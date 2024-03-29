// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	ensurancev1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	versioned "github.com/gocrane/api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gocrane/api/pkg/generated/listers/ensurance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodQOSInformer provides access to a shared informer and lister for
// PodQOSs.
type PodQOSInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodQOSLister
}

type podQOSInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewPodQOSInformer constructs a new informer for PodQOS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodQOSInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodQOSInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredPodQOSInformer constructs a new informer for PodQOS type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodQOSInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnsuranceV1alpha1().PodQOSs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EnsuranceV1alpha1().PodQOSs().Watch(context.TODO(), options)
			},
		},
		&ensurancev1alpha1.PodQOS{},
		resyncPeriod,
		indexers,
	)
}

func (f *podQOSInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodQOSInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podQOSInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ensurancev1alpha1.PodQOS{}, f.defaultInformer)
}

func (f *podQOSInformer) Lister() v1alpha1.PodQOSLister {
	return v1alpha1.NewPodQOSLister(f.Informer().GetIndexer())
}
