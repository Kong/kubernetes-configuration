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

// KongKeySet is the schema for KongKeySet API which defines a KongKeySet entity.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
// +kubebuilder:validation:XValidation:rule="(!self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True')) ? true : oldSelf.spec.controlPlaneRef == self.spec.controlPlaneRef", message="spec.controlPlaneRef is immutable when an entity is already Programmed"
// +kubebuilder:validation:XValidation:rule="!has(self.spec.controlPlaneRef.konnectNamespacedRef) ? true : !has(self.spec.controlPlaneRef.konnectNamespacedRef.__namespace__)", message="spec.controlPlaneRef cannot specify namespace for namespaced resource - it's not supported yet"
// +kgosupported
type KongKeySet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec KongKeySetSpec `json:"spec"`

	// +kubebuilder:default={conditions: {{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}}
	Status KongKeySetStatus `json:"status,omitempty"`
}

// KongKeySetSpec defines the spec for a KongKeySet.
// +kgosupported
type KongKeySetSpec struct {
	// ControlPlaneRef is a reference to a Konnect ControlPlane with which KongKeySet is associated.
	// +optional
	ControlPlaneRef *ControlPlaneRef `json:"controlPlaneRef,omitempty"`

	// KongKeySetAPISpec are the attributes of the KongKeySet itself.
	KongKeySetAPISpec `json:",inline"`
}

// KongKeySetAPISpec defines the attributes of a Kong KeySet.
// +kgosupported
type KongKeySetAPISpec struct {
	// Name is a name of the KeySet.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Tags is an optional set of strings associated with the KeySet for grouping and filtering.
	// +optional
	Tags []string `json:"tags,omitempty"`
}

// KongKeySetStatus defines the status for a KongKeySet.
// +kgosupported
type KongKeySetStatus struct {
	// Konnect contains the Konnect entity status.
	// +optional
	Konnect *konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef `json:"konnect,omitempty"`

	// Conditions describe the status of the Konnect entity.
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// KongKeySetList contains a list of Kong Keys.
// +kubebuilder:object:root=true
// +kgosupported
type KongKeySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KongKeySet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KongKeySet{}, &KongKeySetList{})
}
