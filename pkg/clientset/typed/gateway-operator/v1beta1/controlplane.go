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

package v1beta1

import (
	context "context"

	gatewayoperatorv1beta1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1beta1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ControlPlanesGetter has a method to return a ControlPlaneInterface.
// A group's client should implement this interface.
type ControlPlanesGetter interface {
	ControlPlanes(namespace string) ControlPlaneInterface
}

// ControlPlaneInterface has methods to work with ControlPlane resources.
type ControlPlaneInterface interface {
	Create(ctx context.Context, controlPlane *gatewayoperatorv1beta1.ControlPlane, opts v1.CreateOptions) (*gatewayoperatorv1beta1.ControlPlane, error)
	Update(ctx context.Context, controlPlane *gatewayoperatorv1beta1.ControlPlane, opts v1.UpdateOptions) (*gatewayoperatorv1beta1.ControlPlane, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, controlPlane *gatewayoperatorv1beta1.ControlPlane, opts v1.UpdateOptions) (*gatewayoperatorv1beta1.ControlPlane, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*gatewayoperatorv1beta1.ControlPlane, error)
	List(ctx context.Context, opts v1.ListOptions) (*gatewayoperatorv1beta1.ControlPlaneList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gatewayoperatorv1beta1.ControlPlane, err error)
	ControlPlaneExpansion
}

// controlPlanes implements ControlPlaneInterface
type controlPlanes struct {
	*gentype.ClientWithList[*gatewayoperatorv1beta1.ControlPlane, *gatewayoperatorv1beta1.ControlPlaneList]
}

// newControlPlanes returns a ControlPlanes
func newControlPlanes(c *GatewayoperatorV1beta1Client, namespace string) *controlPlanes {
	return &controlPlanes{
		gentype.NewClientWithList[*gatewayoperatorv1beta1.ControlPlane, *gatewayoperatorv1beta1.ControlPlaneList](
			"controlplanes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *gatewayoperatorv1beta1.ControlPlane { return &gatewayoperatorv1beta1.ControlPlane{} },
			func() *gatewayoperatorv1beta1.ControlPlaneList { return &gatewayoperatorv1beta1.ControlPlaneList{} },
		),
	}
}
