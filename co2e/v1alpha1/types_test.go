package v1alpha1

import (
	"fmt"
	"testing"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

func TestCCF(t *testing.T) {
	carbonFootprint := CloudCarbonFootprint{
		TypeMeta: v1.TypeMeta{
			Kind:       "PowerUsageEffectiveness",
			APIVersion: "co2e.gocrane.io/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "carbonFootprint-shanghai-dc",
		},
		Spec: CloudCarbonFootprintSpec{
			Provider:       ProviderManual,
			Region:         "shanghai",
			Zone:           "china-shanghai-az01",
			Locality:       "AP/China/Shanghai/AZ01",
			PUE:            "1.5",
			EmissionFactor: "0.5810",
			ComputeConfig: []*ComputeConfig{
				{
					NodeSelector: &v1.LabelSelector{
						MatchLabels: map[string]string{"topology.kubernetes.io/zone": "shanghai-zone-01"},
					},
					MinWattsPerCPU:   "0.743",
					MaxWattsPerCPU:   "3.84",
					MemoryWattsPerGB: "0.65",
				},
			},
			StorageConfig: []*StorageConfig{
				{
					StorageClass: "cloudBlockDevice",
					WattsPerTB:   "1.2",
				},
			},
			NetworkingConfig: []*NetworkingConfig{
				{
					NetworkingClass: "golden",
					WattsPerGB:      "0.65",
				},
			},
		},
		Status: CloudCarbonFootprintStatus{
			Conditions: []v1.Condition{
				{
					Type: "Synced",
					LastTransitionTime: v1.Time{
						Time: time.Now(),
					},
					Reason:  "carbon footprint synced",
					Status:  "True",
					Message: "carbon footprint synced",
				},
			},
		},
	}
	cfStr, _ := json.Marshal(carbonFootprint)
	fmt.Println(string(cfStr))
}
