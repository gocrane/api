package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	autoscalingapi "github.com/gocrane/api/autoscaling/v1alpha1"
)

type AnalysisType string

const (
	AnalysisTypeReplicas AnalysisType = "Replicas"
	AnalysisTypeResource AnalysisType = "Resource"
)

type CompletionStrategyType string

const (
	CompletionStrategyPeriodical CompletionStrategyType = "Periodical"
	CompletionStrategyOnce       CompletionStrategyType = "Once"
)

type AdoptionType string

const (
	AdoptionTypeStatus              AdoptionType = "Status"
	AdoptionTypeStatusAndAnnotation AdoptionType = "StatusAndAnnotation"
	AdoptionTypeAuto                AdoptionType = "Auto"
)

const (
	// ReplicasRecommender name
	ReplicasRecommender string = "Replicas"

	// ResourceRecommender name
	ResourceRecommender string = "Resource"

	// HPARecommender name
	HPARecommender string = "HPA"

	// IdleNodeRecommender name
	IdleNodeRecommender string = "IdleNode"
)

var (
	AllRecommenderType []string
)

func init() {
	AllRecommenderType = append(AllRecommenderType, ReplicasRecommender)
	AllRecommenderType = append(AllRecommenderType, ResourceRecommender)
	AllRecommenderType = append(AllRecommenderType, HPARecommender)
	AllRecommenderType = append(AllRecommenderType, IdleNodeRecommender)
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=recommend
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="TargetKind",type=string,JSONPath=`.spec.targetRef.kind`
// +kubebuilder:printcolumn:name="TargetNamespace",type=string,JSONPath=`.spec.targetRef.namespace`
// +kubebuilder:printcolumn:name="TargetName",type=string,JSONPath=`.spec.targetRef.name`
// +kubebuilder:printcolumn:name="Strategy",type=string,JSONPath=`.spec.completionStrategy.completionStrategyType`
// +kubebuilder:printcolumn:name="PeriodSeconds",type=string,JSONPath=`.spec.completionStrategy.periodSeconds`
// +kubebuilder:printcolumn:name="AdoptionType",type=string,JSONPath=`.spec.adoptionType`
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// Recommendation represents the configuration of a single recommendation.
type Recommendation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec RecommendationSpec `json:"spec,omitempty"`

	// +optional
	Status RecommendationStatus `json:"status,omitempty"`
}

// RecommendationSpec describes the recommendation type and what the recommendation is for.
type RecommendationSpec struct {
	// +required
	// +kubebuilder:validation:Required
	TargetRef corev1.ObjectReference `json:"targetRef"`

	// +required
	// +kubebuilder:validation:Required
	Type AnalysisType `json:"type"`

	// CompletionStrategy indicate how to complete a recommendation.
	// the default CompletionStrategy is Once.
	// +optional
	CompletionStrategy CompletionStrategy `json:"completionStrategy,omitempty"`

	// AdoptionType indicate how to adopt recommendation value to target.
	// the default AdoptionType is StatusAndAnnotation.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Status;StatusAndAnnotation;Auto
	// +kubebuilder:default=StatusAndAnnotation
	AdoptionType AdoptionType `json:"adoptionType,omitempty"`
}

// RecommendationStatus represents the current state of a recommendation.
type RecommendationStatus struct {
	RecommendationContent `json:",inline"`

	// Conditions is an array of current recommendation conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastUpdateTime is last time we got an update on this status.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`
}

// RecommendationContent contains results for one recommendation
type RecommendationContent struct {
	// +optional
	RecommendedValue string `json:"recommendedValue,omitempty"`
	// +optional
	TargetRef corev1.ObjectReference `json:"targetRef"`
	// +optional
	RecommendedInfo string `json:"recommendedInfo,omitempty"`
	// +optional
	CurrentInfo string `json:"currentInfo,omitempty"`
	// +optional
	Action string `json:"action,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecommendationList is a list of recommendations
type RecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Recommendation `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=analytics
// +kubebuilder:printcolumn:name="Type",type=string,JSONPath=`.spec.type`
// +kubebuilder:printcolumn:name="Strategy",type=string,JSONPath=`.spec.completionStrategy.completionStrategyType`
// +kubebuilder:printcolumn:name="PeriodSeconds",type=string,JSONPath=`.spec.completionStrategy.periodSeconds`

// Analytics represents the configuration of an analytics object.
type Analytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec AnalyticsSpec `json:"spec"`

	// +optional
	Status AnalyticsStatus `json:"status,omitempty"`
}

// AnalyticsSpec describes the analytics type, what the analysis is for and how often the analysis routine runs.
type AnalyticsSpec struct {
	// Type is the analysis type, including HPA and resource.
	// +required
	// +kubebuilder:validation:Required
	Type AnalysisType `json:"type"`

	// ResourceSelector indicates how to select resources(e.g. a set of Deployments) for an Analytics.
	// +required
	// +kubebuilder:validation:Required
	ResourceSelectors []ResourceSelector `json:"resourceSelectors"`

	// CompletionStrategy indicate how to complete an Analytics.
	// +optional
	CompletionStrategy CompletionStrategy `json:"completionStrategy"`

	// Override Recommendation configs
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

// CompletionStrategy presents how to complete a recommendation or a recommendation request.
type CompletionStrategy struct {
	// CompletionStrategy indicate the strategy to request an Analytics or Recommendation, value can be "Once" and "Periodical"
	// the default CompletionStrategy is Once.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Once;Periodical
	// +kubebuilder:default=Once
	CompletionStrategyType CompletionStrategyType `json:"completionStrategyType,omitempty"`

	// PeriodSeconds is the duration in seconds for an Analytics or Recommendation.
	// +optional
	PeriodSeconds *int64 `json:"periodSeconds,omitempty"`
}

// AnalyticsStatus represents the current state of an analytics item.
type AnalyticsStatus struct {
	// LastUpdateTime is the last time the status updated.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`

	// Conditions is an array of current analytics conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Recommendations is a list of RecommendationMission that run parallel.
	// +optional
	// +listType=atomic
	Recommendations []RecommendationMission `json:"recommendations,omitempty"`
}

type RecommendationMission struct {
	corev1.ObjectReference `json:",inline"`

	// +optional
	TargetRef corev1.ObjectReference `json:"targetRef"`

	// LastStartTime is last time we start a recommendation mission.
	// +optional
	LastStartTime *metav1.Time `json:"lastStartTime,omitempty"`

	// Message presents the running message for this mission
	// +optional
	Message string `json:"message,omitempty"`

	// RecommenderRef presents recommender info for recommendation mission.
	// +optional
	RecommenderRef Recommender `json:"recommenderRef"`
}

// ResourceSelector describes how the resources will be selected.
type ResourceSelector struct {
	// Kind of the resource, e.g. Deployment
	Kind string `json:"kind"`

	// API version of the resource, e.g. "apps/v1"
	// +optional
	APIVersion string `json:"apiVersion"`

	// Name of the resource.
	// +optional
	Name string `json:"name,omitempty"`

	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsList is a list of Analytics items.
type AnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Analytics `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=cs

// ConfigSet represents the configuration set for recommendation.
type ConfigSet struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Configs []Config `json:"configs,omitempty"`
}

type Config struct {
	// +optional
	Targets    []Target          `json:"targets,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

type Target struct {
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigSetList is a list of ConfigSet.
type ConfigSetList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ConfigSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=rr
// +kubebuilder:printcolumn:name="RunInterval",type=string,JSONPath=`.spec.runInterval`
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// RecommendationRule represents the configuration of an RecommendationRule object.
type RecommendationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec RecommendationRuleSpec `json:"spec"`

	// +optional
	Status RecommendationRuleStatus `json:"status,omitempty"`
}

// RecommendationRuleSpec defines resources and runInterval to recommend
type RecommendationRuleSpec struct {
	// ResourceSelector indicates how to select resources(e.g. a set of Deployments) for a Recommendation.
	// +required
	// +kubebuilder:validation:Required
	ResourceSelectors []ResourceSelector `json:"resourceSelectors"`

	// NamespaceSelector indicates resource namespaces to select from
	NamespaceSelector NamespaceSelector `json:"namespaceSelector"`

	// RunInterval between two recommendation
	RunInterval string `json:"runInterval,omitempty"`

	// List of recommender type to run
	Recommenders []Recommender `json:"recommenders"`
}

// Recommender referring to the Recommender in RecommendationConfiguration
type Recommender struct {

	// Recommender's Name
	Name string `json:"name"`
	// Override Recommendation configs
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

// NamespaceSelector describes how to select namespaces for recommend
type NamespaceSelector struct {
	// Select all namespace if true
	Any bool `json:"any,omitempty"`
	// List of namespace names to select from.
	MatchNames []string `json:"matchNames,omitempty"`
}

// RecommendationRuleStatus represents the current state of an RecommendationRule item.
type RecommendationRuleStatus struct {
	// LastUpdateTime is the last time the status updated.
	// +optional
	LastUpdateTime *metav1.Time `json:"lastUpdateTime,omitempty"`

	// Recommendations is a list of RecommendationMission that run parallel.
	// +optional
	// +listType=atomic
	Recommendations []RecommendationMission `json:"recommendations,omitempty"`

	// RunNumber is the numbers of runs
	// +optional
	RunNumber int32 `json:"runNumber,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecommendationRuleList is a list of RecommendationRule items.
type RecommendationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RecommendationRule `json:"items"`
}

// ProposedRecommendation is the result for one recommendation
type ProposedRecommendation struct {
	// EffectiveHPA is the proposed recommendation for type Replicas
	EffectiveHPA *EffectiveHorizontalPodAutoscalerRecommendation `json:"effectiveHPA,omitempty"`

	// ReplicasRecommendation is the proposed replicas for type Replicas
	ReplicasRecommendation *ReplicasRecommendation `json:"replicasRecommendation,omitempty"`

	// ResourceRequest is the proposed recommendation for type Resource
	ResourceRequest *ResourceRequestRecommendation `json:"resourceRequest,omitempty"`
}

type ReplicasRecommendation struct {
	Replicas *int32 `json:"replicas,omitempty"`
}

type EffectiveHorizontalPodAutoscalerRecommendation struct {
	MinReplicas *int32                     `json:"minReplicas,omitempty"`
	MaxReplicas *int32                     `json:"maxReplicas,omitempty"`
	Metrics     []autoscalingv2.MetricSpec `json:"metrics,omitempty"`
	Prediction  *autoscalingapi.Prediction `json:"prediction,omitempty"`
}

type ResourceRequestRecommendation struct {
	Containers []ContainerRecommendation `json:"containers,omitempty"`
}

type ContainerRecommendation struct {
	ContainerName string       `json:"containerName,omitempty"`
	Target        ResourceList `json:"target,omitempty"`
}

type ResourceList map[corev1.ResourceName]string

type PatchReplicas struct {
	Spec PatchReplicasSpec `json:"spec,omitempty"`
}

type PatchReplicasSpec struct {
	Replicas *int32 `json:"replicas,omitempty"`
}

type PatchResource struct {
	Spec PatchResourceSpec `json:"spec,omitempty"`
}

type PatchResourceSpec struct {
	Template PatchResourcePodTemplateSpec `json:"template"`
}

type PatchResourcePodTemplateSpec struct {
	Spec PatchResourcePodSpec `json:"spec,omitempty"`
}

type PatchResourcePodSpec struct {
	// +patchMergeKey=name
	// +patchStrategy=merge
	Containers []corev1.Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name"`
}
