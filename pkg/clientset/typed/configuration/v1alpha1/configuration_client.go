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
	http "net/http"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	scheme "github.com/kong/kubernetes-configuration/pkg/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ConfigurationV1alpha1Interface interface {
	RESTClient() rest.Interface
	IngressClassParametersesGetter
	KongCACertificatesGetter
	KongCertificatesGetter
	KongCredentialACLsGetter
	KongCredentialAPIKeysGetter
	KongCredentialBasicAuthsGetter
	KongCredentialHMACsGetter
	KongCredentialJWTsGetter
	KongCustomEntitiesGetter
	KongDataPlaneClientCertificatesGetter
	KongKeysGetter
	KongKeySetsGetter
	KongLicensesGetter
	KongPluginBindingsGetter
	KongRoutesGetter
	KongSNIsGetter
	KongServicesGetter
	KongTargetsGetter
	KongUpstreamsGetter
	KongVaultsGetter
}

// ConfigurationV1alpha1Client is used to interact with features provided by the configuration.konghq.com group.
type ConfigurationV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ConfigurationV1alpha1Client) IngressClassParameterses(namespace string) IngressClassParametersInterface {
	return newIngressClassParameterses(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCACertificates(namespace string) KongCACertificateInterface {
	return newKongCACertificates(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCertificates(namespace string) KongCertificateInterface {
	return newKongCertificates(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCredentialACLs(namespace string) KongCredentialACLInterface {
	return newKongCredentialACLs(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCredentialAPIKeys(namespace string) KongCredentialAPIKeyInterface {
	return newKongCredentialAPIKeys(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCredentialBasicAuths(namespace string) KongCredentialBasicAuthInterface {
	return newKongCredentialBasicAuths(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCredentialHMACs(namespace string) KongCredentialHMACInterface {
	return newKongCredentialHMACs(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCredentialJWTs(namespace string) KongCredentialJWTInterface {
	return newKongCredentialJWTs(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongCustomEntities(namespace string) KongCustomEntityInterface {
	return newKongCustomEntities(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongDataPlaneClientCertificates(namespace string) KongDataPlaneClientCertificateInterface {
	return newKongDataPlaneClientCertificates(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongKeys(namespace string) KongKeyInterface {
	return newKongKeys(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongKeySets(namespace string) KongKeySetInterface {
	return newKongKeySets(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongLicenses() KongLicenseInterface {
	return newKongLicenses(c)
}

func (c *ConfigurationV1alpha1Client) KongPluginBindings(namespace string) KongPluginBindingInterface {
	return newKongPluginBindings(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongRoutes(namespace string) KongRouteInterface {
	return newKongRoutes(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongSNIs(namespace string) KongSNIInterface {
	return newKongSNIs(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongServices(namespace string) KongServiceInterface {
	return newKongServices(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongTargets(namespace string) KongTargetInterface {
	return newKongTargets(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongUpstreams(namespace string) KongUpstreamInterface {
	return newKongUpstreams(c, namespace)
}

func (c *ConfigurationV1alpha1Client) KongVaults() KongVaultInterface {
	return newKongVaults(c)
}

// NewForConfig creates a new ConfigurationV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ConfigurationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ConfigurationV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ConfigurationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ConfigurationV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ConfigurationV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ConfigurationV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ConfigurationV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ConfigurationV1alpha1Client {
	return &ConfigurationV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := configurationv1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ConfigurationV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
