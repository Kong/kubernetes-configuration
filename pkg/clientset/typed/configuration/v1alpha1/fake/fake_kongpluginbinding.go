/*
Copyright 2021 Kong, Inc.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKongPluginBindings implements KongPluginBindingInterface
type FakeKongPluginBindings struct {
	Fake *FakeConfigurationV1alpha1
	ns   string
}

var kongpluginbindingsResource = v1alpha1.SchemeGroupVersion.WithResource("kongpluginbindings")

var kongpluginbindingsKind = v1alpha1.SchemeGroupVersion.WithKind("KongPluginBinding")

// Get takes name of the kongPluginBinding, and returns the corresponding kongPluginBinding object, and an error if there is any.
func (c *FakeKongPluginBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KongPluginBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kongpluginbindingsResource, c.ns, name), &v1alpha1.KongPluginBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginBinding), err
}

// List takes label and field selectors, and returns the list of KongPluginBindings that match those selectors.
func (c *FakeKongPluginBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KongPluginBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kongpluginbindingsResource, kongpluginbindingsKind, c.ns, opts), &v1alpha1.KongPluginBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KongPluginBindingList{ListMeta: obj.(*v1alpha1.KongPluginBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.KongPluginBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongPluginBindings.
func (c *FakeKongPluginBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kongpluginbindingsResource, c.ns, opts))

}

// Create takes the representation of a kongPluginBinding and creates it.  Returns the server's representation of the kongPluginBinding, and an error, if there is any.
func (c *FakeKongPluginBindings) Create(ctx context.Context, kongPluginBinding *v1alpha1.KongPluginBinding, opts v1.CreateOptions) (result *v1alpha1.KongPluginBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kongpluginbindingsResource, c.ns, kongPluginBinding), &v1alpha1.KongPluginBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginBinding), err
}

// Update takes the representation of a kongPluginBinding and updates it. Returns the server's representation of the kongPluginBinding, and an error, if there is any.
func (c *FakeKongPluginBindings) Update(ctx context.Context, kongPluginBinding *v1alpha1.KongPluginBinding, opts v1.UpdateOptions) (result *v1alpha1.KongPluginBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kongpluginbindingsResource, c.ns, kongPluginBinding), &v1alpha1.KongPluginBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKongPluginBindings) UpdateStatus(ctx context.Context, kongPluginBinding *v1alpha1.KongPluginBinding, opts v1.UpdateOptions) (*v1alpha1.KongPluginBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kongpluginbindingsResource, "status", c.ns, kongPluginBinding), &v1alpha1.KongPluginBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginBinding), err
}

// Delete takes name of the kongPluginBinding and deletes it. Returns an error if one occurs.
func (c *FakeKongPluginBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kongpluginbindingsResource, c.ns, name, opts), &v1alpha1.KongPluginBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongPluginBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kongpluginbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KongPluginBindingList{})
	return err
}

// Patch applies the patch and returns the patched kongPluginBinding.
func (c *FakeKongPluginBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongPluginBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kongpluginbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KongPluginBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KongPluginBinding), err
}