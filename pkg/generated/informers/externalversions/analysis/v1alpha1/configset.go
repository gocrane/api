// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	analysisv1alpha1 "github.com/gocrane/api/analysis/v1alpha1"
	versioned "github.com/gocrane/api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gocrane/api/pkg/generated/listers/analysis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConfigSetInformer provides access to a shared informer and lister for
// ConfigSets.
type ConfigSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ConfigSetLister
}

type configSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConfigSetInformer constructs a new informer for ConfigSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConfigSetInformer constructs a new informer for ConfigSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalysisV1alpha1().ConfigSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AnalysisV1alpha1().ConfigSets(namespace).Watch(context.TODO(), options)
			},
		},
		&analysisv1alpha1.ConfigSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *configSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&analysisv1alpha1.ConfigSet{}, f.defaultInformer)
}

func (f *configSetInformer) Lister() v1alpha1.ConfigSetLister {
	return v1alpha1.NewConfigSetLister(f.Informer().GetIndexer())
}
