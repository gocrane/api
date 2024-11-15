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

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=pq,path=podqoss
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PodQOS struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodQOSSpec `json:"spec"`

	Status PodQOSStatus `json:"status,omitempty"`
}

type PodQOSSpec struct {
	// LabelSelector is a label query over pods that should match the PodQOS
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	// A scope selector represents the AND of the selectors represented
	// by the scoped-resource selector requirements.
	ScopeSelector *ScopeSelector `json:"scopeSelector,omitempty"`

	// ResourceQOS describe the QOS limit for cpu,memory,netIO,diskIO and so on.
	ResourceQOS ResourceQOS `json:"resourceQOS,omitempty"`

	// QualityProbe defines the way to probe a pod
	PodQualityProbe PodQualityProbe `json:"podQualityProbe,omitempty"`

	// AllowedActions limits the set of actions that the pods is allowed to perform by NodeQOS
	// Example: ["Throttle", "Evict"]
	AllowedActions []string `json:"allowedActions,omitempty"`
}

type ScopeName string

const (
	// QOSClassSelectors used to select workload with designated QOSClass, example: ["BestEffort","Guaranteed"]
	QOSClassSelector ScopeName = "QOSClass"
	// PrioritySelectors used to select workload with designated priority, contains an operator and values.
	PrioritySelectors ScopeName = "Priority"
	// NamespaceSelectors used to select workload by namespace
	NamespaceSelectors ScopeName = "Namespace"
)

type ScopeSelector struct {
	// A list of scope selector requirements by scope of the resources.
	// +optional
	MatchExpressions []ScopedResourceSelectorRequirement `json:"matchExpressions,omitempty"`
}

// A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator
// that relates the scope name and values.
type ScopedResourceSelectorRequirement struct {
	// The name of the scope that the selector applies to.
	ScopeName ScopeName `json:"scopeName"`
	// Represents a scope's relationship to a set of values.
	// Valid operators are In, NotIn.
	Operator corev1.ScopeSelectorOperator `json:"operator"`
	// An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty.
	// This array is replaced during a strategic merge patch.
	// +optional
	Values []string `json:"values,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodQOSList contains a list of PodQOS
type PodQOSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodQOS `json:"items"`
}

type ResourceQOS struct {
	CPUQOS    *CPUQOS    `json:"cpuQOS,omitempty"`
	MemoryQOS *MemoryQOS `json:"memoryQOS,omitempty"`
	NetIOQOS  *NetIOQOS  `json:"netIOQOS,omitempty"`
	DiskIOQOS *DiskIOQOS `json:"diskIOQOS,omitempty"`
}

// CompressionPreference provides a quick way to set the frequency, ratio and size of compression.
type CompressionPreference string

type CompressionOversold string

const (
	PreferenceTiny     CompressionPreference = "Tiny"
	PreferenceNormal   CompressionPreference = "Normal"
	PreferenceFileOnly CompressionPreference = "FileOnly"
	PreferenceAnonOnly CompressionPreference = "AnonOnly"

	OversoldTransparent CompressionOversold = "Transparent"
	OversoldNone        CompressionOversold = "None"
	OversoldAllow       CompressionOversold = "Allow"
)

type MemoryCompression struct {
	// +kubebuilder:validation:Default=false
	Enable bool `json:"enable,omitempty"`

	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=4
	CompressionLevel int `json:"compressionLevel,omitempty"`

	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Tiny;Normal;FileOnly;AnonOnly
	// +kubebuilder:default=Tiny
	// +optional
	Preference CompressionPreference `json:"preference"`

	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Transparent;None;Allow
	// +kubebuilder:default=Allow
	// +optional
	Oversold CompressionOversold `json:"oversold"`
}

type CPUQOS struct {
	// CPUPriority define the cpu priority for the pods.
	// CPUPriority range [0,7], 0 is the highest level.
	// When the cpu resource is shortage, the low level pods would be throttled
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	// +optional
	CPUPriority *int32 `json:"cpuPriority,omitempty"`

	ContainerPriority map[string]*int32 `json:"containerPriority,omitempty"`

	CPUBurst CPUBurst `json:"cpuBurst,omitempty"`

	HtIsolation HtIsolation `json:"htIsolation,omitempty"`

	CPUSet CPUSet `json:"cpuSet,omitempty"`

	RDT RDT `json:"rdt,omitempty"`

	ContainerRDT map[string]RDT `json:"containerRdt,omitempty"`
}

type RDTValue map[string]string

type RDT struct {
	L3 RDTValue `json:"l3,omitempty"`
	MB RDTValue `json:"mb,omitempty"`
}

type CPUSet struct {
	// none/exclusive/share
	// Provide three polices for cpuset manager:
	// - none: containers of this pod shares a set of cpus which not allocated to exclusive containers
	// - exclusive:  containers of this pod monopolize the allocated CPUs , other containers not allowed to use.
	// - share: containers of this pod runs in theallocated  CPUs , but other containers can also use.
	CPUSet string `json:"cpuSet,omitempty"`
}

type HtIsolation struct {
	Enable bool `json:"enable,omitempty"`
}

type CPUBurst struct {
	// BurstQuota define the burst quota for the pods.
	BurstQuota string `json:"burstQuota,omitempty"`
}

type MemoryQOS struct {
	// MemoryPriority define the memory priority for the pods.
	// MemoryPriority range [0,7], 0 is the highest level.
	// When the memory resource is shortage, the low level pods would be OOM Killed earlier
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	// +optional
	MemoryPriority *int32 `json:"memPriority,omitempty"`

	MemAsyncReclaim   MemAsyncReclaim   `json:"memAsyncReclaim,omitempty"`
	MemWatermark      MemWatermark      `json:"memWatermark,omitempty"`
	MemPageCacheLimit MemPageCacheLimit `json:"memPageCacheLimit,omitempty"`

	MemoryCompression MemoryCompression `json:"memoryCompression,omitempty"`
}

type MemPageCacheLimit struct {
	PageCacheMaxRatio     *int64 `json:"pageCacheMaxRatio,omitempty"`
	PageCacheReclaimRatio *int64 `json:"pageCacheReclaimRatio,omitempty"`
}

type MemAsyncReclaim struct {
	AsyncRatio          *int64 `json:"asyncRatio,omitempty"`
	AsyncDistanceFactor *int64 `json:"asyncDistanceFactor,omitempty"`
}

// MemWatermark to set memory watermark priority
type MemWatermark struct {
	WatermarkRatio *int `json:"watermarkRatio,omitempty"`
}

type NetIOQOS struct {
	NetIOPriority      *int64           `json:"netIOPriority,omitempty"`
	ContainersPriority map[string]int64 `json:"containersPriority,omitempty"`
	NetIOLimits        NetIOLimits      `json:"netIOLimits,omitempty"`
	DevNetIOLimits     DevNetIOLimits   `json:"devNetIOLimits,omitempty"`
	WhitelistPorts     WhitelistPorts   `json:"whitelistPorts,omitempty"`
}

type DevNetIOLimits map[string]NetIOLimits

type WhitelistPorts struct {
	LPorts string `json:"lports"`
	RPorts string `json:"rports"`
}

type NetIOLimits struct {
	RXBps int64 `json:"rxBps"`
	TXBps int64 `json:"txBps"`
}

type DiskIOQOS struct {
	DiskIOWeight DiskIOWeight `json:"diskIOWeight,omitempty"`
	DiskIOLimit  DiskIOLimit  `json:"diskIOLimit,omitempty"`
}

type DiskIOWeight struct {
	Weight int64 `json:"weight,omitempty"`
}

type DiskIOLimit struct {
	ReadIOPS  int64 `json:"readIOps,omitempty"`
	WriteIOPS int64 `json:"writeIOps,omitempty"`
	ReadBPS   int64 `json:"readBps,omitempty"`
	WriteBPS  int64 `json:"writeBps,omitempty"`
}

type PodQOSStatus struct {
}

type PodQualityProbe struct {
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *corev1.HTTPGetAction `json:"httpGet,omitempty"`

	// TimeoutSeconds is the timeout for request.
	// Defaults to 0, no timeout forever
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

// PodQOSEnsurancePolicyStatus defines the observed status of PodQOSEnsurancePolicy
type PodQOSEnsurancePolicyStatus struct {
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=nq,path=nodeqoss
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeQOS is the Schema for the nodeqos API
type NodeQOS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeQOSSpec   `json:"spec"`
	Status NodeQOSStatus `json:"status,omitempty"`
}

// NodeQOSSpec defines the desired status of NodeQOS
type NodeQOSSpec struct {
	// Selector is a label query over pods that should match the policy
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// NodeQualityProbe defines the way to probe a node
	NodeQualityProbe NodeQualityProbe `json:"nodeQualityProbe,omitempty"`

	// Rules is an array of Rules and its corresponding action
	Rules []Rule `json:"rules,omitempty"`

	// ElasticCpuLimit is the cpu limit for workloads using elastic cpu in the node
	ElasticCpuLimit ElasticCpuLimit `json:"elasticCpuLimit,omitempty"`

	// MemoryLimit is the mem limit in the node
	MemoryLimit MemLimit `json:"memLimit,omitempty"`

	MemoryCompression NodeMemoryCompression `json:"memoryCompression,omitempty"`

	// NetLimits is the net IO limit in the node
	NetLimits NetLimits `json:"netLimits,omitempty"`
}

type NodeMemoryCompression struct {
	// +kubebuilder:validation:Default=false
	Enable bool `json:"enable,omitempty"`
}

type NetLimits struct {
	RXBpsMin *int64 `json:"rxBpsMin"`
	RXBpsMax *int64 `json:"rxBpsMax"`
	TXBpsMin *int64 `json:"txBpsMin"`
	TXBpsMax *int64 `json:"txBpsMax"`
}

type MemLimit struct {
	PageCacheLimitGlobal     *bool  `json:"pageCacheLimitGlobal,omitempty"`
	PageCacheLimitRetryTimes *int64 `json:"pageCacheLimitRetryTimes,omitempty"`
}

type NodeQualityProbe struct {
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *corev1.HTTPGetAction `json:"httpGet,omitempty"`

	// NodeLocalGet specifies how to request node local
	// +optional
	NodeLocalGet *NodeLocalGet `json:"nodeLocalGet,omitempty"`

	// TimeoutSeconds is the timeout for request.
	// Defaults to 0, no timeout forever.
	// +optional
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

type NodeLocalGet struct {
	// LocalCacheTTLSeconds is the cache expired time.
	// Defaults to 60
	// +optional
	// +kubebuilder:default=60
	LocalCacheTTLSeconds int32 `json:"localCacheTTLSeconds,omitempty"`
}

// QOSEnsurance defines the policy that
type QOSEnsurance struct {
	Rule Rule `json:"rule,omitempty"`
}

type ElasticCpuLimit struct {
	// ElasticNodeCpuLimit is the total cpu usage limit for the workloads which use extended resource in node
	// Suppress the workloads which use extended resource when the CPU usage of the node exceedes ElasticNodeCpuLimit
	ElasticNodeCpuLimit ElasticNodeCpuLimit `json:"elasticNodeCpuLimit,omitempty"`

	// Limit the amount of single core CPU that can be used by workloads which use extended resource
	ElasticCoreCpuLimit []ElasticCoreCpuLimit `json:"elasticCoreCpuLimit,omitempty"`

	// Limit the amount of single core CPU and its corresponding time that can be used by workloads which use extended resource
	ElasticCoreCpuLimitPeriod []ElasticCoreCpuLimitPeriod `json:"elasticCoreCpuLimitPeriod,omitempty"`

	// Workloads which use extended resource only run on CPUs where high priority tasks are not running
	ElasticCpuAvoidance ElasticCpuAvoidance `json:"elasticCpuAvoidance,omitempty"`
}

type ElasticNodeCpuLimit struct {
	Percent int64 `json:"percent,omitempty"`
}

type ElasticCoreCpuLimit struct {
	CoreNum string `json:"coreNum,omitempty"`
	Percent int64  `json:"percent,omitempty"`
}

type ElasticCoreCpuLimitPeriod struct {
	SchduleTime string `json:"offlineCpuLimit,omitempty"`
	CoreNum     string `json:"coreNum,omitempty"`
	Percent     int64  `json:"percent,omitempty"`
}

type ElasticCpuAvoidance struct {
	Enable bool `json:"enable,omitempty"`
}

type Rule struct {
	// Name of the objective ensurance
	Name string `json:"name,omitempty"`

	// Metric rule define the metric identifier and target
	MetricRule *MetricRule `json:"metricRule,omitempty"`

	// How many times the rule is reach, to trigger avoidance action.
	// Defaults to 1. Minimum value is 1.
	// +optional
	// +kubebuilder:default=1
	AvoidanceThreshold int32 `json:"avoidanceThreshold,omitempty"`

	// How many times the rule can restore.
	// Defaults to 1. Minimum value is 1.
	// +optional
	// +kubebuilder:default=1
	RestoreThreshold int32 `json:"restoreThreshold,omitempty"`

	// Avoidance action to be executed when the rule triggered
	AvoidanceActionName string `json:"actionName"`

	// Action only preview, not to do the real action.
	// Default AvoidanceActionStrategy is None.
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

// NodeQOSStatus defines the observed status of NodeQOS
type NodeQOSStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeQOSList contains a list of NodeQOS
type NodeQOSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeQOS `json:"items"`
}

type AvoidanceActionSpec struct {
	// CoolDownSeconds is the seconds for cool down when do avoidance.
	// Defaults to 300. Minimum value is 1.
	// +optional
	// +kubebuilder:default=300
	CoolDownSeconds int32 `json:"coolDownSeconds,omitempty"`

	// Throttle describes the throttling action
	// +optional
	Throttle *ThrottleAction `json:"throttle,omitempty"`

	// Eviction describes the eviction action
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
	// MinCPURatio is the min of cpu ratio for low level pods,
	// for example: the pod limit is 4096, ratio is 10, the minimum is 409.
	// MinCPURatio range [0,100]
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	MinCPURatio int32 `json:"minCPURatio,omitempty"`

	// StepCPURatio is the step of cpu share and limit for once down-size.
	// StepCPURatio range [0,100]
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
	// If this value is nil, the pod's terminationGracePeriodSeconds will be used.
	// Otherwise, this value overrides the value provided by the pod spec.
	// Value must be non-negative integer. The value zero indicates delete immediately.
	// +optional
	TerminationGracePeriodSeconds *int32 `json:"terminationGracePeriodSeconds,omitempty"`
}

// AvoidanceActionStatus defines the desired status of AvoidanceAction
type AvoidanceActionStatus struct {
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope="Cluster",shortName=avoid

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
