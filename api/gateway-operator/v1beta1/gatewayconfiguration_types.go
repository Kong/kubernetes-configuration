/*
Copyright 2022 Kong Inc.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
)

func init() {
	SchemeBuilder.Register(&GatewayConfiguration{}, &GatewayConfigurationList{})
}

// GatewayConfiguration is the Schema for the gatewayconfigurations API.
//
// +genclient
// +apireference:kgo:include
// +kong:channels=gateway-operator
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=kogc,categories=kong;all
// +kubebuilder:validation:XValidation:message="Extension not allowed for GatewayConfiguration",rule="has(self.spec.extensions) ? self.spec.extensions.all(e, (e.group == 'konnect.konghq.com' && e.kind == 'KonnectExtension') || (e.group == 'gateway-operator.konghq.com' && e.kind == 'DataPlaneMetricsExtension')) : true"
type GatewayConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewayConfigurationSpec   `json:"spec,omitempty"`
	Status GatewayConfigurationStatus `json:"status,omitempty"`
}

// GatewayConfigurationSpec defines the desired state of GatewayConfiguration
// +apireference:kgo:include
// +kubebuilder:validation:XValidation:message="KonnectExtension must be set at the Gateway level",rule="has(self.dataPlaneOptions) && has(self.dataPlaneOptions.extensions) ? self.dataPlaneOptions.extensions.all(e, (e.group != 'konnect.konghq.com' && e.group != 'gateway-operator.konghq.com') || e.kind != 'KonnectExtension') : true"
// +kubebuilder:validation:XValidation:message="KonnectExtension must be set at the Gateway level",rule="has(self.controlPlaneOptions) && has(self.controlPlaneOptions.extensions) ? self.controlPlaneOptions.extensions.all(e, (e.group != 'konnect.konghq.com' && e.group != 'gateway-operator.konghq.com') || e.kind != 'KonnectExtension') : true"
type GatewayConfigurationSpec struct {
	// DataPlaneOptions is the specification for configuration
	// overrides for DataPlane resources that will be created for the Gateway.
	//
	// +optional
	DataPlaneOptions *GatewayConfigDataPlaneOptions `json:"dataPlaneOptions,omitempty"`

	// ControlPlaneOptions is the specification for configuration
	// overrides for ControlPlane resources that will be created for the Gateway.
	//
	// +optional
	ControlPlaneOptions *ControlPlaneOptions `json:"controlPlaneOptions,omitempty"`

	// Extensions provide additional or replacement features for the Gateway
	// resource to influence or enhance functionality.
	// NOTE: currently, there's only 1 extension that can be attached
	// at the Gateway level (KonnectExtension), so the amount of extensions
	// is limited to 1.
	//
	// +optional
	// +kubebuilder:validation:MinItems=0
	// +kubebuilder:validation:MaxItems=1
	Extensions []commonv1alpha1.ExtensionRef `json:"extensions,omitempty"`
}

// GatewayConfigDataPlaneOptions indicates the specific information needed to
// configure and deploy a DataPlane object.
// +apireference:kgo:include
// +kubebuilder:validation:XValidation:message="Extension not allowed for DataPlane",rule="has(self.extensions) ? self.extensions.all(e, (e.group == 'konnect.konghq.com' || e.group == 'gateway-operator.konghq.com') && e.kind == 'KonnectExtension') : true"
type GatewayConfigDataPlaneOptions struct {
	// +optional
	Deployment DataPlaneDeploymentOptions `json:"deployment"`

	// +optional
	Network GatewayConfigDataPlaneNetworkOptions `json:"network"`

	// Extensions provide additional or replacement features for the DataPlane
	// resources to influence or enhance functionality.
	// NOTE: since we have one extension only (KonnectExtension), we limit the amount of extensions to 1.
	//
	// +optional
	// +kubebuilder:validation:MinItems=0
	// +kubebuilder:validation:MaxItems=1
	Extensions []commonv1alpha1.ExtensionRef `json:"extensions,omitempty"`

	// PluginsToInstall is a list of KongPluginInstallation resources that
	// will be installed and available in the Gateways (DataPlanes) that
	// use this GatewayConfig.
	// +optional
	PluginsToInstall []NamespacedName `json:"pluginsToInstall,omitempty"`
}

// GatewayConfigDataPlaneNetworkOptions defines network related options for a DataPlane.
// +apireference:kgo:include
type GatewayConfigDataPlaneNetworkOptions struct {
	// Services indicates the configuration of Kubernetes Services needed for
	// the topology of various forms of traffic (including ingress, etc.) to
	// and from the DataPlane.
	Services *GatewayConfigDataPlaneServices `json:"services,omitempty"`
}

// GatewayConfigDataPlaneServices contains Services related DataPlane configuration.
// +apireference:kgo:include
type GatewayConfigDataPlaneServices struct {
	// Ingress is the Kubernetes Service that will be used to expose ingress
	// traffic for the DataPlane. Here you can determine whether the DataPlane
	// will be exposed outside the cluster (e.g. using a LoadBalancer type
	// Services) or only internally (e.g. ClusterIP), and inject any additional
	// annotations you need on the service (for instance, if you need to
	// influence a cloud provider LoadBalancer configuration).
	//
	// +optional
	Ingress *GatewayConfigServiceOptions `json:"ingress,omitempty"`
}

// GatewayConfigServiceOptions is used to includes options to customize the ingress service,
// such as the annotations.
// +apireference:kgo:include
type GatewayConfigServiceOptions struct {
	ServiceOptions `json:",inline"`
}

// GatewayConfigurationStatus defines the observed state of GatewayConfiguration
// +apireference:kgo:include
type GatewayConfigurationStatus struct {
	// Conditions describe the current conditions of the GatewayConfigurationStatus.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayConfigurationList contains a list of GatewayConfiguration
// +apireference:kgo:include
type GatewayConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayConfiguration `json:"items"`
}

// GetConditions retrieves the GatewayConfiguration Status Condition
func (g *GatewayConfiguration) GetConditions() []metav1.Condition {
	return g.Status.Conditions
}

// SetConditions sets the GatewayConfiguration Status Condition
func (g *GatewayConfiguration) SetConditions(conditions []metav1.Condition) {
	g.Status.Conditions = conditions
}

// GetExtensions retrieves the GatewayConfiguration Extensions
func (g *GatewayConfiguration) GetExtensions() []commonv1alpha1.ExtensionRef {
	return g.Spec.Extensions
}
