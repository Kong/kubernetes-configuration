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

// FakeKongDataPlaneClientCertificates implements KongDataPlaneClientCertificateInterface
type FakeKongDataPlaneClientCertificates struct {
	Fake *FakeConfigurationV1alpha1
	ns   string
}

var kongdataplaneclientcertificatesResource = v1alpha1.SchemeGroupVersion.WithResource("kongdataplaneclientcertificates")

var kongdataplaneclientcertificatesKind = v1alpha1.SchemeGroupVersion.WithKind("KongDataPlaneClientCertificate")

// Get takes name of the kongDataPlaneClientCertificate, and returns the corresponding kongDataPlaneClientCertificate object, and an error if there is any.
func (c *FakeKongDataPlaneClientCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KongDataPlaneClientCertificate, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongDataPlaneClientCertificate), err
}

// List takes label and field selectors, and returns the list of KongDataPlaneClientCertificates that match those selectors.
func (c *FakeKongDataPlaneClientCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KongDataPlaneClientCertificateList, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificateList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(kongdataplaneclientcertificatesResource, kongdataplaneclientcertificatesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KongDataPlaneClientCertificateList{ListMeta: obj.(*v1alpha1.KongDataPlaneClientCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.KongDataPlaneClientCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongDataPlaneClientCertificates.
func (c *FakeKongDataPlaneClientCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, opts))

}

// Create takes the representation of a kongDataPlaneClientCertificate and creates it.  Returns the server's representation of the kongDataPlaneClientCertificate, and an error, if there is any.
func (c *FakeKongDataPlaneClientCertificates) Create(ctx context.Context, kongDataPlaneClientCertificate *v1alpha1.KongDataPlaneClientCertificate, opts v1.CreateOptions) (result *v1alpha1.KongDataPlaneClientCertificate, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, kongDataPlaneClientCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongDataPlaneClientCertificate), err
}

// Update takes the representation of a kongDataPlaneClientCertificate and updates it. Returns the server's representation of the kongDataPlaneClientCertificate, and an error, if there is any.
func (c *FakeKongDataPlaneClientCertificates) Update(ctx context.Context, kongDataPlaneClientCertificate *v1alpha1.KongDataPlaneClientCertificate, opts v1.UpdateOptions) (result *v1alpha1.KongDataPlaneClientCertificate, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, kongDataPlaneClientCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongDataPlaneClientCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKongDataPlaneClientCertificates) UpdateStatus(ctx context.Context, kongDataPlaneClientCertificate *v1alpha1.KongDataPlaneClientCertificate, opts v1.UpdateOptions) (result *v1alpha1.KongDataPlaneClientCertificate, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(kongdataplaneclientcertificatesResource, "status", c.ns, kongDataPlaneClientCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongDataPlaneClientCertificate), err
}

// Delete takes name of the kongDataPlaneClientCertificate and deletes it. Returns an error if one occurs.
func (c *FakeKongDataPlaneClientCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, name, opts), &v1alpha1.KongDataPlaneClientCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongDataPlaneClientCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KongDataPlaneClientCertificateList{})
	return err
}

// Patch applies the patch and returns the patched kongDataPlaneClientCertificate.
func (c *FakeKongDataPlaneClientCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongDataPlaneClientCertificate, err error) {
	emptyResult := &v1alpha1.KongDataPlaneClientCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(kongdataplaneclientcertificatesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.KongDataPlaneClientCertificate), err
}
