package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ProviderManual = "Manual"
)

type ComputeConfig struct {
	// when there are various types of compute server, node selector selects the targets
	NodeSelector *metav1.LabelSelector `json:"nodeSelector,omitempty"`
	// power when cpu is idle
	MinWattsPerCPU string `json:"minWattsPerCPU,omitempty"`
	// power when cpu utilization is 100%
	MaxWattsPerCPU string `json:"maxWattsPerCPU,omitempty"`
	// sometimes it's hard to measure memory, storage, and networking energy consumption
	// CPUEnergyConsumptionRatio can be defined to specify the percentage of cpu energy consumption vs all IT equipments consumption
	CPUEnergyConsumptionRatio string `json:"cpuEnergyConsumptionRatio,omitempty"`
	// power of per BG memory
	MemoryWattsPerGB string `json:"memoryWattsPerGB,omitempty"`
}

type StorageConfig struct {
	// storage class, e.g. cephfs
	StorageClass string `json:"storageClass,omitempty"`
	// power per TB for the class
	WattsPerTB string `json:"wattsPerTB,omitempty"`
}

type NetworkingConfig struct {
	// networking class, e.g. golden, server, bronze, which define different redundancies of networking links, and has different energy consumption
	NetworkingClass string `json:"storageClass,omitempty"`
	// power per GB for the class
	WattsPerGB string `json:"wattsPerGB,omitempty"`
}

type CloudCarbonFootprintSpec struct {
	// Provider is the provider of the ccf, when provider is manual, all the properties of ccf would be configured manually
	// when a cloud provider exposes query API, a cloud provider controller can query cloud api and fill the properties automatically
	Provider string `json:"provider,omitempty"`
	// region of the datacenter, e.g. shanghai
	Region string `json:"region,omitempty"`
	// availability zone of the datacenter, e.g. shanghai-az01
	Zone string `json:"zone,omitempty"`
	// locality holds more information of location, e.g. ap/china/shanghai/az01/floor3
	Locality string `json:"locality,omitempty"`
	// power usage effectiveness = IT equipment energy usage / total facility energy usage
	PUE string `json:"pue,omitempty"`
	// emission factor of the data center, unit is tCO2/MWh, the average emission factor of China is 0.5810
	EmissionFactor string `json:"emissionFactor,omitempty"`
	// compute power infos
	// when there are multiple node types in the cluster, define multiple compute configs
	ComputeConfig []*ComputeConfig `json:"computeConfig,omitempty"`
	// storage power info
	StorageConfig []*StorageConfig `json:"storageConfig,omitempty"`
	// networking power info
	NetworkingConfig []*NetworkingConfig `json:"networkingConfig,omitempty"`
}

type CloudCarbonFootprintStatus struct {
	Conditions []metav1.Condition `json:"condition,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:resource:scope=Cluster,shortName=ccf,path=cloudcarbonfootprints
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudCarbonFootprint defines carbon footprint configuration of a datacenter
type CloudCarbonFootprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudCarbonFootprintSpec   `json:"spec,omitempty"`
	Status CloudCarbonFootprintStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudCarbonFootprintList contains a list of PodQOS
type CloudCarbonFootprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CloudCarbonFootprint `json:"items"`
}
