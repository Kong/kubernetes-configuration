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

// KongTargetsGetter has a method to return a KongTargetInterface.
// A group's client should implement this interface.
type KongTargetsGetter interface {
	KongTargets(namespace string) KongTargetInterface
}

// KongTargetInterface has methods to work with KongTarget resources.
type KongTargetInterface interface {
	Create(ctx context.Context, kongTarget *configurationv1alpha1.KongTarget, opts v1.CreateOptions) (*configurationv1alpha1.KongTarget, error)
	Update(ctx context.Context, kongTarget *configurationv1alpha1.KongTarget, opts v1.UpdateOptions) (*configurationv1alpha1.KongTarget, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongTarget *configurationv1alpha1.KongTarget, opts v1.UpdateOptions) (*configurationv1alpha1.KongTarget, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*configurationv1alpha1.KongTarget, error)
	List(ctx context.Context, opts v1.ListOptions) (*configurationv1alpha1.KongTargetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configurationv1alpha1.KongTarget, err error)
	KongTargetExpansion
}

// kongTargets implements KongTargetInterface
type kongTargets struct {
	*gentype.ClientWithList[*configurationv1alpha1.KongTarget, *configurationv1alpha1.KongTargetList]
}

// newKongTargets returns a KongTargets
func newKongTargets(c *ConfigurationV1alpha1Client, namespace string) *kongTargets {
	return &kongTargets{
		gentype.NewClientWithList[*configurationv1alpha1.KongTarget, *configurationv1alpha1.KongTargetList](
			"kongtargets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *configurationv1alpha1.KongTarget { return &configurationv1alpha1.KongTarget{} },
			func() *configurationv1alpha1.KongTargetList { return &configurationv1alpha1.KongTargetList{} },
		),
	}
}
