package v1alpha1

import (
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// HighestLevel is the highest level for pod.
	HighestLevel = int32(9999)

	// LowestLevel is the lowest level pod.
	LowestLevel = int32(0)
)

// URIScheme identifies the scheme used for connection to a host for Get actions
type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "HTTPS"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceLevel defines the service level for pods
type ServiceLevel struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceLevelSpec `json:"spec"`

	Status ServiceLevelStatus `json:"status,omitempty"`
}

type ServiceLevelSpec struct {
	// The value of pods level. This is the actual level that pods
	// receive when they have the name of this class in their pod ensurance policy.
	// Integer value range （0～9999), the highest level is 9999.
	// The priority is PodQOSClassValue+Value
	// Example: pod qos guaranteed, value 11 , priority 20011
	// Example: pod qos burstable, value 10 , priority 10010
	// Example: pod qos bestEffort, value 19 , priority 19
	Value int32 `json:"value,omitempty"`

	// The pod qos class is define the pod's qos class,which is align with k8s's pod qos class
	// The service level value for pod qos : (guaranteed 20000, burstable 10000, besteffort 0)
	PodQosClass v1.PodQOSClass `json:"podQosClass,omitempty"`

	// QosClassDefault specifies whether this serviceLevel should be considered as
	// the default level for pods that do not selected by any serviceLevel object.
	// +optional
	QosClassDefault bool `json:"qosClassDefault,omitempty"`

	// Description is an arbitrary string that usually provides guidelines on
	// when this qos level class should be used.
	// +optional
	Description string `json:"description,omitempty"`
}

// ServiceLevelStatus defines the desired status of ServiceLevel
type ServiceLevelStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceLevelList contains a list of ServiceLevel
type ServiceLevelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceLevel `json:"items"`
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

// PodQOSEnsurancePolicySpec defines the desired status of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicySpec struct {

	// service level for pods
	ServiceLevelName string `json:"serviceLevelName"`

	// select pod used labels
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	//pod quality probe
	QualityProbe QualityProbe `json:"qualityProbe,omitempty"`
}

// PodQOSEnsurancePolicyStatus defines the observed status of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicyStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodQOSEnsurancePolicyList contains a list of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodQOSEnsurancePolicy `json:"items"`
}

// +genclient
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
	//select nodes use labels
	LabelSelector metav1.LabelSelector `json:"labelSelector"`

	//node quality probe
	NodeQualityProbe NodeQualityProbe `json:"nodeQualityProbe,omitempty"`
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

type QualityProbe struct {
	Handler             `json:",inline"`
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty"`
}

// Handler defines a specific action that should be taken
type Handler struct {
	HTTPGet *HTTPGet `json:"httpGet,omitempty"`
}

type HTTPGet struct {
	// Path to access on the HTTP server.
	// +optional
	Path string `json:"path,omitempty"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host string `json:"host,omitempty"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme URIScheme `json:"scheme,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
}

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string `json:"name"`
	// The header field value
	Value string `json:"value"`
}

type NodeQualityProbe struct {
	Handler NodeHandler `json:",inline"`
	// +optional
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
	// +optional
	PeriodSeconds int32 `json:"periodSeconds,omitempty"`
}

type NodeHandler struct {
	// currently supported
	// CPU usage, CPU load, Memory Usage, DiskIO
	// +optional
	HTTPGet *HTTPGet `json:"httpGet,omitempty"`

	// Get node metric from local
	// +optional
	NodeLocalGet *NodeLocalGet `json:"nodeLocalGet,omitempty"`
}

type NodeLocalGet struct {
	// +optional
	LocalCacheTTL time.Duration `json:"localCacheTTL,omitempty"`
	// +optional
	MaxHousekeepingInterval time.Duration `json:"maxHousekeepingInterval,omitempty"`
}
