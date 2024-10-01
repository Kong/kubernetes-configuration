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
	"context"

	v1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KongCredentialJWTsGetter has a method to return a KongCredentialJWTInterface.
// A group's client should implement this interface.
type KongCredentialJWTsGetter interface {
	KongCredentialJWTs(namespace string) KongCredentialJWTInterface
}

// KongCredentialJWTInterface has methods to work with KongCredentialJWT resources.
type KongCredentialJWTInterface interface {
	Create(ctx context.Context, kongCredentialJWT *v1alpha1.KongCredentialJWT, opts v1.CreateOptions) (*v1alpha1.KongCredentialJWT, error)
	Update(ctx context.Context, kongCredentialJWT *v1alpha1.KongCredentialJWT, opts v1.UpdateOptions) (*v1alpha1.KongCredentialJWT, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kongCredentialJWT *v1alpha1.KongCredentialJWT, opts v1.UpdateOptions) (*v1alpha1.KongCredentialJWT, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KongCredentialJWT, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KongCredentialJWTList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KongCredentialJWT, err error)
	KongCredentialJWTExpansion
}

// kongCredentialJWTs implements KongCredentialJWTInterface
type kongCredentialJWTs struct {
	*gentype.ClientWithList[*v1alpha1.KongCredentialJWT, *v1alpha1.KongCredentialJWTList]
}

// newKongCredentialJWTs returns a KongCredentialJWTs
func newKongCredentialJWTs(c *ConfigurationV1alpha1Client, namespace string) *kongCredentialJWTs {
	return &kongCredentialJWTs{
		gentype.NewClientWithList[*v1alpha1.KongCredentialJWT, *v1alpha1.KongCredentialJWTList](
			"kongcredentialjwts",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.KongCredentialJWT { return &v1alpha1.KongCredentialJWT{} },
			func() *v1alpha1.KongCredentialJWTList { return &v1alpha1.KongCredentialJWTList{} }),
	}
}
