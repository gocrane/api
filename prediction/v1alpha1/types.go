package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
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

type AlgorithmType string

const (
	AlgorithmTypePercentile AlgorithmType = "percentile"
	AlgorithmTypeDSP        AlgorithmType = "dsp"
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
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodePrediction is the node prediction resource, which is associated with a node.
type NodePrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NodePredictionResourceSpec `json:"spec"`

	// +optional
	Status NodePredictionResourceStatus `json:"status"`
}

// NodePredictionResourceSpec is the specification of a node prediction.
type NodePredictionResourceSpec struct {
	// Period is the prediction time series interval or step.
	Period metav1.Duration `json:"period,omitempty"`
	// Mode is the prediction time series mode
	Mode PredictionMode `json:"mode,omitempty"`
	// MetricPredictionConfigs is the prediction configs of metric. each metric has its config for different prediction behaviors
	MetricPredictionConfigs []MetricPredictionConfig `json:"metricPredictionConfigs,omitempty"`
}

// NodePredictionResourceStatus represents information about the status of NodePrediction
type NodePredictionResourceStatus struct {
	// NextPossible is the predicted resource usage in next resolution point based on previous series.
	NextPossible Prediction `json:"nextPossible,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:webhooks:path=/mutate-podgroupprediction,mutating=true,failurePolicy=fail,groups=prediction.crane.io,resources=podgrouppredictions,verbs=create;update,versions=v1alpha1,name=prediction.crane.io_podgrouppredictions_webhook,sideEffects=none,admissionReviewVersions=v1
// +kubebuilder:webhooks:verbs=create;update,path=/validate-podgroupprediction,mutating=false,failurePolicy=fail,groups=prediction.crane.io,resources=podgrouppredictions,versions=v1,name=prediction.crane.io_podgrouppredictions_webhook,sideEffects=none,admissionReviewVersions=v1
// +kubebuilder:subresource:status
// +kubebuilder:object:root=true

// PodGroupPrediction is a prediction on the resource consumed by a pod group.
// In kubernetes context, a pod group often refers to a batch of pods that satisfy a label selector.
type PodGroupPrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PodGroupPredictionSpec `json:"spec,omitempty"`

	// +optional
	Status PodGroupPredictionStatus `json:"status,omitempty"`
}

type PredictionStatus string

const (
	// PredictionStatusPending - no valid prediction series available, wait for prediction.
	PredictionStatusPending PredictionStatus = "Pending"
	// PredictionStatusPredicting - prediction is on the way, result is ready and value is valid.
	PredictionStatusPredicting PredictionStatus = "Predicting"
	// PredictionStatusNotStarted - the prediction has not start.
	PredictionStatusNotStarted PredictionStatus = "NotStarted"
	// PredictionStatusCompleted - the prediction has competed.
	PredictionStatusCompleted PredictionStatus = "Completed"
)

// PodGroupPredictionSpec is a description of a PodGroupPrediction.
type PodGroupPredictionSpec struct {
	// Prediction start time. If not specified, the prediction starts from the object creationTimestamp.
	// +optional
	Start *metav1.Time `json:"start,omitempty"`
	// Prediction end time. If current time is after end, the prediction will be stopped and the status will not be updated afterward.
	// If end is null, the prediction will never stop.
	// +optional
	End *metav1.Time `json:"end,omitempty"`
	// PredictionWindow, for example, 24-hours means predicting time series in next 24 hours.
	// This should be used only for PredictionModeRange.
	PredictionWindow metav1.Duration `json:"predictionWindow,omitempty"`
	// Mode is the prediction time series mode. instant or range
	Mode PredictionMode `json:"mode,omitempty"`
	// Pods is a list of pod names that belong to this pod group.
	// If not specified then WorkloadRef is invalid.
	// The aggregator aggregate priority is  Pods > WorkloadRef > LabelSelector
	// +optional
	Pods []string `json:"pods,omitempty"`
	// WorkloadRef is a ref of workload(deployment/statefulsets).
	// +optional
	WorkloadRef *autoscalingv2.CrossVersionObjectReference `json:"workloadRef,omitempty"`
	// LabelSelector is the aggregator label selector. aggregator group all data by same key . for example, [online: label=v1] denotes all pods with label label=v1 will aggregate by sum all the resources.
	// +optional
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`
	// MetricPredictionConfigs is the prediction configs of metric. each metric has its config for different prediction behaviors
	MetricPredictionConfigs []MetricPredictionConfig `json:"metricPredictionConfigs,omitempty"`
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
	Status metav1.ConditionStatus `json:"status,omitempty"`
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
	Value     string `json:"value,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
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

type MetricPredictionConfig struct {
	MetricName    string        `json:"metricName,omitempty"`
	AlgorithmType AlgorithmType `json:"algorithmType,omitempty"`
	// +optional
	DSP *Dsp `json:"dsp,omitempty"`
	// +optional
	Percentile *Percentile `json:"percentile,omitempty"`
}

type Dsp struct {
	// SampleInterval is the sampling interval of metrics.
	SampleInterval string `json:"sampleInterval,omitempty"`
	// HistoryLength describes how long back should be queried against provider to get historical metrics for prediction.
	HistoryLength string `json:"historyLength,omitempty"`
	// Estimator
	Estimator Estimator `json:"estimators,omitempty"`
}

type Estimator struct {
	// +optional
	MaxValueEstimators []*MaxValueEstimator `json:"maxValue,omitempty"`
	// +optional
	FFTEstimators []*FFTEstimator `json:"fft,omitempty"`
}

type MaxValueEstimator struct{}

type FFTEstimator struct {
	MarginFraction         string `json:"marginFraction,omitempty"`
	LowAmplitudeThreshold  string `json:"lowAmplitudeThreshold,omitempty"`
	HighFrequencyThreshold string `json:"highFrequencyThreshold,omitempty"`
	MinNumOfSpectrumItems  int32  `json:"minNumOfSpectrumItems,omitempty"`
	MaxNumOfSpectrumItems  int32  `json:"maxNumOfSpectrumItems,omitempty"`
}

type Percentile struct {
	SampleInterval  string          `json:"sampleInterval,omitempty"`
	Histogram       HistogramConfig `json:"histogram,omitempty"`
	MinSampleWeight string          `json:"minSampleWeight,omitempty"`
}

type HistogramConfig struct {
	MaxValue              string `json:"maxValue,omitempty"`
	Epsilon               string `json:"epsilon,omitempty"`
	HalfLife              string `json:"halfLife,omitempty"`
	BucketSize            string `json:"bucketSize,omitempty"`
	FirstBucketSize       string `json:"firstBucketSize,omitempty"`
	BucketSizeGrowthRatio string `json:"bucketSizeGrowthRatio,omitempty"`
}
