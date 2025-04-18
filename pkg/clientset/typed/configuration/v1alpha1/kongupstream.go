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

package v1alpha1

import (
	context "context"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongUpstreamsGetter has a method to return a KongUpstreamInterface.
// A group's client should implement this interface.
type KongUpstreamsGetter interface {
	KongUpstreams(namespace string) KongUpstreamInterface
}

// KongUpstreamInterface has methods to work with KongUpstream resources.
type KongUpstreamInterface interface {
	Create(ctx context.Context, kongUpstream *configurationv1alpha1.KongUpstream, opts v1.CreateOptions) (*configurationv1alpha1.KongUpstream, error)
	Update(ctx context.Context, kongUpstream *configurationv1alpha1.KongUpstream, opts v1.UpdateOptions) (*configurationv1alpha1.KongUpstream, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongUpstream *configurationv1alpha1.KongUpstream, opts v1.UpdateOptions) (*configurationv1alpha1.KongUpstream, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*configurationv1alpha1.KongUpstream, error)
	List(ctx context.Context, opts v1.ListOptions) (*configurationv1alpha1.KongUpstreamList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configurationv1alpha1.KongUpstream, err error)
	KongUpstreamExpansion
}

// kongUpstreams implements KongUpstreamInterface
type kongUpstreams struct {
	*gentype.ClientWithList[*configurationv1alpha1.KongUpstream, *configurationv1alpha1.KongUpstreamList]
}

// newKongUpstreams returns a KongUpstreams
func newKongUpstreams(c *ConfigurationV1alpha1Client, namespace string) *kongUpstreams {
	return &kongUpstreams{
		gentype.NewClientWithList[*configurationv1alpha1.KongUpstream, *configurationv1alpha1.KongUpstreamList](
			"kongupstreams",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *configurationv1alpha1.KongUpstream { return &configurationv1alpha1.KongUpstream{} },
			func() *configurationv1alpha1.KongUpstreamList { return &configurationv1alpha1.KongUpstreamList{} },
		),
	}
}
