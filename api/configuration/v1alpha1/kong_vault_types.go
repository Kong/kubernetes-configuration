/*
Copyright 2023 Kong, Inc.

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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

const (
	// KongVaultKind is the kind name of KongVault resource.
	KongVaultKind = "KongVault"
)

// KongVault is the schema for kongvaults API which defines a custom Kong vault.
// A Kong vault is a storage to store sensitive data, where the values can be referenced in configuration of plugins.
// See: https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/
//
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster,shortName=kv,categories=kong-ingress-controller;kong,path=kongvaults
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Backend Type",type=string,JSONPath=`.spec.backend`,description="Name of the backend of the vault"
// +kubebuilder:printcolumn:name="Prefix",type=string,JSONPath=`.spec.prefix`,description="Prefix of vault URI to reference the values in the vault"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`,description="Age"
// +kubebuilder:printcolumn:name="Description",type=string,JSONPath=`.spec.description`,description="Description",priority=1
// +kubebuilder:printcolumn:name="Programmed",type=string,JSONPath=`.status.conditions[?(@.type=="Programmed")].status`
// +kubebuilder:validation:XValidation:rule="self.spec.prefix == oldSelf.spec.prefix", message="The spec.prefix field is immutable"
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.controlPlaneRef) || has(self.spec.controlPlaneRef)", message="controlPlaneRef is required once set"
// +kubebuilder:validation:XValidation:rule="(!has(self.status) || !self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True') || !has(self.spec.controlPlaneRef)) ? true : oldSelf.spec.controlPlaneRef == self.spec.controlPlaneRef", message="spec.controlPlaneRef is immutable when an entity is already Programmed"
// +apireference:kgo:include
// +apireference:kic:include
// +kong:channels=ingress-controller;gateway-operator
type KongVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KongVaultSpec   `json:"spec"`
	Status            KongVaultStatus `json:"status,omitempty"`
}

// KongVaultSpec defines specification of a custom Kong vault.
// +apireference:kgo:include
// +apireference:kic:include
type KongVaultSpec struct {
	// Backend is the type of the backend storing the secrets in the vault.
	// The supported backends of Kong is listed here:
	// https://docs.konghq.com/gateway/latest/kong-enterprise/secrets-management/backends/
	// +kubebuilder:validation:MinLength=1
	Backend string `json:"backend"`
	// Prefix is the prefix of vault URI for referencing values in the vault.
	// It is immutable after created.
	// +kubebuilder:validation:MinLength=1
	Prefix string `json:"prefix"`
	// Description is the additional information about the vault.
	Description string `json:"description,omitempty"`
	// Config is the configuration of the vault. Varies for different backends.
	Config apiextensionsv1.JSON `json:"config,omitempty"`
	// Tags are the tags associated to the vault for grouping and filtering.
	Tags commonv1alpha1.Tags `json:"tags,omitempty"`
	// ControlPlaneRef is a reference to a Konnect ControlPlane this KongVault is associated with.
	// +kubebuilder:validation:XValidation:message="'konnectID' type is not supported", rule="self.type != 'konnectID'"
	// +optional
	ControlPlaneRef *commonv1alpha1.ControlPlaneRef `json:"controlPlaneRef,omitempty"`
}

// KongVaultStatus represents the current status of the KongVault resource.
// +apireference:kgo:include
// +apireference:kic:include
type KongVaultStatus struct {
	// Konnect contains the Konnect entity status.
	// +optional
	// +apireference:kic:exclude
	Konnect *konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef `json:"konnect,omitempty"`

	// Conditions describe the current conditions of the KongVaultStatus.
	//
	// Known condition types are:
	//
	// * "Programmed"
	//
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions"`
}

// KongVaultList contains a list of KongVault.
// +kubebuilder:object:root=true
// +apireference:kgo:include
// +apireference:kic:include
type KongVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KongVault `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KongVault{}, &KongVaultList{})
}
