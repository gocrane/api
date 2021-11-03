package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	predictionv1alpha1 "github.com/gocrane-io/api/prediction/v1alpha1"
)

// AdvancedHorizontalPodAutoscalerSpec defines the desired spec of AdvancedHorizontalPodAutoscaler
type AdvancedHorizontalPodAutoscalerSpec struct {
	// ScaleTargetRef is the reference to the workload that should be scaled.
	ScaleTargetRef autoscalingv2.CrossVersionObjectReference `json:"scaleTargetRef"`
	// MinReplicaCount is the lower limit replicas to the scale target which the autoscaler can scale down to.
	// the default MinReplicaCount is 1.
	// +optional
	MinReplicaCount *int32 `json:"minReplicaCount,omitempty"`
	// MaxReplicaCount is the upper limit replicas to the scale target which the autoscaler can scale up to.
	// It cannot be less that MinReplicaCount.
	MaxReplicaCount int32 `json:"maxReplicaCount"`
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
	// PredictionConfig defines config for predict resources.
	// If unspecified, defaults don't enable prediction.
	PredictionConfig PredictionConfig `json:"predictionConfig,omitempty"`
}

// PredictionConfig defines config for predict resources
type PredictionConfig struct {
	// PredictionWindow is the time window to predict metrics in the future.
	PredictionWindow *int32 `json:"predictionWindow,omitempty"`
	// PredictionAlgorithm contains all algorithm config that provider by prediction api.
	// +optional
	PredictionAlgorithm *PredictionAlgorithm `json:"predictionAlgorithm,omitempty"`
}

// PredictionAlgorithm defines the algorithm to predict resources
type PredictionAlgorithm struct {
	// +optional
	DSP *predictionv1alpha1.DspConfig `json:"dsp,omitempty"`
	// +optional
	Percentile *predictionv1alpha1.PercentileConfig `json:"percentile,omitempty"`
}

type AdvancedHorizontalPodAutoscalerStatus struct {
	// ExpectReplicaCount is the expected replica count to scale to.
	// +optional
	ExpectReplicaCount *int32 `json:"expectReplicaCount,omitempty"`
	// CurrentReplicaCount is the current replica count to the scale target.
	// +optional
	CurrentReplicaCount *int32 `json:"currentReplicaCount,omitempty"`
	// LastScaleTime indicate the last time to execute scaling.
	// +optional
	LastScaleTime *metav1.Time `json:"lastScaleTime,omitempty"`
	// Conditions is an array of current autoscaler conditions.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AdvancedHorizontalPodAutoscaler is the Schema for the advancedhorizontalpodautoscaler API
type AdvancedHorizontalPodAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec AdvancedHorizontalPodAutoscalerSpec `json:"spec"`

	Status AdvancedHorizontalPodAutoscalerStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AdvancedHorizontalPodAutoscalerList contains a list of AdvancedHorizontalPodAutoscaler
type AdvancedHorizontalPodAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AdvancedHorizontalPodAutoscaler `json:"items"`
}
