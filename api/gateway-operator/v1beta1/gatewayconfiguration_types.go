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
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/common/v1alpha1"
)

func init() {
	SchemeBuilder.Register(&GatewayConfiguration{}, &GatewayConfigurationList{})
}

// GatewayConfiguration is the Schema for the gatewayconfigurations API.
//
// +genclient
// +apireference:kgo:include
// +kong:channels=gateway-operator
// +kubebuilder:deprecatedversion:warning="GatewayConfiguration v1beta1 has been deprecated in favor of v2beta1 and it will be removed in future."
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=kogc,categories=kong
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

	// +optional
	Resources *GatewayConfigDataPlaneResources `json:"resources,omitempty"`

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
	// The `port` in the ports of ingress service must be the same as the port
	// in one of the Gateway's listeners.
	//
	// +optional
	Ingress *GatewayConfigServiceOptions `json:"ingress,omitempty"`
}

// GatewayConfigDataPlaneResources defines the resources that will be
// created and managed for Gateway's DataPlane.
// +apireference:kgo:include
type GatewayConfigDataPlaneResources struct {

	// PodDisruptionBudget is the configuration for the PodDisruptionBudget
	// that will be created for the DataPlane.
	//
	// +optional
	PodDisruptionBudget *PodDisruptionBudget `json:"podDisruptionBudget,omitempty"`
}

// GatewayConfigServiceOptions is used to includes options to customize the ingress service,
// such as the annotations and ports.
// +apireference:kgo:include
// +kubebuilder:validation:XValidation:message="Cannot set NodePort when service type is not NodePort or LoadBalancer", rule="!has(self.ports) || !(self.ports.exists(p, has(p.nodePort))) ? true : has(self.type) && ['LoadBalancer', 'NodePort'].exists(t, t == self.type)"
type GatewayConfigServiceOptions struct {
	ServiceOptions `json:",inline"`

	// Ports defines the list of ports that are exposed by the service.
	// The ports field allows defining the name, port and targetPort of
	// the underlying service ports, while the protocol is defaulted to TCP,
	// as it is the only protocol currently supported.
	//
	// +kubebuilder:validation:MaxItems=4
	Ports []GatewayConfigurationServicePort `json:"ports,omitempty"`
}

// GatewayConfigurationServicePort contains information on service's port.
// +apireference:kgo:include
type GatewayConfigurationServicePort struct {
	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names. When considering
	// the endpoints for a Service, this must match the 'name' field in the
	// EndpointPort.
	// Optional if only one ServicePort is defined on this service.
	// +optional
	Name string `json:"name,omitempty"`

	// The port that will be exposed by this service.
	Port int32 `json:"port"`

	// Number or name of the port to access on the pods targeted by the service.
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// If this is a string, it will be looked up as a named port in the
	// target Pod's container ports. If this is not specified, the value
	// of the 'port' field is used (an identity map).
	// This field is ignored for services with clusterIP=None, and should be
	// omitted or set equal to the 'port' field.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	// +optional
	TargetPort intstr.IntOrString `json:"targetPort,omitempty"`

	// The port on each node on which this service is exposed when type is
	// NodePort or LoadBalancer. Usually assigned by the system. If a value is
	// specified, in-range, and not in use it will be used, otherwise the
	// operation will fail. If not specified, a port will be allocated if this
	// Service requires one. If this field is specified when creating a
	// Service which does not need it, creation will fail. This field will be
	// wiped when updating a Service to no longer need it (e.g. changing type
	// from NodePort to ClusterIP).
	//
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	// Can only be specified if type is NodePort or LoadBalancer.
	//
	// +optional
	NodePort int32 `json:"nodePort,omitempty"`
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

// ConvertTo converts this GatewayConfiguration (v1beta1) to the Hub version.
func (g *GatewayConfiguration) ConvertTo(dstRaw conversion.Hub) error {
	_ = dstRaw
	panic("v1beta1 GatewayConfiguration ConvertTo(...) is not implemented yet")
}

// ConvertFrom converts from the Hub version to this version (v1beta1).
func (g *GatewayConfiguration) ConvertFrom(srcRaw conversion.Hub) error {
	_ = srcRaw
	panic("v1beta1 GatewayConfiguration ConvertFrom(...) is not implemented yet")
}
