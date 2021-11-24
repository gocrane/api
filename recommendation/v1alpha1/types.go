package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RecommendationType string

const (
	RecommendationTypeAdvancedHPA RecommendationType = "AdvancedHPA"
	RecommendationTypeResource    RecommendationType = "Resource"
)

type AnalyticsType string

const (
	AnalyticsTypeAdvancedHPA AnalyticsType = "HPA"
	AnalyticsTypeResource    AnalyticsType = "Resource"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Recommendation represents the configuration of a single recommendation.
type Recommendation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec RecommendationSpec `json:"spec,omitempty"`
	Status RecommendationStatus `json:"status,omitempty"`
}

// RecommendationSpec describes the recommendation type and what the recommendation is for.
type RecommendationSpec struct {
	TargetRef autoscalingv2.CrossVersionObjectReference `json:"targetRef"`
	Namespace string                                    `json:"namespace"`
	Type      RecommendationType                        `json:"type"`
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
}

type AdvancedHorizontalPodAutoscalerRecommendation struct {
	// optional
	MinReplicas *int32
	// optional
	MaxReplicas *int32
}

type ResourceRequestRecommendation struct {
	Containers []ContainerRecommendation `json:"containers,omitempty"`
}

type ContainerRecommendation struct {
	ContainerName string              `json:"containerName"`
	Target        corev1.ResourceList `json:"target"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RecommendationList is a list of recommendations
type RecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Recommendation `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Analytics represents the configuration of an analytics.
type Analytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec AnalyticsSpec `json:"spec,omitempty"`
	Status AnalyticsStatus `json:"status,omitempty"`
}

// AnalyticsSpec describes the analytics type and what the analysis is for.
type AnalyticsSpec struct {
	// Type is the analytics type, including HPA and resource.
	Type AnalyticsType `json:"type"`
	// ResourceSelector indicates how to select resource(e.g. a set of Deployments) for the analytics.
	ResourceSelector []ResourceSelector `json:"resourceSelector"`
	// Interval is the time interval between two continuous analysis.
	Interval *metav1.Duration `json:"interval,omitempty"`
}

// AnalyticsStatus represents the current state of an analytics item.
type AnalyticsStatus struct {
	// LastAnalysisTime indicates the last time to perform analysis.
	// +optional
	LastAnalysisTime *metav1.Time `json:"lastAnalysisTime,omitempty"`
	// Conditions is an array of current analytics conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	// Recommendations is a list of pointers to recommendations that are updated by this analytics.
	// +optional
	// +listType=atomic
	Recommendations []corev1.ObjectReference `json:"active,omitempty"`
}

// ResourceSelector describes how the resources will be selected.
type ResourceSelector struct {
	// +optional
	APIGroup string `json:"apiGroup,omitempty"`
	// +optional
	Resource string `json:"resource,omitempty"`
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// +optional
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AnalyticsList is a list of analytics items.
type AnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Analytics `json:"items"`
}