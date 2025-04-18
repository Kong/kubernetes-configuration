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

// DataPlanesGetter has a method to return a DataPlaneInterface.
// A group's client should implement this interface.
type DataPlanesGetter interface {
	DataPlanes(namespace string) DataPlaneInterface
}

// DataPlaneInterface has methods to work with DataPlane resources.
type DataPlaneInterface interface {
	Create(ctx context.Context, dataPlane *gatewayoperatorv1beta1.DataPlane, opts v1.CreateOptions) (*gatewayoperatorv1beta1.DataPlane, error)
	Update(ctx context.Context, dataPlane *gatewayoperatorv1beta1.DataPlane, opts v1.UpdateOptions) (*gatewayoperatorv1beta1.DataPlane, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, dataPlane *gatewayoperatorv1beta1.DataPlane, opts v1.UpdateOptions) (*gatewayoperatorv1beta1.DataPlane, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*gatewayoperatorv1beta1.DataPlane, error)
	List(ctx context.Context, opts v1.ListOptions) (*gatewayoperatorv1beta1.DataPlaneList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *gatewayoperatorv1beta1.DataPlane, err error)
	DataPlaneExpansion
}

// dataPlanes implements DataPlaneInterface
type dataPlanes struct {
	*gentype.ClientWithList[*gatewayoperatorv1beta1.DataPlane, *gatewayoperatorv1beta1.DataPlaneList]
}

// newDataPlanes returns a DataPlanes
func newDataPlanes(c *GatewayOperatorV1beta1Client, namespace string) *dataPlanes {
	return &dataPlanes{
		gentype.NewClientWithList[*gatewayoperatorv1beta1.DataPlane, *gatewayoperatorv1beta1.DataPlaneList](
			"dataplanes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *gatewayoperatorv1beta1.DataPlane { return &gatewayoperatorv1beta1.DataPlane{} },
			func() *gatewayoperatorv1beta1.DataPlaneList { return &gatewayoperatorv1beta1.DataPlaneList{} },
		),
	}
}
