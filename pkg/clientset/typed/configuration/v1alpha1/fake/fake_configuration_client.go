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
	v1alpha1 "github.com/kong/kubernetes-configuration/pkg/clientset/typed/configuration/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeConfigurationV1alpha1 struct {
	*testing.Fake
}

func (c *FakeConfigurationV1alpha1) IngressClassParameterses(namespace string) v1alpha1.IngressClassParametersInterface {
	return &FakeIngressClassParameterses{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCACertificates(namespace string) v1alpha1.KongCACertificateInterface {
	return &FakeKongCACertificates{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCertificates(namespace string) v1alpha1.KongCertificateInterface {
	return &FakeKongCertificates{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCredentialACLs(namespace string) v1alpha1.KongCredentialACLInterface {
	return &FakeKongCredentialACLs{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCredentialAPIKeys(namespace string) v1alpha1.KongCredentialAPIKeyInterface {
	return &FakeKongCredentialAPIKeys{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCredentialBasicAuths(namespace string) v1alpha1.KongCredentialBasicAuthInterface {
	return &FakeKongCredentialBasicAuths{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCredentialHMACs(namespace string) v1alpha1.KongCredentialHMACInterface {
	return &FakeKongCredentialHMACs{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCredentialJWTs(namespace string) v1alpha1.KongCredentialJWTInterface {
	return &FakeKongCredentialJWTs{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongCustomEntities(namespace string) v1alpha1.KongCustomEntityInterface {
	return &FakeKongCustomEntities{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongDataPlaneClientCertificates(namespace string) v1alpha1.KongDataPlaneClientCertificateInterface {
	return &FakeKongDataPlaneClientCertificates{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongKeys(namespace string) v1alpha1.KongKeyInterface {
	return &FakeKongKeys{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongKeySets(namespace string) v1alpha1.KongKeySetInterface {
	return &FakeKongKeySets{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongLicenses() v1alpha1.KongLicenseInterface {
	return &FakeKongLicenses{c}
}

func (c *FakeConfigurationV1alpha1) KongPluginBindings(namespace string) v1alpha1.KongPluginBindingInterface {
	return &FakeKongPluginBindings{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongRoutes(namespace string) v1alpha1.KongRouteInterface {
	return &FakeKongRoutes{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongSNIs(namespace string) v1alpha1.KongSNIInterface {
	return &FakeKongSNIs{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongServices(namespace string) v1alpha1.KongServiceInterface {
	return &FakeKongServices{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongTargets(namespace string) v1alpha1.KongTargetInterface {
	return &FakeKongTargets{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongUpstreams(namespace string) v1alpha1.KongUpstreamInterface {
	return &FakeKongUpstreams{c, namespace}
}

func (c *FakeConfigurationV1alpha1) KongVaults() v1alpha1.KongVaultInterface {
	return &FakeKongVaults{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeConfigurationV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
