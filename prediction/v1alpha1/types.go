package v1alpha1

import (
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
	// Type is the type of metric, now support ResourceQuery、MetricQuery、ExpressionQuery
	Type MetricType `json:"type,omitempty"`
	// +optional
	// ResourceQuery is a kubernetes built in metric, only support cpu, memory
	ResourceQuery *v1.ResourceName `json:"resourceQuery,omitempty"`
	// following QueryExpressions depend on your crane system data source configured when the system start.
	// if you use different sources with your system start params, it is not valid.
	// +optional
	// MetricQuery is a query against a metric with a set of conditions
	MetricQuery *MetricQuery `json:"metricQuery,omitempty"`
	// +optional
	// ExpressionQuery is a query with a DSL-style expression, such as prometheus promQL
	ExpressionQuery *ExpressionQuery `json:"expressionQuery,omitempty"`
	// Algorithm is the algorithm used by this prediction metric.
	Algorithm Algorithm `json:"algorithm,omitempty"`
}

// MetricType is the type of metric
type MetricType string

const (
	// ResourceQueryMetricType is kubernetes built in metric, only support cpu and memory now.
	ResourceQueryMetricType MetricType = "ResourceQuery"
	// MetricQueryMetricType is an selector style metric, it queried from a system which supports it.
	MetricQueryMetricType MetricType = "MetricQuery"
	// ExpressionQueryMetricType is an raw query style metric, it is queried from a system which supports it, such as prometheus
	ExpressionQueryMetricType MetricType = "ExpressionQuery"
)

// MetricQuery
type MetricQuery struct {
	// MetricName is the name of the metric.
	// +required
	// +kubebuilder:validation:Required
	MetricName string `json:"metricName,omitempty"`

	// QueryConditions is a query condition list.
	// + optional
	QueryConditions []QueryCondition `json:"labels,omitempty"`
}

// ExpressionQuery
type ExpressionQuery struct {
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
