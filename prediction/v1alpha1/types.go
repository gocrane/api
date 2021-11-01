package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceName represents the name of the resource.
type ResourceName string

const (
	// ResourceCPU represents CPU in milli cores (1 core = 1000 milli cores).
	ResourceCPU ResourceName = "cpu"
	// ResourceMemory represents memory in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024).
	ResourceMemory ResourceName = "memory"
)

// PredictionMode represents the prediction time series mode.
type PredictionMode string

const (
	// PredictionModeInstant means predicting a single point in the future, for example the maximum value for the next hour
	PredictionModeInstant = "instant"
	// PredictionModeRange means predicting a time series during a range of time in the future.
	PredictionModeRange = "range"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodePrediction is the node prediction resource. it is associated with a node.
type NodePrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodePredictionResourceSpec `json:"spec"`

	// +optional
	Status NodePredictionResourceStatus `json:"status"`
}

// NodePredictionResourceSpec
type NodePredictionResourceSpec struct {
	// Period is the prediction time series interval or step.
	Period metav1.Duration `json:"period"`
	// Mode is the prediction time series mode
	Mode PredictionMode `json:"mode"`
	// MetricPredictionConfigs is the prediction configs of metric. each metric has its config for different prediction behaviors
	MetricPredictionConfigs []AlgorithmProviderConfig `json:"metricPredictionConfigs"`
}

// NodePredictionResourceStatus
type NodePredictionResourceStatus struct {
	// Consumed is the predicted resource usage in next resolution point based on past time series.
	Consumed Prediction `json:"consumed"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroupPrediction is a prediction on the resource consumed by a pod group.
// In kubernetes context, a pod group often refers to a batch of pods that satisfy a label selector.
type PodGroupPrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodGroupPredictionSpec `json:"spec"`

	// +optional
	Status PodGroupPredictionStatus `json:"status"`
}

type PredictionStatus string

const (
	// PredictionStatusCharging means no valid prediction series is available, just wait to predicting.
	PredictionStatusCharging PredictionStatus = "Charging"
	// PredictionStatusPredicting means the prediction routine is ongoing and the prediction data is valid.
	PredictionStatusPredicting PredictionStatus = "Predicting"
	// PredictionStatusNotStarted means the prediction routine has not started yet.
	PredictionStatusNotStarted PredictionStatus = "NotStarted"
	// PredictionStatusFinished means the prediction has finished.
	PredictionStatusFinished PredictionStatus = "Finished"
)

// PodGroupPredictionSpec is a description of a PodGroupPrediction.
type PodGroupPredictionSpec struct {
	// Prediction start time. If not specified, the prediction routine will start as soon as the CR is created.
	// +optional
	Start *metav1.Time `json:"start"`
	// Prediction end time, after which the prediction routine will stop, and the prediction data will not get updated any more.
	// If not specified, the prediction process will keep running forever.
	// +optional
	End *metav1.Time `json:"end"`
	// PredictionLength, for example, 24-hours means predicting time series in next 24 hours. This should be used only for PredictionModeRange.
	PredictionLength metav1.Duration `json:"predictionLength"`
	// Prediction mode
	Mode PredictionMode `json:"mode"`
	// Pods is a list of pod names that belong to this pod group. If not specified then WorkloadRef is invalid. the aggregator aggregate priority is  Pods > WorkloadRef > LabelSelector
	// +optional
	Pods []string `json:"pods"`
	// WorkloadRef is a ref of workload(deployment/statefulsets).
	// +optional
	WorkloadRef *autoscalingv2.CrossVersionObjectReference `json:"workloadRef"`
	// LabelSelector is the aggregator label selector. aggregator group all data by same key . for example, [online: label=v1] denotes all pods with label label=v1 will aggregate by sum all the resources.
	// +optional
	LabelSelector metav1.LabelSelector `json:"labelSelector"`
	// MetricPredictionConfigs is the prediction configs of metric. each metric has its config for different prediction behaviors
	MetricPredictionConfigs []AlgorithmProviderConfig `json:"metricPredictionConfigs"`
}

// PodGroupPredictionStatus
type PodGroupPredictionStatus struct {
	// Conditions is the condition of PodGroupPrediction
	Conditions []PodGroupPredictionCondition `json:"conditions,omitempty"`
	// Status
	Status PredictionStatus `json:"status,omitempty"`
	// Aggregation is the aggregated prediction value of all pods.
	Aggregation Prediction `json:"aggregation,omitempty"`
	// Containers is all the containers in pod group. excludes pause container. key is the namesapce/podname/containername
	Containers map[string]Prediction `json:"containers,omitempty"`
}

// PodGroupPredictionConditionType is a valid value for PodGroupPredictionCondition.Type
type PodGroupPredictionConditionType string

// These are valid conditions of PodGroupPrediction.
const (
	// PredictionConditionCharging means no valid prediction series is available, just wait to predict.
	PredictionConditionCharging PodGroupPredictionConditionType = "Charging"
	// PredictionConditionPredicting means the prediction routine is ongoing and the prediction data is valid.
	PredictionConditionPredicting PodGroupPredictionConditionType = "Predicting"
	// PredictionConditionNotStarted means the prediction routine has not started yet.
	PredictionConditionNotStarted PodGroupPredictionConditionType = "NotStarted"
	// PredictionStatusFinished means the prediction has finished, the prediction data will not be updated anymore.
	PredictionConditionFinished PodGroupPredictionConditionType = "Finished"
)

// PodGroupPredictionCondition contains details for the current condition of this pod.
type PodGroupPredictionCondition struct {
	// Type is the type of the condition.
	Type PodGroupPredictionConditionType `json:"type,omitempty"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status v1.ConditionStatus `json:"status,omitempty"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// Prediction define metrics prediction
type Prediction map[string]TimeSeries

// TimeSeries
type TimeSeries []*Vector

// Vector
type Vector struct {
	// CRD not support float64
	Value     string `json:"value"`
	Timestamp int64  `json:"timestamp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodePredictionList is a list of NodePrediction resources
type NodePredictionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NodePrediction `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodGroupPredictionList is a list of PodGroupPrediction
type PodGroupPredictionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PodGroupPrediction `json:"items"`
}

type AlgorithmProviderConfig struct {
	MetricName string `json:"metricName"`
	// +optional
	DSP *DspConfig `json:"dsp"`
	// +optional
	Percentile *PercentileConfig `json:"percentile"`
}

type DspConfig struct {
	// SampleInterval is the sampling interval of metrics.
	SampleInterval string `json:"sampleInterval"`
	// HistoryLength describes how long back should be queried against provider to get historical metrics for prediction.
	HistoryLength string `json:"historyLength"`
	// Estimators
	Estimators *EstimatorConfigs `json:"estimators"`
}

type EstimatorConfigs struct {
	// +optional
	MaxValue *MaxValueEstimatorConfig `json:"maxValue"`
	// +optional
	FFT *FFTEstimatorConfig `json:"fft"`
}

type MaxValueEstimatorConfig struct{}

type FFTEstimatorConfig struct {
	MarginFraction         string `json:"marginFraction"`
	LowAmplitudeThreshold  string `json:"lowAmplitudeThreshold"`
	HighFrequencyThreshold string `json:"highFrequencyThreshold"`
	MinNumOfSpectrumItems  int32  `json:"minNumOfSpectrumItems"`
	MaxNumOfSpectrumItems  int32  `json:"maxNumOfSpectrumItems"`
}

type PercentileConfig struct {
	SampleInterval  string          `json:"sampleInterval"`
	Histogram       HistogramConfig `json:"histogram"`
	MinSampleWeight string          `json:"minSampleWeight"`
}

type HistogramConfig struct {
	MaxValue              string `json:"maxValue"`
	Epsilon               string `json:"epsilon"`
	HalfLife              string `json:"halfLife"`
	BucketSize            string `json:"bucketSize"`
	FirstBucketSize       string `json:"firstBucketSize"`
	BucketSizeGrowthRatio string `json:"bucketSizeGrowthRatio"`
}
