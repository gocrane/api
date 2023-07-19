// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTimeSeriesPredictions implements TimeSeriesPredictionInterface
type FakeTimeSeriesPredictions struct {
	Fake *FakePredictionV1alpha1
	ns   string
}

var timeseriespredictionsResource = schema.GroupVersionResource{Group: "prediction.crane.io", Version: "v1alpha1", Resource: "timeseriespredictions"}

var timeseriespredictionsKind = schema.GroupVersionKind{Group: "prediction.crane.io", Version: "v1alpha1", Kind: "TimeSeriesPrediction"}

// Get takes name of the timeSeriesPrediction, and returns the corresponding timeSeriesPrediction object, and an error if there is any.
func (c *FakeTimeSeriesPredictions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(timeseriespredictionsResource, c.ns, name), &v1alpha1.TimeSeriesPrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), err
}

// List takes label and field selectors, and returns the list of TimeSeriesPredictions that match those selectors.
func (c *FakeTimeSeriesPredictions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TimeSeriesPredictionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(timeseriespredictionsResource, timeseriespredictionsKind, c.ns, opts), &v1alpha1.TimeSeriesPredictionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TimeSeriesPredictionList{ListMeta: obj.(*v1alpha1.TimeSeriesPredictionList).ListMeta}
	for _, item := range obj.(*v1alpha1.TimeSeriesPredictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested timeSeriesPredictions.
func (c *FakeTimeSeriesPredictions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(timeseriespredictionsResource, c.ns, opts))

}

// Create takes the representation of a timeSeriesPrediction and creates it.  Returns the server's representation of the timeSeriesPrediction, and an error, if there is any.
func (c *FakeTimeSeriesPredictions) Create(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.CreateOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(timeseriespredictionsResource, c.ns, timeSeriesPrediction), &v1alpha1.TimeSeriesPrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), err
}

// Update takes the representation of a timeSeriesPrediction and updates it. Returns the server's representation of the timeSeriesPrediction, and an error, if there is any.
func (c *FakeTimeSeriesPredictions) Update(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (result *v1alpha1.TimeSeriesPrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(timeseriespredictionsResource, c.ns, timeSeriesPrediction), &v1alpha1.TimeSeriesPrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTimeSeriesPredictions) UpdateStatus(ctx context.Context, timeSeriesPrediction *v1alpha1.TimeSeriesPrediction, opts v1.UpdateOptions) (*v1alpha1.TimeSeriesPrediction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(timeseriespredictionsResource, "status", c.ns, timeSeriesPrediction), &v1alpha1.TimeSeriesPrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), err
}

// Delete takes name of the timeSeriesPrediction and deletes it. Returns an error if one occurs.
func (c *FakeTimeSeriesPredictions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(timeseriespredictionsResource, c.ns, name, opts), &v1alpha1.TimeSeriesPrediction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTimeSeriesPredictions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(timeseriespredictionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TimeSeriesPredictionList{})
	return err
}

// Patch applies the patch and returns the patched timeSeriesPrediction.
func (c *FakeTimeSeriesPredictions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TimeSeriesPrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(timeseriespredictionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TimeSeriesPrediction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TimeSeriesPrediction), err
}