package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
)

func init() {
	SchemeBuilder.Register(&KonnectCloudGatewayNetwork{}, &KonnectCloudGatewayNetworkList{})
}

// KonnectCloudGatewayNetwork is the Schema for the Konnect Network API.
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
type KonnectCloudGatewayNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of KonnectCloudGatewayNetwork.
	Spec KonnectCloudGatewayNetworkSpec `json:"spec,omitempty"`

	// Status defines the observed state of KonnectCloudGatewayNetwork.
	Status KonnectCloudGatewayNetworkStatus `json:"status,omitempty"`
}

// KonnectCloudGatewayNetworkSpec defines the desired state of KonnectCloudGatewayNetwork.
// +apireference:kgo:include
type KonnectCloudGatewayNetworkSpec struct {
	// NOTE: These fields are extracted from sdkkonnectcomp.CreateNetworkRequest
	// because for some reason when embedding the struct, the fields deserialization
	// doesn't work (konnect field is always empty).

	Name                          string `json:"name"`
	CloudGatewayProviderAccountID string `json:"cloud_gateway_provider_account_id"`
	// Region ID for cloud provider region.
	Region string `json:"region"`
	// List of availability zones that the network is attached to.
	AvailabilityZones []string `json:"availability_zones"`
	// CIDR block configuration for the network.
	CidrBlock string `json:"cidr_block"`
	// Initial state for creating a network.
	// +optional
	State *sdkkonnectcomp.NetworkCreateState `json:"state"`

	// +kubebuilder:validation:Required
	KonnectConfiguration KonnectConfiguration `json:"konnect"`
}

// KonnectCloudGatewayNetworkStatus defines the observed state of KonnectCloudGatewayNetwork.
// +apireference:kgo:include
type KonnectCloudGatewayNetworkStatus struct {
	KonnectEntityStatus `json:",inline"`

	// Conditions describe the current conditions of the KonnectCloudGatewayNetwork.
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
func (c *KonnectCloudGatewayNetwork) GetKonnectAPIAuthConfigurationRef() KonnectAPIAuthConfigurationRef {
	return c.Spec.KonnectConfiguration.APIAuthConfigurationRef
}

// KonnectCloudGatewayNetworkList contains a list of KonnectCloudGatewayNetwork.
// +kubebuilder:object:root=true
// +apireference:kgo:include
type KonnectCloudGatewayNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectCloudGatewayNetwork `json:"items"`
}
