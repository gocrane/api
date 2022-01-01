// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/analysis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigSets implements ConfigSetInterface
type FakeConfigSets struct {
	Fake *FakeAnalysisV1alpha1
	ns   string
}

var configsetsResource = schema.GroupVersionResource{Group: "analysis.crane.io", Version: "v1alpha1", Resource: "configsets"}

var configsetsKind = schema.GroupVersionKind{Group: "analysis.crane.io", Version: "v1alpha1", Kind: "ConfigSet"}

// Get takes name of the configSet, and returns the corresponding configSet object, and an error if there is any.
func (c *FakeConfigSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configsetsResource, c.ns, name), &v1alpha1.ConfigSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigSet), err
}

// List takes label and field selectors, and returns the list of ConfigSets that match those selectors.
func (c *FakeConfigSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configsetsResource, configsetsKind, c.ns, opts), &v1alpha1.ConfigSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigSetList{ListMeta: obj.(*v1alpha1.ConfigSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configSets.
func (c *FakeConfigSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configsetsResource, c.ns, opts))

}

// Create takes the representation of a configSet and creates it.  Returns the server's representation of the configSet, and an error, if there is any.
func (c *FakeConfigSets) Create(ctx context.Context, configSet *v1alpha1.ConfigSet, opts v1.CreateOptions) (result *v1alpha1.ConfigSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configsetsResource, c.ns, configSet), &v1alpha1.ConfigSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigSet), err
}

// Update takes the representation of a configSet and updates it. Returns the server's representation of the configSet, and an error, if there is any.
func (c *FakeConfigSets) Update(ctx context.Context, configSet *v1alpha1.ConfigSet, opts v1.UpdateOptions) (result *v1alpha1.ConfigSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configsetsResource, c.ns, configSet), &v1alpha1.ConfigSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigSet), err
}

// Delete takes name of the configSet and deletes it. Returns an error if one occurs.
func (c *FakeConfigSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configsetsResource, c.ns, name), &v1alpha1.ConfigSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigSetList{})
	return err
}

// Patch applies the patch and returns the patched configSet.
func (c *FakeConfigSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigSet), err
}
