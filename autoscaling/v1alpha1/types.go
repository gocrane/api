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
	// ScaleStrategyManual execute scale manually.
	ScaleStrategyManual ScaleStrategy = "Manual"
)

// AdvancedHorizontalPodAutoscalerSpec defines the desired spec of AdvancedHorizontalPodAutoscaler
type AdvancedHorizontalPodAutoscalerSpec struct {
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
	// ScaleStrategy indicate the strategy to scaling target, value can be "Auto" and "Manual"
	// the default ScaleStrategy is Auto.
	// +optional
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Enum=Auto;Manual
	// +kubebuilder:default=Auto
	ScaleStrategy ScaleStrategy `json:"scaleStrategy"`
	// SpecificReplicas specify the target replicas if ScaleStrategy is Manual
	// If not set, when ScaleStrategy is setting to Manual, it will just stop scaling
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
	DSP *predictionapi.Dsp `json:"dsp,omitempty"`
	// +optional
	Percentile *predictionapi.Percentile `json:"percentile,omitempty"`
}

type ConditionType string

const (
	// PredictionReady indicates that the prediction is ready.
	// For enabled prediction advanced-hpa.
	PredictionReady ConditionType = "PredictionReady"
	// Ready indicates that whole autoscaling is running as expect.
	Ready ConditionType = "Ready"
)

type AdvancedHorizontalPodAutoscalerStatus struct {
	// ExpectReplicas is the expected replicas to scale to.
	// +optional
	ExpectReplicas *int32 `json:"expectReplicas,omitempty"`
	// CurrentReplicas is the current replicas to the scale target.
	// +optional
	CurrentReplicas *int32 `json:"currentReplicas,omitempty"`
	// LastScaleTime indicate the last time to execute scaling.
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
// +kubebuilder:resource:shortName=ahpa
// +kubebuilder:printcolumn:name="STRATEGY",type="string",JSONPath=".spec.scaleStrategy",description="The scale strategy of ahpa."
// +kubebuilder:printcolumn:name="MINPODS",type="integer",JSONPath=".spec.minReplicas",description="The min replicas of target workload."
// +kubebuilder:printcolumn:name="MAXPODS",type="integer",JSONPath=".spec.maxReplicas",description="The max replicas of target workload."
// +kubebuilder:printcolumn:name="SPECIFICPODS",type="integer",JSONPath=".spec.specificReplicas",description="The specific replicas of target workload."
// +kubebuilder:printcolumn:name="REPLICAS",type="integer",JSONPath=".status.expectReplicas",description="The desired replicas of target workload."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."

// AdvancedHorizontalPodAutoscaler is the Schema for the advancedhorizontalpodautoscaler API
type AdvancedHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec AdvancedHorizontalPodAutoscalerSpec `json:"spec,omitempty"`

	// +optional
	Status AdvancedHorizontalPodAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AdvancedHorizontalPodAutoscalerList contains a list of AdvancedHorizontalPodAutoscaler
type AdvancedHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AdvancedHorizontalPodAutoscaler `json:"items"`
}
