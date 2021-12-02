package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Type string

const (
	TypeHPA      Type = "HPA"
	TypeResource Type = "Resource"
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
	TargetRef autoscalingv2.CrossVersionObjectReference `json:"targetRef"`

	// +required
	// +kubebuilder:validation:Required
	Type Type `json:"type"`
}

// RecommendationStatus represents the current state of a recommendation.
type RecommendationStatus struct {
	// AdvancedHPA is the recommendation for advanced HPA.
	// +optional
	AdvancedHPA *AdvancedHorizontalPodAutoscalerRecommendation `json:"advancedHPA,omitempty"`

	// ResourceRequest is the recommendation for containers' cpu/mem requests.
	// +optional
	ResourceRequest *ResourceRequestRecommendation `json:"resourceRequest,omitempty"`

	// Conditions is an array of current recommendation conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// LastUpdateTime is last time we got an update on this status.
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// ValidityPeriod is the suggested validity period (e.g. 24h) of this recommendation since LastUpdateTime.
	// +optional
	ValidityPeriod metav1.Duration `json:"ValidityPeriod,omitempty"`
}

type AdvancedHorizontalPodAutoscalerRecommendation struct {
	// +optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`

	// +optional
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
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
	// Type is the analytics type, including HPA and resource.
	// +required
	// +kubebuilder:validation:Required
	Type Type `json:"type"`

	// ResourceSelector indicates how to select resources(e.g. a set of Deployments) for the analytics.
	// +required
	// +kubebuilder:validation:Required
	ResourceSelectors []ResourceSelector `json:"resourceSelectors"`

	// IntervalSeconds is the duration in seconds between two continuous analysis. Setting it to 0 means this is a one-off analysis.
	// +required
	// +kubebuilder:validation:Required
	IntervalSeconds *int64 `json:"intervalSeconds,omitempty"`
}

// AnalyticsStatus represents the current state of an analytics item.
type AnalyticsStatus struct {
	// LastAnalysisTime indicates the last time to perform analysis.
	// +optional
	LastAnalysisTime metav1.Time `json:"lastAnalysisTime,omitempty"`

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

// AnalyticsList is a list of analytics items.
type AnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Analytics `json:"items"`
}
