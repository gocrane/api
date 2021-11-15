package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	// select pod used labels
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	//pod quality probe
	QualityProbe QualityProbe `json:"qualityProbe,omitempty"`

	//pod objective ensurance check and action
	ObjectiveEnsurance []ObjectiveEnsurance `json:"objectiveEnsurance,omitempty"`
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
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	//node quality probe
	NodeQualityProbe NodeQualityProbe `json:"nodeQualityProbe,omitempty"`

	//node objective ensurance check and action
	ObjectiveEnsurances []ObjectiveEnsurance `json:"objectiveEnsurances,omitempty"`
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

type AvoidanceActionSpec struct {
	// how long it should wait between triggered scheduling
	// default is 300s
	// +optional
	CoolDownSeconds *int64 `json:"coolDownSeconds,omitempty"`

	//Action to Throttle resource
	// +optional
	Throttle *ThrottleAction `json:"Throttle,omitempty"`

	//Action to evict low level pods
	// +optional
	Eviction *EvictionAction `json:"eviction,omitempty"`

	// Description is an arbitrary string that usually provides guidelines on
	// when this action should be used.
	// +optional
	Description string `json:"description,omitempty"`
}

// AvoidanceActionStatus defines the desired status of AvoidanceAction
type AvoidanceActionStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AvoidanceActionList contains a list of AvoidanceAction
type AvoidanceActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AvoidanceAction `json:"items"`
}

type QualityProbe struct {
	Handler `json:",inline"`
	// Init delay time for handler, default is 5s
	// +optional
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`

	// Timeout for request, default is 0, instead not timeout
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
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
	// currently supported
	// CPU usage, CPU load, Memory Usage, DiskIO
	// +optional
	HTTPGet *HTTPGet `json:"httpGet,omitempty"`

	// Get node metric from local
	// +optional
	NodeLocalGet *NodeLocalGet `json:"nodeLocalGet,omitempty"`

	// Init delay time for handler, default is 5s
	// +optional
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty"`
	// Timeout for request, default is 0, instead not timeout
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

type NodeLocalGet struct {
	// default is 60s
	// +optional
	LocalCacheTTLSeconds *int32 `json:"localCacheTTLSeconds,omitempty"`
	// default is 60s
	// +optional
	MaxHousekeepingIntervalSeconds *int32 `json:"maxHousekeepingIntervalSeconds,omitempty"`
}

type ThrottleAction struct {
	// +optional
	CPUThrottle CPUThrottle `json:"cpuThrottle,omitempty"`

	// +optional
	MemoryThrottle MemoryThrottle `json:"memoryThrottle,omitempty"`
}

type CPUThrottle struct {
	// how long it waits for each compress step
	// default is 10s
	// +optional
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty"`

	//the min of cpu ratio for low level pods
	//example: the pod limit is 4096, ratio is 10, the min is 409
	// +optional
	MinCPURatio uint64 `json:"minCPURatio,omitempty"`

	//the step of cpu share and limit for once down-size (1-100)
	// +optional
	StepCPURatio uint64 `json:"stepCPURatio,omitempty"`
}

type MemoryThrottle struct {
	// how long it waits for each compress step
	// default is 10s
	// +optional
	IntervalSeconds *int32 `json:"intervalSeconds,omitempty"`

	// to force gc the page cache of low level pods
	// +optional
	ForceGC bool `json:"forceGC,omitempty"`
}

type EvictionAction struct {
	// duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates delete immediately.
	// +optional
	DeletionGracePeriodSeconds *int32 `json:"deletionGracePeriodSeconds,omitempty"`
}

// ObjectiveEnsurance defines the rule if anomaly reached
// and if the rule reached, do what action
type ObjectiveEnsurance struct {
	// Name of the objective ensurance
	Name string `json:"name,omitempty"`

	// Metric rule define the metric identifier and target
	MetricRule *MetricRule `json:"metricRule,omitempty"`

	// How many times the rule is reach, to trigger action, default is 1
	ReachedThreshold int32 `json:"reachedThreshold,omitempty"`

	// How many times the rule can restore, default is 1
	RestoredThreshold int32 `json:"restoredThreshold,omitempty"`

	// Avoidance action when be triggered
	AvoidanceActionName string `json:"actionName"`

	// Action only dry run,not to do the real action
	// +optional
	DryRun bool `json:"dryRun,omitempty"`
}

type MetricRule struct {
	// Metric identifies the target metric by name and selector
	Metric MetricIdentifier `json:"metric"`

	// Target specifies the target value for the given metric
	Target *MetricTarget `json:"target"`
}

// MetricIdentifier defines the name and optionally selector for a metric
type MetricIdentifier struct {
	// Name is the name of the given metric
	Name string `json:"name"`
	// Selector is the selector for the given metric
	// it is the string-encoded form of a standard kubernetes label selector
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

// MetricTarget defines the target value or utilization of a specific metric
type MetricTarget struct {
	// Type represents whether the metric type is Value and Utilization
	Type MetricTargetType `json:"type"`
	// Value is the target value of the metric (as a quantity).
	Value *resource.Quantity `json:"value,omitempty"`
	// Utilization is the target value of a percentage of the resource for pods.
	Utilization *int32 `json:"utilization,omitempty"`
}

// MetricTargetType specifies the type of metric being targeted, and should be either
// "Value", "AverageValue", or "Utilization"
type MetricTargetType string

const (
	// UtilizationMetricType is a possible value for MetricTarget.Type.
	UtilizationMetricType MetricTargetType = "Utilization"
	// ValueMetricType is a possible value for MetricTarget.Type.
	ValueMetricType MetricTargetType = "Value"
)
