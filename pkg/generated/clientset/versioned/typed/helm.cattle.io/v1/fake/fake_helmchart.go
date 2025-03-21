/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/k3s-io/helm-controller/pkg/apis/helm.cattle.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelmCharts implements HelmChartInterface
type FakeHelmCharts struct {
	Fake *FakeHelmV1
	ns   string
}

var helmchartsResource = v1.SchemeGroupVersion.WithResource("helmcharts")

var helmchartsKind = v1.SchemeGroupVersion.WithKind("HelmChart")

// Get takes name of the helmChart, and returns the corresponding helmChart object, and an error if there is any.
func (c *FakeHelmCharts) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HelmChart, err error) {
	emptyResult := &v1.HelmChart{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(helmchartsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.HelmChart), err
}

// List takes label and field selectors, and returns the list of HelmCharts that match those selectors.
func (c *FakeHelmCharts) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HelmChartList, err error) {
	emptyResult := &v1.HelmChartList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(helmchartsResource, helmchartsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.HelmChartList{ListMeta: obj.(*v1.HelmChartList).ListMeta}
	for _, item := range obj.(*v1.HelmChartList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmCharts.
func (c *FakeHelmCharts) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(helmchartsResource, c.ns, opts))

}

// Create takes the representation of a helmChart and creates it.  Returns the server's representation of the helmChart, and an error, if there is any.
func (c *FakeHelmCharts) Create(ctx context.Context, helmChart *v1.HelmChart, opts metav1.CreateOptions) (result *v1.HelmChart, err error) {
	emptyResult := &v1.HelmChart{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(helmchartsResource, c.ns, helmChart, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.HelmChart), err
}

// Update takes the representation of a helmChart and updates it. Returns the server's representation of the helmChart, and an error, if there is any.
func (c *FakeHelmCharts) Update(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (result *v1.HelmChart, err error) {
	emptyResult := &v1.HelmChart{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(helmchartsResource, c.ns, helmChart, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.HelmChart), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelmCharts) UpdateStatus(ctx context.Context, helmChart *v1.HelmChart, opts metav1.UpdateOptions) (result *v1.HelmChart, err error) {
	emptyResult := &v1.HelmChart{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(helmchartsResource, "status", c.ns, helmChart, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.HelmChart), err
}

// Delete takes name of the helmChart and deletes it. Returns an error if one occurs.
func (c *FakeHelmCharts) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(helmchartsResource, c.ns, name, opts), &v1.HelmChart{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmCharts) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(helmchartsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.HelmChartList{})
	return err
}

// Patch applies the patch and returns the patched helmChart.
func (c *FakeHelmCharts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelmChart, err error) {
	emptyResult := &v1.HelmChart{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(helmchartsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.HelmChart), err
}
