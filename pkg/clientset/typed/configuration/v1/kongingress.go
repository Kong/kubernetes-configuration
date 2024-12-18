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

package v1

import (
	context "context"

	configurationv1 "github.com/kong/kubernetes-configuration/api/configuration/v1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongIngressesGetter has a method to return a KongIngressInterface.
// A group's client should implement this interface.
type KongIngressesGetter interface {
	KongIngresses(namespace string) KongIngressInterface
}

// KongIngressInterface has methods to work with KongIngress resources.
type KongIngressInterface interface {
	Create(ctx context.Context, kongIngress *configurationv1.KongIngress, opts metav1.CreateOptions) (*configurationv1.KongIngress, error)
	Update(ctx context.Context, kongIngress *configurationv1.KongIngress, opts metav1.UpdateOptions) (*configurationv1.KongIngress, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*configurationv1.KongIngress, error)
	List(ctx context.Context, opts metav1.ListOptions) (*configurationv1.KongIngressList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *configurationv1.KongIngress, err error)
	KongIngressExpansion
}

// kongIngresses implements KongIngressInterface
type kongIngresses struct {
	*gentype.ClientWithList[*configurationv1.KongIngress, *configurationv1.KongIngressList]
}

// newKongIngresses returns a KongIngresses
func newKongIngresses(c *ConfigurationV1Client, namespace string) *kongIngresses {
	return &kongIngresses{
		gentype.NewClientWithList[*configurationv1.KongIngress, *configurationv1.KongIngressList](
			"kongingresses",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *configurationv1.KongIngress { return &configurationv1.KongIngress{} },
			func() *configurationv1.KongIngressList { return &configurationv1.KongIngressList{} },
		),
	}
}
