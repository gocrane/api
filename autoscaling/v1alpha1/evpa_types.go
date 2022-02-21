package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vpatypes "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
)

type EffectiveVerticalPodAutoscalerSpec struct {

	// TargetRef points to the controller managing the set of pods for the
	// autoscaler to control - e.g. Deployment, StatefulSet. VerticalPodAutoscaler
	// can be targeted at controller implementing scale subresource (the pod set is
	// retrieved from the controller's ScaleStatus) or some well known controllers
	// (e.g. for DaemonSet the pod set is read from the controller's spec).
	// If VerticalPodAutoscaler cannot use specified target it will report
	// ConfigUnsupported condition.
	// Note that VerticalPodAutoscaler does not require full implementation
	// of scale subresource - it will not use it to modify the replica count.
	// The only thing retrieved is a label selector matching pods grouped by
	// the target resource.
	TargetRef *autoscalingv2.CrossVersionObjectReference `json:"targetRef"`

	// Describes the rules on how changes are applied to the pods.
	// If not specified, all fields in the `PodUpdatePolicy` are set to their
	// default values.
	// +optional
	UpdatePolicy *vpatypes.PodUpdatePolicy `json:"updatePolicy,omitempty"`

	// Controls how the autoscaler computes recommended resources.
	// The resource policy may be used to set constraints on the recommendations
	// for individual containers. If not specified, the autoscaler computes recommended
	// resources for all containers in the pod, without additional constraints.
	// +optional
	ResourcePolicy *PodResourcePolicy `json:"resourcePolicy,omitempty"`

	// ResourceEstimators Contains the specifications for estimators.
	// +optional
	ResourceEstimators []ResourceEstimator `json:"resourceEstimators,omitempty"`
}

// PodResourcePolicy controls how autoscaler computes the recommended resources
// for containers belonging to the pod. There can be at most one entry for every
// named container and optionally a single wildcard entry with `containerName` = '*',
// which handles all containers that don't have individual policies.
type PodResourcePolicy struct {
	// Per-container resource policies.
	// +optional
	// +patchMergeKey=containerName
	// +patchStrategy=merge
	ContainerPolicies []ContainerResourcePolicy `json:"containerPolicies,omitempty" patchStrategy:"merge" patchMergeKey:"containerName"`
}

// ContainerResourcePolicy controls how autoscaler computes the recommended
// resources for a specific container.
type ContainerResourcePolicy struct {
	// Name of the container or DefaultContainerResourcePolicy, in which
	// case the policy is used by the containers that don't have their own
	// policy specified.
	ContainerName string `json:"containerName,omitempty"`
	// ScaleUpPolicy define the policy when scale up containers resources.
	// +optional
	ScaleUpPolicy *ContainerScalingPolicy `json:"scaleUpPolicy,omitempty"`
	// ScaleDownPolicy define the policy when scale down containers resources.
	// +optional
	ScaleDownPolicy *ContainerScalingPolicy `json:"scaleDownPolicy,omitempty"`
	// Specifies the minimal amount of resources that will be recommended
	// for the container. The default is no minimum.
	// +optional
	MinAllowed v1.ResourceList `json:"minAllowed,omitempty"`
	// Specifies the maximum amount of resources that will be recommended
	// for the container. The default is no maximum.
	// +optional
	MaxAllowed v1.ResourceList `json:"maxAllowed,omitempty"`

	// Specifies the type of recommendations that will be computed
	// (and possibly applied) by VPA.
	// If not specified, the default of [ResourceCPU, ResourceMemory] will be used.
	ControlledResources *[]ResourceName `json:"controlledResources,omitempty" patchStrategy:"merge"`

	// Specifies which resource values should be controlled.
	// The default is "RequestsAndLimits".
	// +optional
	ControlledValues *vpatypes.ContainerControlledValues `json:"controlledValues,omitempty"`
}

// ContainerScalingPolicy define effective policy for vertical scaling in certain direction: Up or Down
type ContainerScalingPolicy struct {
	// ScaleMode controls Whether autoscaler is enabled for the container.
	// The default is "Auto".
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Auto;Off
	// +kubebuilder:default=Auto
	ScaleMode *vpatypes.ContainerScalingMode `json:"mode,omitempty"`

	// MetricThresholds defines resource usages thresholds for vertical scaling.
	// Only if actual usage is reached to threshold, autoscaling estimator will be triggered.
	MetricThresholds *ResourceMetricList `json:"metricThresholds,omitempty"`

	// StabilizationWindowSeconds is the number of seconds for which past vertical scaling
	// considered while scaling up or scaling down.
	// +optional
	// +kubebuilder:validation:Type=integer
	// +kubebuilder:default=3600
	StabilizationWindowSeconds *int32 `json:"stabilizationWindowSeconds,omitempty"`
}

type ResourceName string

type ResourceMetricList map[ResourceName]ResourceMetric

type ResourceMetric struct {
	// averageValue is the target value of the average of the
	// metric across all relevant pods (as a quantity)
	// +optional
	AverageValue *resource.Quantity `json:"averageValue,omitempty"`
}

// ResourceEstimator defines the spec for resource estimator
type ResourceEstimator struct {
	// Type defines the type for this estimator.
	// +optional
	Type string `json:"type,omitempty"`

	// Priority defines the priority for this estimator.
	// +optional
	Priority int `json:"priority,omitempty"`

	// Config contains key value pairs for this estimator
	// for example, we can define configs like:
	// key1: value1
	// key2: value2
	// these configs will pass into estimators when execute scaling
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

// EffectiveVerticalPodAutoscalerStatus describes the runtime state of the autoscaler.
type EffectiveVerticalPodAutoscalerStatus struct {
	// CurrentEstimators is the last state of the estimators used by this autoscaler
	CurrentEstimators []ResourceEstimatorStatus `json:"currentEstimators,omitempty"`

	// The most recently computed amount of resources recommended by the
	// autoscaler for the controlled pods.
	// +optional
	Recommendation *vpatypes.RecommendedPodResources `json:"recommendation,omitempty"`

	// Conditions is the set of conditions required for this autoscaler to scale its target,
	// and indicates whether or not those conditions are met.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []vpatypes.VerticalPodAutoscalerCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// ResourceEstimatorStatus contains state for resource estimator
type ResourceEstimatorStatus struct {
	// Type defines the type for this estimator.
	// +optional
	Type string `json:"type,omitempty"`

	// LastUpdateTime is the last time the status updated.
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// The most recently computed amount of resources recommended by the
	// estimator for the controlled pods.
	// +optional
	Recommendation *vpatypes.RecommendedPodResources `json:"recommendation,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=evpa
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// EffectiveVerticalPodAutoscaler is the Schema for the effectiveverticalpodautoscaler API
type EffectiveVerticalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec EffectiveVerticalPodAutoscalerSpec `json:"spec"`

	// +optional
	Status EffectiveVerticalPodAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EffectiveVerticalPodAutoscalerList contains a list of EffectiveVerticalPodAutoscaler
type EffectiveVerticalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EffectiveVerticalPodAutoscaler `json:"items"`
}
