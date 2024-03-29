// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/ensurance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeQOSs implements NodeQOSInterface
type FakeNodeQOSs struct {
	Fake *FakeEnsuranceV1alpha1
}

var nodeqossResource = schema.GroupVersionResource{Group: "ensurance.crane.io", Version: "v1alpha1", Resource: "nodeqoss"}

var nodeqossKind = schema.GroupVersionKind{Group: "ensurance.crane.io", Version: "v1alpha1", Kind: "NodeQOS"}

// Get takes name of the nodeQOS, and returns the corresponding nodeQOS object, and an error if there is any.
func (c *FakeNodeQOSs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodeqossResource, name), &v1alpha1.NodeQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeQOS), err
}

// List takes label and field selectors, and returns the list of NodeQOSs that match those selectors.
func (c *FakeNodeQOSs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeQOSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodeqossResource, nodeqossKind, opts), &v1alpha1.NodeQOSList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeQOSList{ListMeta: obj.(*v1alpha1.NodeQOSList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeQOSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeQOSs.
func (c *FakeNodeQOSs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodeqossResource, opts))
}

// Create takes the representation of a nodeQOS and creates it.  Returns the server's representation of the nodeQOS, and an error, if there is any.
func (c *FakeNodeQOSs) Create(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.CreateOptions) (result *v1alpha1.NodeQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodeqossResource, nodeQOS), &v1alpha1.NodeQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeQOS), err
}

// Update takes the representation of a nodeQOS and updates it. Returns the server's representation of the nodeQOS, and an error, if there is any.
func (c *FakeNodeQOSs) Update(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (result *v1alpha1.NodeQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodeqossResource, nodeQOS), &v1alpha1.NodeQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeQOS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeQOSs) UpdateStatus(ctx context.Context, nodeQOS *v1alpha1.NodeQOS, opts v1.UpdateOptions) (*v1alpha1.NodeQOS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodeqossResource, "status", nodeQOS), &v1alpha1.NodeQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeQOS), err
}

// Delete takes name of the nodeQOS and deletes it. Returns an error if one occurs.
func (c *FakeNodeQOSs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodeqossResource, name), &v1alpha1.NodeQOS{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeQOSs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodeqossResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeQOSList{})
	return err
}

// Patch applies the patch and returns the patched nodeQOS.
func (c *FakeNodeQOSs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeQOS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodeqossResource, name, pt, data, subresources...), &v1alpha1.NodeQOS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeQOS), err
}
