package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
)

func init() {
	SchemeBuilder.Register(&KonnectNetwork{}, &KonnectNetworkList{})
}

// KonnectNetwork is the Schema for the Konnect Network API.
//
// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
// +kubebuilder:printcolumn:name="ID",description="Konnect ID",type=string,JSONPath=`.status.id`
// +kubebuilder:printcolumn:name="OrgID",description="Konnect Organization ID this resource belongs to.",type=string,JSONPath=`.status.organizationID`
// +apireference:kgo:include
// +kong:channels=gateway-operator
type KonnectNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of KonnectNetwork.
	Spec KonnectNetworkSpec `json:"spec,omitempty"`

	// Status defines the observed state of KonnectNetwork.
	Status KonnectNetworkStatus `json:"status,omitempty"`
}

// KonnectNetworkSpec defines the desired state of KonnectNetwork.
// +apireference:kgo:include
type KonnectNetworkSpec struct {
	sdkkonnectcomp.CreateNetworkRequest `json:",inline"`

	// +kubebuilder:validation:Required
	KonnectConfiguration KonnectConfiguration `json:"konnect,omitempty"`
}

// KonnectNetworkStatus defines the observed state of KonnectNetwork.
// +apireference:kgo:include
type KonnectNetworkStatus struct {
	KonnectEntityStatus `json:",inline"`

	// Conditions describe the current conditions of the KonnectNetwork.
	//
	// Known condition types are:
	//
	// * "Programmed"
	//
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	// +kubebuilder:default={{type: "Programmed", status: "Unknown", reason:"Pending", message:"Waiting for controller", lastTransitionTime: "1970-01-01T00:00:00Z"}}
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// GetKonnectAPIAuthConfigurationRef returns the Konnect API Auth Configuration Ref.
func (c *KonnectNetwork) GetKonnectAPIAuthConfigurationRef() KonnectAPIAuthConfigurationRef {
	return c.Spec.KonnectConfiguration.APIAuthConfigurationRef
}

// KonnectNetworkList contains a list of KonnectNetwork.
// +kubebuilder:object:root=true
// +apireference:kgo:include
type KonnectNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectNetwork `json:"items"`
}
