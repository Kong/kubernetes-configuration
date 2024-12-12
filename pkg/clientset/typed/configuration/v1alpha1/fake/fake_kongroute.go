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
	v1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/pkg/clientset/typed/configuration/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeKongRoutes implements KongRouteInterface
type fakeKongRoutes struct {
	*gentype.FakeClientWithList[*v1alpha1.KongRoute, *v1alpha1.KongRouteList]
	Fake *FakeConfigurationV1alpha1
}

func newFakeKongRoutes(fake *FakeConfigurationV1alpha1, namespace string) configurationv1alpha1.KongRouteInterface {
	return &fakeKongRoutes{
		gentype.NewFakeClientWithList[*v1alpha1.KongRoute, *v1alpha1.KongRouteList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("kongroutes"),
			v1alpha1.SchemeGroupVersion.WithKind("KongRoute"),
			func() *v1alpha1.KongRoute { return &v1alpha1.KongRoute{} },
			func() *v1alpha1.KongRouteList { return &v1alpha1.KongRouteList{} },
			func(dst, src *v1alpha1.KongRouteList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.KongRouteList) []*v1alpha1.KongRoute { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.KongRouteList, items []*v1alpha1.KongRoute) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
