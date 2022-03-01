package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AvoidanceActionStrategy string

const (
	// AvoidanceActionStrategyNone do the action when the rules triggered.
	AvoidanceActionStrategyNone AvoidanceActionStrategy = "None"
	// AvoidanceActionStrategyPreview is the preview for QosEnsuranceStrategyNone.
	AvoidanceActionStrategyPreview AvoidanceActionStrategy = "Preview"
)

// PodQOSEnsurancePolicySpec defines the desired status of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicySpec struct {
	// Selector is a label query over pods that should match the policy
	Selector metav1.LabelSelector `json:"selector,omitempty"`

	//QualityProbe defines the way to probe a pod
	QualityProbe QualityProbe `json:"qualityProbe,omitempty"`

	// ObjectiveEnsurances is an array of ObjectiveEnsurance
	ObjectiveEnsurances []ObjectiveEnsurance `json:"objectiveEnsurance,omitempty"`
}

type QualityProbe struct {
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *corev1.HTTPGetAction `json:"httpGet,omitempty"`
	// Init delay time for handler, default is 5s
	// +optional
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// Timeout for request, default is 0, instead not timeout
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// PodQOSEnsurancePolicyStatus defines the observed status of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicyStatus struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodQOSEnsurancePolicy is the Schema for the podqosensurancepolicies API
type PodQOSEnsurancePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodQOSEnsurancePolicySpec   `json:"spec"`
	Status PodQOSEnsurancePolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodQOSEnsurancePolicyList contains a list of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodQOSEnsurancePolicy `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeQOSEnsurancePolicy is the Schema for the nodeqosensurancepolicies API
type NodeQOSEnsurancePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeQOSEnsurancePolicySpec   `json:"spec"`
	Status NodeQOSEnsurancePolicyStatus `json:"status,omitempty"`
}

// NodeQOSEnsurancePolicySpec defines the desired status of NodeQOSEnsurancePolicy
type NodeQOSEnsurancePolicySpec struct {
	// Selector is a label query over pods that should match the policy
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// NodeQualityProbe defines the way to probe a node
	NodeQualityProbe NodeQualityProbe `json:"nodeQualityProbe,omitempty"`

	// ObjectiveEnsurances is an array of ObjectiveEnsurance
	ObjectiveEnsurances []ObjectiveEnsurance `json:"objectiveEnsurances,omitempty"`
}

type NodeQualityProbe struct {
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *corev1.HTTPGetAction `json:"httpGet,omitempty"`

	// NodeLocalGet specifies how to request node local
	// +optional
	NodeLocalGet *NodeLocalGet `json:"nodeLocalGet,omitempty"`

	// TimeoutSeconds is the timeout for request
	// The default value is 10 seconds
	// +optional
	// +kubebuilder:default=10
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

type NodeLocalGet struct {
	// the default LocalCacheTTLSeconds is 60s
	// +optional
	// +kubebuilder:default=60
	LocalCacheTTLSeconds *int32 `json:"localCacheTTLSeconds,omitempty"`
}

// ObjectiveEnsurance defines the policy that
type ObjectiveEnsurance struct {
	// Name of the objective ensurance
	Name string `json:"name,omitempty"`

	// Metric rule define the metric identifier and target
	MetricRule *MetricRule `json:"metricRule,omitempty"`

	// How many times the rule is reach, to trigger avoidance action, default is 1
	// +optional
	// +kubebuilder:default=1
	AvoidanceThreshold *int32 `json:"avoidanceThreshold,omitempty"`

	// How many times the rule can restore, default is 1
	// +optional
	// +kubebuilder:default=1
	RestoreThreshold *int32 `json:"restoreThreshold,omitempty"`

	// Avoidance action when be triggered
	AvoidanceActionName string `json:"actionName"`

	// Action only preview, not to do the real action
	// the default AvoidanceActionStrategy is None.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=None;Preview
	// +kubebuilder:default=None
	Strategy AvoidanceActionStrategy `json:"strategy,omitempty"`
}

type MetricRule struct {
	// Name is the name of the given metric
	Name string `json:"name"`
	// Selector is the selector for the given metric
	// it is the string-encoded form of a standard kubernetes label selector
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
	// Value is the target value of the metric (as a quantity).
	Value resource.Quantity `json:"value,omitempty"`
}

// NodeQOSEnsurancePolicyStatus defines the observed status of NodeQOSEnsurancePolicy
type NodeQOSEnsurancePolicyStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeQOSEnsurancePolicyList contains a list of NodeQOSEnsurancePolicy
type NodeQOSEnsurancePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeQOSEnsurancePolicy `json:"items"`
}

type AvoidanceActionSpec struct {
	// CoolDownSeconds is the seconds for cool down when do avoidance
	// default is 300s
	// +optional
	// +kubebuilder:default=300
	CoolDownSeconds *int32 `json:"coolDownSeconds,omitempty"`

	// Throttle describes the throttling action
	// +optional
	Throttle *ThrottleAction `json:"throttle,omitempty"`

	//Eviction describes the eviction action
	// +optional
	Eviction *EvictionAction `json:"eviction,omitempty"`

	// Description is an arbitrary string that usually provides guidelines on
	// when this action should be used.
	// +optional
	// +kubebuilder:validation:MaxLength=1024
	Description string `json:"description,omitempty"`
}

type ThrottleAction struct {
	// +optional
	CPUThrottle CPUThrottle `json:"cpuThrottle,omitempty"`

	// +optional
	MemoryThrottle MemoryThrottle `json:"memoryThrottle,omitempty"`
}

type CPUThrottle struct {
	// MinCPURatio is the min of cpu ratio for low level pods.
	// example: the pod limit is 4096, ratio is 10, the minimum is 409.
	// MinCPURatio range [0,100]
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	MinCPURatio int32 `json:"minCPURatio,omitempty"`

	// StepCPURatio is the step of cpu share and limit for once down-size.
	// StepCPURatio range [0,100]
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	StepCPURatio int32 `json:"stepCPURatio,omitempty"`
}

type MemoryThrottle struct {
	// ForceGC means force gc page cache for pods with low priority
	// +optional
	ForceGC bool `json:"forceGC,omitempty"`
}

type EvictionAction struct {
	// TerminationGracePeriodSeconds is the duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates delete immediately.
	// +optional
	// +kubebuilder:default=30
	TerminationGracePeriodSeconds *int32 `json:"terminationGracePeriodSeconds,omitempty"`
}

// AvoidanceActionStatus defines the desired status of AvoidanceAction
type AvoidanceActionStatus struct {
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope="Cluster"

// AvoidanceAction defines Avoidance action
type AvoidanceAction struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AvoidanceActionSpec   `json:"spec"`
	Status AvoidanceActionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AvoidanceActionList contains a list of AvoidanceAction
type AvoidanceActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvoidanceAction `json:"items"`
}
