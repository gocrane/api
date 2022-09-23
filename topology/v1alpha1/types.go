package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CPUManagerPolicy represents policy of the crane agent cpu manager.
type CPUManagerPolicy string

const (
	// CPUManagerPolicyNone name of none policy.
	CPUManagerPolicyNone CPUManagerPolicy = "None"
	// CPUManagerPolicyStatic is the name of the static policy.
	CPUManagerPolicyStatic CPUManagerPolicy = "Static"
)

// TopologyManagerPolicy represents policy of the crane agent resource management component.
type TopologyManagerPolicy string

// Constants of type TopologyManagerPolicy represent policy of the agent
// node's resource management component. It's TopologyManager in kubele.
const (
	// TopologyManagerPolicyNone policy is the default policy and does not perform any topology alignment.
	TopologyManagerPolicyNone TopologyManagerPolicy = "None"
	// TopologyManagerPolicySingleNUMANodePodLevel enables pod level resource counting, this policy assumes
	// TopologyManager policy single-numa-node also was set on the node.
	TopologyManagerPolicySingleNUMANodePodLevel TopologyManagerPolicy = "SingleNUMANodePodLevel"
)

// ZoneType string describes a topology type for a zone
type ZoneType string

const (
	ZoneTypeNode   ZoneType = "Node"
	ZoneTypeSocket ZoneType = "Socket"
	ZoneTypeCore   ZoneType = "Core"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope="Cluster",shortName=nrt
// +kubebuilder:printcolumn:name="CRANE CPU MANAGER POLICY",type=string,JSONPath=".craneManagerPolicy.cpuManagerPolicy",description="CPUManagerPolicy represents policy of the crane agent cpu manager."
// +kubebuilder:printcolumn:name="CRANE TOPOLOGY MANAGER POLICY",type=string,JSONPath=".craneManagerPolicy.topologyManagerPolicy",description="TopologyManagerPolicy represents policy of the crane agent resource management component. Defaults to SingleNUMANodePodLevel."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// NodeResourceTopology describes node resources and their topology.
type NodeResourceTopology struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// CraneManagerPolicy describes the associated manager policy of crane agent.
	// +required
	CraneManagerPolicy ManagerPolicy `json:"craneManagerPolicy"`

	// Reserved describes resources reserved for system and kubernetes components.
	// +optional
	Reserved corev1.ResourceList `json:"reserved,omitempty"`

	// Zones represents all resource topology zones of a node.
	// +optional
	Zones ZoneList `json:"zones,omitempty"`

	// Attributes represents node attributes if any.
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`
}

// ManagerPolicy describes the associated manager policy.
type ManagerPolicy struct {
	// CPUManagerPolicy represents policy of the crane agent cpu manager.
	// +kubebuilder:validation:Enum=None;Static
	// +required
	CPUManagerPolicy CPUManagerPolicy `json:"cpuManagerPolicy"`

	// TopologyManagerPolicy represents policy of the crane agent resource management component.
	// Defaults to SingleNUMANodePodLevel.
	// +kubebuilder:validation:Enum=None;SingleNUMANodePodLevel
	// +kubebuilder:default=SingleNUMANodePodLevel
	// +required
	TopologyManagerPolicy TopologyManagerPolicy `json:"topologyManagerPolicy"`
}

// Zone represents a resource topology zone, e.g. socket, node, die or core.
type Zone struct {
	// Name represents the zone name.
	// +required
	Name string `json:"name"`

	// Type represents the zone type.
	// +kubebuilder:validation:Enum=Node;Socket;Core
	// +required
	Type ZoneType `json:"type"`

	// Parent represents the name of parent zone.
	// +optional
	Parent string `json:"parent,omitempty"`

	// Costs represents the cost between different zones.
	// +optional
	Costs CostList `json:"costs,omitempty"`

	// Attributes represents zone attributes if any.
	// +optional
	Attributes map[string]string `json:"attributes,omitempty"`

	// Resources represents the resource info of the zone.
	// +optional
	Resources *ResourceInfo `json:"resources,omitempty"`
}

// ZoneList contains an array of Zone objects.
type ZoneList []Zone

// ResourceInfo contains information about one resource type.
type ResourceInfo struct {
	// Capacity of the resource, corresponding to capacity in node status, i.e.
	// total amount of this resource that the node has.
	// +optional
	Capacity corev1.ResourceList `json:"capacity,omitempty"`

	// Allocatable quantity of the resource, corresponding to allocatable in
	// node status, i.e. total amount of this resource available to be used by
	// pods.
	// +optional
	Allocatable corev1.ResourceList `json:"allocatable,omitempty"`

	// ReservedCPUNums specifies the cpu numbers reserved for the host level system threads and kubernetes related threads.
	// +optional
	ReservedCPUNums int32 `json:"reservedCPUNums,omitempty"`
}

// CostInfo describes the cost (or distance) between two Zones.
type CostInfo struct {
	// Name represents the zone name.
	// +required
	Name string `json:"name"`

	// Value represents the cost value.
	// +required
	Value int64 `json:"value"`
}

// CostList contains an array of CostInfo objects.
type CostList []CostInfo

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeResourceTopologyList is a list of NodeResourceTopology resources
type NodeResourceTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodeResourceTopology `json:"items"`
}
