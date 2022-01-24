package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=recommend

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

	// timeoutSeconds specifies the seconds of one recommendation process.
	// Default value is 600(for 10 minutes).
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty"`
}

// RecommendationStatus represents the current state of a recommendation.
type RecommendationStatus struct {
	// +optional
	RecommendedValue string `json:"recommendedValue,omitempty"`

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
	PeriodSeconds *int64 `json:"periodSeconds,omitempty"`
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=cs

// ConfigSet represents the configuration set for recommendation.
type ConfigSet struct {
	metav1.TypeMeta   `json:",inline"`
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
	Namespace  string `json:"namespace,omitempty"`
	// +optional
	Kind       string `json:"kind,omitempty"`
	// +optional
	Name       string `json:"name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigSetList is a list of ConfigSet.
type ConfigSetList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ConfigSet `json:"items" protobuf:"bytes,2,rep,name=items"`
}