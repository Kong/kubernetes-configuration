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

	gatewayoperatorv1alpha1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongPluginInstallationsGetter has a method to return a KongPluginInstallationInterface.
// A group's client should implement this interface.
type KongPluginInstallationsGetter interface {
	KongPluginInstallations(namespace string) KongPluginInstallationInterface
}

// KongPluginInstallationInterface has methods to work with KongPluginInstallation resources.
type KongPluginInstallationInterface interface {
	Create(ctx context.Context, kongPluginInstallation *gatewayoperatorv1alpha1.KongPluginInstallation, opts v1.CreateOptions) (*gatewayoperatorv1alpha1.KongPluginInstallation, error)
	Update(ctx context.Context, kongPluginInstallation *gatewayoperatorv1alpha1.KongPluginInstallation, opts v1.UpdateOptions) (*gatewayoperatorv1alpha1.KongPluginInstallation, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongPluginInstallation *gatewayoperatorv1alpha1.KongPluginInstallation, opts v1.UpdateOptions) (*gatewayoperatorv1alpha1.KongPluginInstallation, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*gatewayoperatorv1alpha1.KongPluginInstallation, error)
	List(ctx context.Context, opts v1.ListOptions) (*gatewayoperatorv1alpha1.KongPluginInstallationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gatewayoperatorv1alpha1.KongPluginInstallation, err error)
	KongPluginInstallationExpansion
}

// kongPluginInstallations implements KongPluginInstallationInterface
type kongPluginInstallations struct {
	*gentype.ClientWithList[*gatewayoperatorv1alpha1.KongPluginInstallation, *gatewayoperatorv1alpha1.KongPluginInstallationList]
}

// newKongPluginInstallations returns a KongPluginInstallations
func newKongPluginInstallations(c *GatewayOperatorV1alpha1Client, namespace string) *kongPluginInstallations {
	return &kongPluginInstallations{
		gentype.NewClientWithList[*gatewayoperatorv1alpha1.KongPluginInstallation, *gatewayoperatorv1alpha1.KongPluginInstallationList](
			"kongplugininstallations",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *gatewayoperatorv1alpha1.KongPluginInstallation {
				return &gatewayoperatorv1alpha1.KongPluginInstallation{}
			},
			func() *gatewayoperatorv1alpha1.KongPluginInstallationList {
				return &gatewayoperatorv1alpha1.KongPluginInstallationList{}
			},
		),
	}
}
