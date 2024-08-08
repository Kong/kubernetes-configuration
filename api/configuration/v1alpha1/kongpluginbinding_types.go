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

	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

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
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
type KongPluginBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:message="One field between global and targets must be set",rule="(((has(self.global) && self.global == true) ? !has(self.targets) : true) && (((!has(self.global) || self.global == false) ? has(self.targets) : true) || (!has(self.targets) ? (has(self.global) && self.global == true) : true)))"
	Spec   KongPluginBindingSpec   `json:"spec"`
	Status KongPluginBindingStatus `json:"status,omitempty"`
}

func (c *KongPluginBinding) GetKonnectStatus() *konnectv1alpha1.KonnectEntityStatus {
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

	// +optional
	Global *bool `json:"global,omitempty"`

	// +optional
	// +kubebuilder:validation:XValidation:message="The combination of entities set is not allowed",rule="(has(self.consumerRef) && has(self.routeRef) && has(self.serviceRef) && !has(self.consumerGroupRef)) || (has(self.consumerGroupRef) && has(self.serviceRef) && has(self.routeRef) && !has(self.consumerRef)) || (has(self.consumerRef) && has(self.routeRef) && !has(self.consumerGroupRef) && !has(self.serviceRef)) || (has(self.consumerRef) && has(self.serviceRef) && !has(self.routeRef) && !has(self.consumerGroupRef)) || (has(self.consumerGroupRef) && has(self.routeRef) && !has(self.serviceRef) && !has(self.consumerRef)) || (has(self.consumerGroupRef) && has(self.serviceRef) && !has(self.consumerRef) && !has(self.routeRef)) || (has(self.routeRef) && has(self.serviceRef) && !has(self.consumerRef) && !has(self.consumerGroupRef)) || (has(self.consumerRef) && !has(self.serviceRef) && !has(self.routeRef) && !has(self.consumerGroupRef)) || (has(self.consumerGroupRef) && !has(self.serviceRef) && !has(self.routeRef) && !has(self.consumerRef)) || (has(self.routeRef) && !has(self.serviceRef) && !has(self.consumerRef) && !has(self.consumerGroupRef)) || (has(self.serviceRef) && !has(self.routeRef) && !has(self.consumerGroupRef) && !has(self.consumerRef))"
	// +kubebuilder:validation:XValidation:message="At least one entity reference must be set",rule="has(self.routeRef) || has(self.serviceRef) || has(self.consumerRef) || has(self.consumerGroupRef)"
	// +kubebuilder:validation:XValidation:message="KongRoute can be used only when serviceRef is unset or set to KongService, and viceversa",rule="(has(self.routeRef) && self.routeRef.kind == 'KongRoute') ? (!has(self.serviceRef) || self.serviceRef.kind == 'KongService') : true"
	Targets *KongPluginBindingTargets `json:"targets,omitempty"`
}

type KongPluginBindingTargets struct {
	// Kong contains the Kong entity references. It is possible to set multiple combinations
	// of references, as described in https://docs.konghq.com/gateway/latest/key-concepts/plugins/#precedence
	// The complete set of allowed combinations and their order of precedence for plugins
	// configured to multiple entities is:
	//
	// 1. Consumer + route + service
	// 2. Consumer group + service + route
	// 3. Consumer + route
	// 4. Consumer + service
	// 5. Consumer group + route
	// 6. Consumer group + service
	// 7. Route + service
	// 8. Consumer
	// 9. Consumer group
	// 10. Route
	// 11. Service
	//
	// +optional
	// +kubebuilder:validation:XValidation:message="group/kind not allowed for the routeRef",rule="(self.kind == 'KongRoute' && self.group == 'configuration.konghq.com') || (self.kind == 'Ingress' && self.group == 'networking.k8s.io') || (self.kind == 'HTTPRoute' && self.group == 'gateway.networking.k8s.io') || (self.kind == 'GCPRoute' && self.group == 'gateway.networking.k8s.io')"
	RouteReference *TargetRefWithGroupKind `json:"routeRef,omitempty"`

	// +optional
	// +kubebuilder:validation:XValidation:message="group/kind not allowed for the serviceRef",rule="(self.kind == 'KongService' && self.group == 'configuration.konghq.com') || (self.kind == 'Service' && (self.group == '' || self.group == 'core'))"
	ServiceReference       *TargetRefWithGroupKind `json:"serviceRef,omitempty"`
	ConsumerReference      *TargetRef              `json:"consumerRef,omitempty"`
	ConsumerGroupReference *TargetRef              `json:"consumerGroupRef,omitempty"`
}

type PluginRef struct {
	// TODO(mattia): cross-namespace references are not supported yet.
	// https://github.com/Kong/kubernetes-configuration/issues/9

	// Name is the name of the KongPlugin or KongClusterPlugin resource.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// kind can be KongPlugin or KongClusterPlugin. If not set, it is assumed to be KongPlugin.
	// +kubebuilder:validation:Enum=KongPlugin;KongClusterPlugin
	// +kubebuilder:default:=KongPlugin
	Kind *string `json:"kind,omitempty"`
}

type TargetRef struct {
	// Name is the name of the entity.
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

// TODO(mlavacca): add xvalidation between group and kind
type TargetRefWithGroupKind struct {
	// Name is the name of the entity.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=KongService;KongRoute;Service;HTTPRoute;GCPRoute;Ingress
	Kind string `json:"kind"`

	// +kubebuilder:validation:Enum="";core;gateway.networking.k8s.io;networking.k8s.io;configuration.konghq.com
	Group string `json:"group"`
}

// KongPluginBindingStatus represents the current status of the KongBinding resource.
type KongPluginBindingStatus struct {
	// Konnect contains the Konnect entity status.
	// +optional
	Konnect *konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef `json:"konnect,omitempty"`

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
