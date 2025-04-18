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

	incubatorv1alpha1 "github.com/kong/kubernetes-configuration/api/incubator/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongServiceFacadesGetter has a method to return a KongServiceFacadeInterface.
// A group's client should implement this interface.
type KongServiceFacadesGetter interface {
	KongServiceFacades(namespace string) KongServiceFacadeInterface
}

// KongServiceFacadeInterface has methods to work with KongServiceFacade resources.
type KongServiceFacadeInterface interface {
	Create(ctx context.Context, kongServiceFacade *incubatorv1alpha1.KongServiceFacade, opts v1.CreateOptions) (*incubatorv1alpha1.KongServiceFacade, error)
	Update(ctx context.Context, kongServiceFacade *incubatorv1alpha1.KongServiceFacade, opts v1.UpdateOptions) (*incubatorv1alpha1.KongServiceFacade, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongServiceFacade *incubatorv1alpha1.KongServiceFacade, opts v1.UpdateOptions) (*incubatorv1alpha1.KongServiceFacade, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*incubatorv1alpha1.KongServiceFacade, error)
	List(ctx context.Context, opts v1.ListOptions) (*incubatorv1alpha1.KongServiceFacadeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *incubatorv1alpha1.KongServiceFacade, err error)
	KongServiceFacadeExpansion
}

// kongServiceFacades implements KongServiceFacadeInterface
type kongServiceFacades struct {
	*gentype.ClientWithList[*incubatorv1alpha1.KongServiceFacade, *incubatorv1alpha1.KongServiceFacadeList]
}

// newKongServiceFacades returns a KongServiceFacades
func newKongServiceFacades(c *IncubatorV1alpha1Client, namespace string) *kongServiceFacades {
	return &kongServiceFacades{
		gentype.NewClientWithList[*incubatorv1alpha1.KongServiceFacade, *incubatorv1alpha1.KongServiceFacadeList](
			"kongservicefacades",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *incubatorv1alpha1.KongServiceFacade { return &incubatorv1alpha1.KongServiceFacade{} },
			func() *incubatorv1alpha1.KongServiceFacadeList { return &incubatorv1alpha1.KongServiceFacadeList{} },
		),
	}
}
