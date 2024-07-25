/*
Copyright 2024 Kong, Inc.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TODO(pmalek): expression field is not in OpenAPI spec but error message references it:
//   when protocols has 'http', at least one of 'hosts', 'methods', 'paths', 'headers' or 'expression' must be set

// KongPluginBinding is the schema for Plugin Bindings API which defines a Kong Plugin Binding.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Plugin-kind",type=string,JSONPath=`.spec.pluginReference.kind`,description="Kind of the plugin"
// +kubebuilder:printcolumn:name="Plugin-name",type=string,JSONPath=`.spec.pluginReference.name`,description="Name of the plugin"
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
type KongPluginBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KongPluginBindingSpec   `json:"spec"`
	Status KongPluginBindingStatus `json:"status,omitempty"`
}

func (c *KongPluginBinding) GetKonnectStatus() *KonnectEntityStatus {
	return &c.Status.Konnect.KonnectEntityStatus
}

func (c KongPluginBinding) GetTypeName() string {
	return "KongPluginBinding"
}

// GetConditions returns the Status Conditions
func (c *KongPluginBinding) GetConditions() []metav1.Condition {
	return c.Status.Conditions
}

// SetConditions sets the Status Conditions
func (c *KongPluginBinding) SetConditions(conditions []metav1.Condition) {
	c.Status.Conditions = conditions
}

// KongPluginBindingSpec defines specification of a KongPluginBinding.
type KongPluginBindingSpec struct {
	// PluginReference is a reference to the KongPlugin or KongClusterPlugin resource. It is required
	PluginReference PluginRef `json:"pluginRef"`

	// Kong contains the Kong entity references.
	Kong *KongReferences `json:"kong,omitempty"`

	// Let's defer this one to the future as we are not sure about the shape we want to give it.
	// EntityReference        *GenericEntityRef `json:"genericEntityRef,omitempty"`
}

type KongReferences struct {
	// TODO(mlavacca): add validation that at least one of the following field is properly set.
	// TODO(mlavacca): add validation to ensure that when genericEntityReference is set, no other reference is set.

	RouteReference         *EntityRef `json:"routeRef,omitempty"`
	ServiceReference       *EntityRef `json:"serviceRef,omitempty"`
	ConsumerReference      *EntityRef `json:"consumerRef,omitempty"`
	ConsumerGroupReference *EntityRef `json:"consumerGroupRef,omitempty"`
}

type PluginRef struct {
	// TODO(mattia): What about cross-namespace references? Do we want to introduce a namespace field
	// to allow such a reference? We could allow it and require a RefGrant.

	// Name is the name of the KongPlugin or KongClusterPlugin resource.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// kind can be KongPlugin or KongClusterPlugin. If not set, it is assumed to be KongPlugin.
	// +kubebuilder:validation:Enum=KongPlugin;KongClusterPlugin
	// +kubebuilder:default:=KongPlugin
	Kind *string `json:"kind,omitempty"`
}

type EntityRef struct {
	// Name is the name of the entity.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

type GenericEntityRef struct {
	// Name is the name of the generic entity.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// kind is the kind of the entity.
	// +kubebuilder:validation:Enum=Service;HTTPRoute;GCPRoute;TLSRoute;TCPRoute;UDPRoute;Ingress
	Kind string `json:"kind"`

	// TODO(mlavacca): add cross-field validation. Kind can be set depending on the group.
	// Group is the group of the entity.
	// +kubebuilder:validation:Enum="";core;gateway.networking.k8s.io;networking.k8s.io
	Group string `json:"group"`
}

// KongPluginBindingStatus represents the current status of the KongBinding resource.
type KongPluginBindingStatus struct {
	// Konnect contains the Konnect entity status.
	Konnect KonnectEntityStatusWithControlPlaneAndServiceRefs `json:"konnect,omitempty"`

	// Conditions describe the status of the Konnect entity.
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true

// KongPluginBindingList contains a list of KongPluginBindings.
type KongPluginBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KongPluginBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KongPluginBinding{}, &KongPluginBindingList{})
}
