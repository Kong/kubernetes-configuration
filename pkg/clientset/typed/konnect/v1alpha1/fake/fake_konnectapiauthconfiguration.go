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
	v1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/pkg/clientset/typed/konnect/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeKonnectAPIAuthConfigurations implements KonnectAPIAuthConfigurationInterface
type fakeKonnectAPIAuthConfigurations struct {
	*gentype.FakeClientWithList[*v1alpha1.KonnectAPIAuthConfiguration, *v1alpha1.KonnectAPIAuthConfigurationList]
	Fake *FakeKonnectV1alpha1
}

func newFakeKonnectAPIAuthConfigurations(fake *FakeKonnectV1alpha1, namespace string) konnectv1alpha1.KonnectAPIAuthConfigurationInterface {
	return &fakeKonnectAPIAuthConfigurations{
		gentype.NewFakeClientWithList[*v1alpha1.KonnectAPIAuthConfiguration, *v1alpha1.KonnectAPIAuthConfigurationList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("konnectapiauthconfigurations"),
			v1alpha1.SchemeGroupVersion.WithKind("KonnectAPIAuthConfiguration"),
			func() *v1alpha1.KonnectAPIAuthConfiguration { return &v1alpha1.KonnectAPIAuthConfiguration{} },
			func() *v1alpha1.KonnectAPIAuthConfigurationList { return &v1alpha1.KonnectAPIAuthConfigurationList{} },
			func(dst, src *v1alpha1.KonnectAPIAuthConfigurationList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.KonnectAPIAuthConfigurationList) []*v1alpha1.KonnectAPIAuthConfiguration {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.KonnectAPIAuthConfigurationList, items []*v1alpha1.KonnectAPIAuthConfiguration) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
