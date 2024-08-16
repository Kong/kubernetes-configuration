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

// FakeKongCustomEntities implements KongCustomEntityInterface
type FakeKongCustomEntities struct {
	Fake *FakeConfigurationV1alpha1
	ns   string
}

var kongcustomentitiesResource = v1alpha1.SchemeGroupVersion.WithResource("kongcustomentities")

var kongcustomentitiesKind = v1alpha1.SchemeGroupVersion.WithKind("KongCustomEntity")

// Get takes name of the kongCustomEntity, and returns the corresponding kongCustomEntity object, and an error if there is any.
func (c *FakeKongCustomEntities) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KongCustomEntity, err error) {
	emptyResult := &v1alpha1.KongCustomEntity{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(kongcustomentitiesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongCustomEntity), err
}

// List takes label and field selectors, and returns the list of KongCustomEntities that match those selectors.
func (c *FakeKongCustomEntities) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KongCustomEntityList, err error) {
	emptyResult := &v1alpha1.KongCustomEntityList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(kongcustomentitiesResource, kongcustomentitiesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KongCustomEntityList{ListMeta: obj.(*v1alpha1.KongCustomEntityList).ListMeta}
	for _, item := range obj.(*v1alpha1.KongCustomEntityList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongCustomEntities.
func (c *FakeKongCustomEntities) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(kongcustomentitiesResource, c.ns, opts))

}

// Create takes the representation of a kongCustomEntity and creates it.  Returns the server's representation of the kongCustomEntity, and an error, if there is any.
func (c *FakeKongCustomEntities) Create(ctx context.Context, kongCustomEntity *v1alpha1.KongCustomEntity, opts v1.CreateOptions) (result *v1alpha1.KongCustomEntity, err error) {
	emptyResult := &v1alpha1.KongCustomEntity{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(kongcustomentitiesResource, c.ns, kongCustomEntity, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongCustomEntity), err
}

// Update takes the representation of a kongCustomEntity and updates it. Returns the server's representation of the kongCustomEntity, and an error, if there is any.
func (c *FakeKongCustomEntities) Update(ctx context.Context, kongCustomEntity *v1alpha1.KongCustomEntity, opts v1.UpdateOptions) (result *v1alpha1.KongCustomEntity, err error) {
	emptyResult := &v1alpha1.KongCustomEntity{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(kongcustomentitiesResource, c.ns, kongCustomEntity, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongCustomEntity), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKongCustomEntities) UpdateStatus(ctx context.Context, kongCustomEntity *v1alpha1.KongCustomEntity, opts v1.UpdateOptions) (result *v1alpha1.KongCustomEntity, err error) {
	emptyResult := &v1alpha1.KongCustomEntity{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(kongcustomentitiesResource, "status", c.ns, kongCustomEntity, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongCustomEntity), err
}

// Delete takes name of the kongCustomEntity and deletes it. Returns an error if one occurs.
func (c *FakeKongCustomEntities) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kongcustomentitiesResource, c.ns, name, opts), &v1alpha1.KongCustomEntity{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongCustomEntities) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(kongcustomentitiesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KongCustomEntityList{})
	return err
}

// Patch applies the patch and returns the patched kongCustomEntity.
func (c *FakeKongCustomEntities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongCustomEntity, err error) {
	emptyResult := &v1alpha1.KongCustomEntity{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(kongcustomentitiesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongCustomEntity), err
}
