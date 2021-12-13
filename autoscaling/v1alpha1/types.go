package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	_ "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/metrics/pkg/apis/custom_metrics"
	_ "k8s.io/metrics/pkg/apis/custom_metrics/v1beta1"
	_ "k8s.io/metrics/pkg/apis/custom_metrics/v1beta2"
	_ "k8s.io/metrics/pkg/apis/metrics"
	_ "k8s.io/metrics/pkg/apis/metrics/v1beta1"

	predictionapi "github.com/gocrane/api/prediction/v1alpha1"
)

type ScaleStrategy string

const (
	// ScaleStrategyAuto execute scale based on metrics.
	ScaleStrategyAuto ScaleStrategy = "Auto"
	// ScaleStrategyPreview is the preview for ScaleStrategyAuto.
	ScaleStrategyPreview ScaleStrategy = "Preview"
)

// EffectiveHorizontalPodAutoscalerSpec defines the desired spec of EffectiveHorizontalPodAutoscaler
type EffectiveHorizontalPodAutoscalerSpec struct {
	// ScaleTargetRef is the reference to the workload that should be scaled.
	ScaleTargetRef autoscalingv2.CrossVersionObjectReference `json:"scaleTargetRef"`
	// MinReplicas is the lower limit replicas to the scale target which the autoscaler can scale down to.
	// the default MinReplicas is 1.
	// +optional
	// +kubebuilder:default=1
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// MaxReplicas is the upper limit replicas to the scale target which the autoscaler can scale up to.
	// It cannot be less that MinReplicas.
	MaxReplicas int32 `json:"maxReplicas"`
	// ScaleStrategy indicate the strategy to scaling target, value can be "Auto" and "Preview"
	// the default ScaleStrategy is Auto.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Auto;Preview
	// +kubebuilder:default=Auto
	ScaleStrategy ScaleStrategy `json:"scaleStrategy"`
	// SpecificReplicas specify the target replicas if ScaleStrategy is Preview
	// If not set, when ScaleStrategy is setting to Preview, it will just stop scaling
	// +optional
	// +kubebuilder:validation:Type=integer
	SpecificReplicas *int32 `json:"specificReplicas"`
	// metrics contains the specifications for which to use to calculate the
	// desired replica count (the maximum replica count across all metrics will
	// be used).  The desired replica count is calculated multiplying the
	// ratio between the target value and the current value by the current
	// number of pods.  Ergo, metrics used must decrease as the pod count is
	// increased, and vice-versa.  See the individual metric source types for
	// more information about how each type of metric must respond.
	// If not set, the default metric will be set to 80% average CPU utilization.
	// +optional
	Metrics []autoscalingv2.MetricSpec `json:"metrics,omitempty"`
	// behavior configures the scaling behavior of the target
	// in both Up and Down directions (scaleUp and scaleDown fields respectively).
	// If not set, the default HPAScalingRules for scale up and scale down are used.
	// +optional
	Behavior *autoscalingv2.HorizontalPodAutoscalerBehavior `json:"behavior,omitempty"`
	// Prediction defines configurations for predict resources.
	// If unspecified, defaults don't enable prediction.
	Prediction *Prediction `json:"prediction,omitempty"`
}

// Prediction defines configurations for predict resources
type Prediction struct {
	// PredictionWindowSeconds is the time window seconds to predict metrics in the future.
	// +optional
	// +kubebuilder:validation:Type=integer
	// +kubebuilder:default=3600
	PredictionWindowSeconds *int32 `json:"predictionWindowSeconds,omitempty"`
	// PredictionAlgorithm contains all algorithm config that provider by prediction api.
	// +optional
	PredictionAlgorithm *PredictionAlgorithm `json:"predictionAlgorithm,omitempty"`
}

// PredictionAlgorithm defines the algorithm to predict resources
type PredictionAlgorithm struct {
	// AlgorithmType specifies algorithm to predict resource
	AlgorithmType predictionapi.AlgorithmType `json:"algorithmType,omitempty"`
	// +optional
	DSP *predictionapi.DSP `json:"dsp,omitempty"`
	// +optional
	Percentile *predictionapi.Percentile `json:"percentile,omitempty"`
}

type ConditionType string

const (
	// PredictionReady indicates that the prediction is ready.
	// For enabled prediction effective-hpa.
	PredictionReady ConditionType = "PredictionReady"
	// Ready indicates that whole autoscaling is running as expect.
	Ready ConditionType = "Ready"
)

type EffectiveHorizontalPodAutoscalerStatus struct {
	// ExpectReplicas is the expected replicas to scale to.
	// +optional
	ExpectReplicas *int32 `json:"expectReplicas,omitempty"`
	// CurrentReplicas is the current replicas to the scale target.
	// +optional
	CurrentReplicas *int32 `json:"currentReplicas,omitempty"`
	// LastScaleTime indicates the last time to execute scaling.
	// +optional
	LastScaleTime *metav1.Time `json:"lastScaleTime,omitempty"`
	// Conditions is an array of current autoscaler conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ehpa
// +kubebuilder:printcolumn:name="STRATEGY",type="string",JSONPath=".spec.scaleStrategy",description="The scale strategy of ahpa."
// +kubebuilder:printcolumn:name="MINPODS",type="integer",JSONPath=".spec.minReplicas",description="The min replicas of target workload."
// +kubebuilder:printcolumn:name="MAXPODS",type="integer",JSONPath=".spec.maxReplicas",description="The max replicas of target workload."
// +kubebuilder:printcolumn:name="SPECIFICPODS",type="integer",JSONPath=".spec.specificReplicas",description="The specific replicas of target workload."
// +kubebuilder:printcolumn:name="REPLICAS",type="integer",JSONPath=".status.expectReplicas",description="The desired replicas of target workload."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// EffectiveHorizontalPodAutoscaler is the Schema for the effectivehorizontalpodautoscaler API
type EffectiveHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec EffectiveHorizontalPodAutoscalerSpec `json:"spec,omitempty"`

	// +optional
	Status EffectiveHorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EffectiveHorizontalPodAutoscalerList contains a list of EffectiveHorizontalPodAutoscaler
type EffectiveHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EffectiveHorizontalPodAutoscaler `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=.status.labelSelector
// +kubebuilder:resource:shortName=subs
// +kubebuilder:printcolumn:name="REPLICAS",type="integer",JSONPath=".status.replicas",description="The replicas presents the dryRun result."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// Substitute is the Schema for the Substitute API
type Substitute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec SubstituteSpec `json:"spec,omitempty"`

	// +optional
	Status SubstituteStatus `json:"status,omitempty"`
}

// SubstituteSpec defines the desired spec of Substitute
type SubstituteSpec struct {
	// Replicas is used by presents dryRun replicas for SubstituteTargetRef.
	// +optional
	Replicas *int32 `json:"replicas,omitempty"`

	// SubstituteTargetRef is the reference to the workload that should be Substituted.
	SubstituteTargetRef autoscalingv2.CrossVersionObjectReference `json:"substituteTargetRef"`
}

// SubstituteStatus defines the status of Substitute
type SubstituteStatus struct {
	// LabelSelector is label selectors that is sync with SubstituteTargetRef's labelSelector which used by HPA.
	// +optional
	LabelSelector string `json:"labelSelector,omitempty"`

	// Replicas is used by presents dryRun replicas for SubstituteTargetRef.
	// status.replicas should be equal to spec.replicas.
	Replicas int32 `json:"replicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubstituteList contains a list of Substitute
type SubstituteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Substitute `json:"items"`
}
