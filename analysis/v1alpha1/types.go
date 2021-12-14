package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	autoscalingapi "github.com/gocrane/api/autoscaling/v1alpha1"
)

type AnalysisType string

const (
	AnalysisTypeHPA      AnalysisType = "HPA"
	AnalysisTypeResource AnalysisType = "Resource"
)

type CompletionStrategyType string

const (
	CompletionStrategyPeriodical CompletionStrategyType = "Periodical"
	CompletionStrategyOnce       CompletionStrategyType = "Once"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
}

// RecommendationStatus represents the current state of a recommendation.
type RecommendationStatus struct {
	// EffectiveHPA is the recommendation for effective HPA.
	// +optional
	EffectiveHPA *EffectiveHorizontalPodAutoscalerRecommendation `json:"effectiveHPA,omitempty"`

	// ResourceRequest is the recommendation for containers' cpu/mem requests.
	// +optional
	ResourceRequest *ResourceRequestRecommendation `json:"resourceRequest,omitempty"`

	// Conditions is an array of current recommendation conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastUpdateTime is last time we got an update on this status.
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// LastSuccessfulTime is the last time the recommendation successfully completed.
	// +optional
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`
}

type EffectiveHorizontalPodAutoscalerRecommendation struct {
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`

	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`

	// +optional
	Metrics []autoscalingv2.MetricSpec `json:"metrics,omitempty"`

	// +optional
	Prediction *autoscalingapi.Prediction `json:"prediction,omitempty"`
}

type ResourceRequestRecommendation struct {
	// +optional
	Containers []ContainerRecommendation `json:"containers,omitempty"`
}

type ContainerRecommendation struct {
	// +required
	// +kubebuilder:validation:Required
	ContainerName string `json:"containerName"`

	// +required
	// +kubebuilder:validation:Required
	Target corev1.ResourceList `json:"target"`
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
	PeriodSeconds *int64 `json:"periodSeconds"`
}

// AnalyticsStatus represents the current state of an analytics item.
type AnalyticsStatus struct {
	// LastSuccessfulTime is the last time the recommendation successfully completed.
	// +optional
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`

	// Conditions is an array of current analytics conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Recommendations is a list of pointers to recommendations that are updated by this analytics.
	// +optional
	// +listType=atomic
	Recommendations []corev1.ObjectReference `json:"recommendations,omitempty"`
}

// ResourceSelector describes how the resources will be selected.
type ResourceSelector struct {
	// Kind of the resource, e.g. Deployment
	Kind string `json:"kind"`

	// API version of the resource, e.g. "apps/v1"
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// Name of the resource.
	// +optional
	Name string `json:"name,omitempty"`

	// +optional
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsList is a list of Analytics items.
type AnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Analytics `json:"items"`
}
