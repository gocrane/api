package v1alpha1

import (
	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterNodePredictionList is a list of TimeSeriesPrediction resources
type ClusterNodePredictionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterNodePrediction `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=cnp
// +kubebuilder:subresource:status

// ClusterNodePrediction must be created in crane root namespace
// as TimeSeriesPrediction is a namespaced object now
type ClusterNodePrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterNodePredictionSpec   `json:"spec,omitempty"`
	Status ClusterNodePredictionStatus `json:"status,omitempty"`
}

type ClusterNodePredictionSpec struct {
	NodeSelector       map[string]string   `json:"nodeSelector,omitempty"`
	PredictionTemplate *PredictionTemplate `json:"template,omitempty"`
}

type ClusterNodePredictionStatus struct {
	DesiredNumberCreated int                `json:"desiredNumberCreated,omitempty"`
	CurrentNumberCreated int                `json:"currentNumberCreated,omitempty"`
	Conditions           []metav1.Condition `json:"conditions,omitempty"`
}

type PredictionTemplate struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TimeSeriesPredictionSpec `json:"spec,omitempty"`
}

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
	DSP *DSP `json:"dsp,omitempty"`
	// +optional
	Percentile *Percentile `json:"percentile,omitempty"`
}

type DSP struct {
	// SampleInterval is the sampling interval of metrics.
	SampleInterval string `json:"sampleInterval,omitempty"`
	// HistoryLength describes how long back should be queried against provider to get historical metrics for prediction.
	HistoryLength string `json:"historyLength,omitempty"`
	// Estimator
	Estimators Estimators `json:"estimators,omitempty"`
}

type Estimators struct {
	// +optional
	MaxValueEstimators []*MaxValueEstimator `json:"maxValue,omitempty"`
	// +optional
	FFTEstimators []*FFTEstimator `json:"fft,omitempty"`
}

type MaxValueEstimator struct {
	MarginFraction string `json:"marginFraction,omitempty"`
}

type FFTEstimator struct {
	MarginFraction         string `json:"marginFraction,omitempty"`
	LowAmplitudeThreshold  string `json:"lowAmplitudeThreshold,omitempty"`
	HighFrequencyThreshold string `json:"highFrequencyThreshold,omitempty"`
	MinNumOfSpectrumItems  *int32 `json:"minNumOfSpectrumItems,omitempty"`
	MaxNumOfSpectrumItems  *int32 `json:"maxNumOfSpectrumItems,omitempty"`
}

type Percentile struct {
	SampleInterval  string          `json:"sampleInterval,omitempty"`
	Histogram       HistogramConfig `json:"histogram,omitempty"`
	MinSampleWeight string          `json:"minSampleWeight,omitempty"`
	MarginFraction  string          `json:"marginFraction,omitempty"`
	Percentile      string          `json:"percentile,omitempty"`
}

type HistogramConfig struct {
	MaxValue              string `json:"maxValue,omitempty"`
	Epsilon               string `json:"epsilon,omitempty"`
	HalfLife              string `json:"halfLife,omitempty"`
	BucketSize            string `json:"bucketSize,omitempty"`
	FirstBucketSize       string `json:"firstBucketSize,omitempty"`
	BucketSizeGrowthRatio string `json:"bucketSizeGrowthRatio,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=tsp
// +kubebuilder:printcolumn:name="TargetRefName",type="string",JSONPath=".spec.targetRef.name",description="The target ref name of tsp."
// +kubebuilder:printcolumn:name="TargetRefKind",type="string",JSONPath=".spec.targetRef.kind",description="The target ref kind of tsp."
// +kubebuilder:printcolumn:name="PredictionWindowSeconds",type="integer",JSONPath=".spec.predictionWindowSeconds",description="The predictionWindowSeconds of tsp."
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created."
// +kubebuilder:webhooks:path=/mutate-timeseriesprediction,mutating=true,failurePolicy=fail,groups=prediction.crane.io,resources=timeseriesprediction,verbs=create;update,versions=v1alpha1,name=timeseriespredictions.webhook.prediction.crane.io,sideEffects=none,admissionReviewVersions=v1
// +kubebuilder:webhooks:verbs=create;update,path=/validate-timeseriesprediction,mutating=false,failurePolicy=fail,groups=prediction.crane.io,resources=timeseriesprediction,versions=v1,name=timeseriespredictions.webhook.prediction.crane.io,sideEffects=none,admissionReviewVersions=v1

// TimeSeriesPrediction is a prediction for a time series.
type TimeSeriesPrediction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TimeSeriesPredictionSpec `json:"spec,omitempty"`

	// +optional
	Status TimeSeriesPredictionStatus `json:"status,omitempty"`
}

// TimeSeriesPredictionSpec is a description of a TimeSeriesPrediction.
type TimeSeriesPredictionSpec struct {
	// PredictionMetrics is an array of PredictionMetric
	PredictionMetrics []PredictionMetric `json:"predictionMetrics,omitempty"`
	// Target is the target referent of time series prediction. each TimeSeriesPrediction associate with just only one target ref.
	// all metrics in PredictionMetricConfigurations is about the TargetRef
	TargetRef v1.ObjectReference `json:"targetRef,omitempty"`
	// PredictionWindowSeconds is a time window in seconds, indicating how long to predict in the future.
	PredictionWindowSeconds int32 `json:"predictionWindowSeconds,omitempty"`
}

// TimeSeriesPredictionStatus is the status of a TimeSeriesPrediction.
type TimeSeriesPredictionStatus struct {
	// PredictionMetrics is an array of PredictionMetricStatus of all PredictionMetrics, PredictionMetricStatus include predicted time series data
	PredictionMetrics []PredictionMetricStatus `json:"predictionMetrics,omitempty"`
	// Conditions is the condition of TimeSeriesPrediction
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// PredictionConditionType is a valid value for TimeSeriesPredictionCondition.Type
type PredictionConditionType string

// These are valid conditions of TimeSeriesPrediction.
const (
	// TimeSeriesPredictionConditionReady means the prediction data is available to consume
	TimeSeriesPredictionConditionReady PredictionConditionType = "Ready"
)

// PredictionMetric describe what metric of your time series prediction, how to query, use which algorithm to predict.
type PredictionMetric struct {
	// ResourceIdentifier is a resource to identify the metric, but now it is just an identifier now. reference otlp https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/resource/sdk.md
	ResourceIdentifier string `json:"resourceIdentifier,omitempty"`
	// Type is the type of metric, now support ResourceQuery、ExpressionQuery、RawQuery
	Type MetricType `json:"type,omitempty"`
	// +optional
	// ResourceQuery is a kubernetes built in metric, only support cpu, memory
	ResourceQuery *v1.ResourceName `json:"resourceQuery,omitempty"`
	// following QueryExpressions depend on your crane system data source configured when the system start.
	// if you use different sources with your system start params, it is not valid.
	// +optional
	// ExpressionQuery is a query expression of non-prometheus style, usually is api style
	ExpressionQuery *ExpressionQuery `json:"expressionQuery,omitempty"`
	// +optional
	// RawQuery is a query expression of DSL style, such as prometheus query language
	RawQuery *RawQuery `json:"rawQuery,omitempty"`
	// Algorithm is the algorithm used by this prediction metric.
	Algorithm Algorithm `json:"algorithm,omitempty"`
}

// MetricType is the type of metric
type MetricType string

const (
	// ResourceQueryMetricType is kubernetes built in metric, only support cpu and memory now.
	ResourceQueryMetricType MetricType = "ResourceQuery"
	// ExpressionQueryMetricType is an selector style metric, it queried from a system which supports it.
	ExpressionQueryMetricType MetricType = "ExpressionQuery"
	// RawQueryMetricType is an raw query style metric, it is queried from a system which supports it, such as prometheus
	RawQueryMetricType MetricType = "RawQuery"
)

// ExpressionQuery
type ExpressionQuery struct {
	// MetricName is the name of the metric.
	// +required
	// +kubebuilder:validation:Required
	MetricName string `json:"metricName,omitempty"`

	// QueryConditions is a query condition list.
	// + optional
	QueryConditions []QueryCondition `json:"labels,omitempty"`
}

// RawQuery
type RawQuery struct {
	// Expression is the query expression. For prometheus, it is promQL.
	Expression string `json:"expression,omitempty"`
}

// QueryCondition is a key, operator, value triple.
// E.g. 'namespace = default', 'role in [Admin, Developer]'
type QueryCondition struct {
	// Key is the key of the query condition
	Key string `json:"key,omitempty"`
	// Operator
	Operator Operator `json:"operator,omitempty"`
	// Value is the query value list.
	Value []string `json:"value,omitempty"`
}

type Operator string

const (
	OperatorEqual         Operator = "="
	OperatorNotEqual      Operator = "!="
	OperatorRegexMatch    Operator = "=~"
	OperatorNotRegexMatch Operator = "!~"
	OperatorIn            Operator = "in"
)

// Algorithm describe the algorithm params
type Algorithm struct {
	// AlgorithmType is the algorithm type, currently supports dsp and percentile.
	AlgorithmType AlgorithmType `json:"algorithmType,omitempty"`
	// +optional
	// DSP is an algorithm which use FFT to deal with time series, typically it is used to predict some periodic time series
	DSP *DSP `json:"dsp,omitempty"`
	// +optional
	// Percentile is an algorithm which use exponential time decay histogram, it can predict a reasonable value according your history time series
	Percentile *Percentile `json:"percentile,omitempty"`
}

type MetricTimeSeriesList []*MetricTimeSeries

// MetricPredictedData is predicted data of an metric, which denote a metric by ResourceIdentifier in the PredictionMetric
type PredictionMetricStatus struct {
	// ResourceIdentifier is a resource to identify the metric, but now it is just an identifier now.
	// such as cpu, memory
	ResourceIdentifier string `json:"resourceIdentifier,omitempty"`
	// Prediction is the predicted time series data of the metric
	Prediction []*MetricTimeSeries `json:"prediction,omitempty"`
}

// MetricTimeSeries is a stream of samples that belong to a metric with a set of labels
type MetricTimeSeries struct {
	// A collection of Labels that are attached by monitoring system as metadata
	// for the metrics, which are known as dimensions.
	Labels []Label `json:"labels,omitempty"`
	// A collection of Samples in chronological order.
	Samples []Sample `json:"samples,omitempty"`
}

// Sample pairs a Value with a Timestamp.
type Sample struct {
	Value     string `json:"value,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

// A Label is a Name and Value pair that provides additional information about the metric.
// It is metadata for the metric. For example, Kubernetes pod metrics always have
// 'namespace' label that represents which namespace it belongs to.
type Label struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TimeSeriesPredictionList is a list of NodePrediction resources
type TimeSeriesPredictionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TimeSeriesPrediction `json:"items"`
}
