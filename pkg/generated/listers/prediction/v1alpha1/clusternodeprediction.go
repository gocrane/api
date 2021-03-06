// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gocrane/api/prediction/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterNodePredictionLister helps list ClusterNodePredictions.
// All objects returned here must be treated as read-only.
type ClusterNodePredictionLister interface {
	// List lists all ClusterNodePredictions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterNodePrediction, err error)
	// ClusterNodePredictions returns an object that can list and get ClusterNodePredictions.
	ClusterNodePredictions(namespace string) ClusterNodePredictionNamespaceLister
	ClusterNodePredictionListerExpansion
}

// clusterNodePredictionLister implements the ClusterNodePredictionLister interface.
type clusterNodePredictionLister struct {
	indexer cache.Indexer
}

// NewClusterNodePredictionLister returns a new ClusterNodePredictionLister.
func NewClusterNodePredictionLister(indexer cache.Indexer) ClusterNodePredictionLister {
	return &clusterNodePredictionLister{indexer: indexer}
}

// List lists all ClusterNodePredictions in the indexer.
func (s *clusterNodePredictionLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterNodePrediction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterNodePrediction))
	})
	return ret, err
}

// ClusterNodePredictions returns an object that can list and get ClusterNodePredictions.
func (s *clusterNodePredictionLister) ClusterNodePredictions(namespace string) ClusterNodePredictionNamespaceLister {
	return clusterNodePredictionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterNodePredictionNamespaceLister helps list and get ClusterNodePredictions.
// All objects returned here must be treated as read-only.
type ClusterNodePredictionNamespaceLister interface {
	// List lists all ClusterNodePredictions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterNodePrediction, err error)
	// Get retrieves the ClusterNodePrediction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterNodePrediction, error)
	ClusterNodePredictionNamespaceListerExpansion
}

// clusterNodePredictionNamespaceLister implements the ClusterNodePredictionNamespaceLister
// interface.
type clusterNodePredictionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterNodePredictions in the indexer for a given namespace.
func (s clusterNodePredictionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterNodePrediction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterNodePrediction))
	})
	return ret, err
}

// Get retrieves the ClusterNodePrediction from the indexer for a given namespace and name.
func (s clusterNodePredictionNamespaceLister) Get(name string) (*v1alpha1.ClusterNodePrediction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusternodeprediction"), name)
	}
	return obj.(*v1alpha1.ClusterNodePrediction), nil
}
