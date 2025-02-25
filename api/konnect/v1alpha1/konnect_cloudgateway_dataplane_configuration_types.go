package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
)

func init() {
	SchemeBuilder.Register(&KonnectCloudGatewayDataPlaneGroupConfiguration{}, &KonnectCloudGatewayDataPlaneGroupConfigurationList{})
}

// KonnectCloudGatewayDataPlaneGroupConfiguration is the Schema for the Konnect Network API.
//
// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Programmed",description="The Resource is Programmed on Konnect",type=string,JSONPath=`.status.conditions[?(@.type=='Programmed')].status`
// +kubebuilder:printcolumn:name="ID",description="Konnect ID",type=string,JSONPath=`.status.id`
// +kubebuilder:printcolumn:name="ControlPlaneID",description="ControlPlane ID",type=string,JSONPath=`.status.controlPlaneID`
// +kubebuilder:printcolumn:name="OrgID",description="Konnect Organization ID this resource belongs to.",type=string,JSONPath=`.status.organizationID`
// +apireference:kgo:include
// +kong:channels=gateway-operator
type KonnectCloudGatewayDataPlaneGroupConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of KonnectCloudGatewayDataPlaneGroupConfiguration.
	Spec KonnectCloudGatewayDataPlaneGroupConfigurationSpec `json:"spec"`

	// Status defines the observed state of KonnectCloudGatewayDataPlaneGroupConfiguration.
	Status KonnectCloudGatewayDataPlaneGroupConfigurationStatus `json:"status,omitempty"`
}

// KonnectCloudGatewayDataPlaneGroupConfigurationSpec defines the desired state of KonnectCloudGatewayDataPlaneGroupConfiguration.
//
// +apireference:kgo:include
type KonnectCloudGatewayDataPlaneGroupConfigurationSpec struct {
	// Version specifies the desired Kong Gateway version.
	//
	// +kubebuilder:validation:Required
	Version string `json:"version"`

	// DataplaneGroups is a list of desired data-plane groups that describe where
	// to deploy instances, along with how many instances.
	//
	// +kubebuilder:validation:Optional
	DataplaneGroups []KonnectConfigurationDataPlaneGroup `json:"dataplane_groups"`

	// APIAccess is the desired type of API access for data-plane groups.
	//
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=private+public
	// +kubebuilder:validation:Enum=private;public;private+public
	APIAccess *sdkkonnectcomp.APIAccess `json:"api_access"`

	// ControlPlaneRef is a reference to a ControlPlane which DataPlanes from this
	// configuration will connect to.
	//
	// +kubebuilder:validation:Required
	ControlPlaneRef commonv1alpha1.ControlPlaneRef `json:"controlPlaneRef"`
}

// APIAccess defines the API access type for data-plane groups.
type APIAccess string

const (
	// APIAccessPrivate is the API access type for private data-plane groups.
	APIAccessPrivate APIAccess = "private"
	// APIAccessPublic is the API access type for public data-plane groups.
	APIAccessPublic APIAccess = "public"
	// APIAccessPrivatePublic is the API access type for private and public data-plane groups.
	APIAccessPrivatePublic APIAccess = "private+public"
)

const (
	// NetworkRefKonnectID is the type for the KonnectID Network ref.
	// It is used to reference a Konnect network entity by its ID on the Konnect platform.
	NetworkRefKonnectID = "konnectID"
)

// NetworkRef is the schema for the NetworkRef type.
// It is used to reference a Network entity.
//
// +kubebuilder:object:generate=true
// +kubebuilder:validation:XValidation:rule="(has(self.type) && self.type == 'konnectID') ? has(self.konnectID) : true", message="when type is konnectID, konnectID must be set"
// +apireference:kgo:include
type NetworkRef struct {
	// Type indicates the type of the control plane being referenced.
	//
	// +kubebuilder:validation:Enum=konnectID
	Type string `json:"type,omitempty"`

	// KonnectID is the schema for the KonnectID type.
	// This field is required when the Type is konnectID.
	// +optional
	KonnectID *string `json:"konnectID,omitempty"`
}

// KonnectConfigurationDataPlaneGroup is the schema for the KonnectConfiguration type.
type KonnectConfigurationDataPlaneGroup struct {
	// Name of cloud provider.
	//
	// +kubebuilder:validation:Required
	Provider sdkkonnectcomp.ProviderName `json:"provider"`

	// Region for cloud provider region.
	//
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	// NetworkRef is the schema for the NetworkRef type.
	//
	// +kubebuilder:validation:Required
	NetworkRef NetworkRef `json:"networkRef"`

	// Autoscale configuration for the data-plane group.
	//
	// +kubebuilder:validation:Required
	Autoscale ConfigurationDataPlaneGroupAutoscale `json:"autoscale"`

	// Array of environment variables to set for a data-plane group.
	//
	// +kubebuilder:validation:Optional
	Environment []ConfigurationDataPlaneGroupEnvironmentField `json:"environment,omitempty"`
}

// ConfigurationDataPlaneGroupEnvironmentField specifies an environment variable field for the data-plane group.
type ConfigurationDataPlaneGroupEnvironmentField struct {
	// Name of the environment variable field to set for the data-plane group. Must be prefixed by KONG_.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^KONG_."
	Name string `json:"name"`
	// Value assigned to the environment variable field for the data-plane group.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Required
	Value string `json:"value"`
}

// ConfigurationDataPlaneGroupAutoscale specifies the autoscale configuration for the data-plane group.
//
// +kubebuilder:validation:XValidation:rule="!(has(self.autopilot) && has(self.static))",message="can't provide both autopilot and static configuration"
// +kubebuilder:validation:XValidation:rule="self.type == 'ConfigurationDataPlaneGroupAutoscaleStatic' ? has(self.static) : true",message="static is required when type is ConfigurationDataPlaneGroupAutoscaleStatic"
// +kubebuilder:validation:XValidation:rule="self.type == 'ConfigurationDataPlaneGroupAutoscaleAutopilot' ? has(self.autopilot) : true",message="autopilot is required when type is ConfigurationDataPlaneGroupAutoscaleAutopilot"
type ConfigurationDataPlaneGroupAutoscale struct {
	// ConfigurationDataPlaneGroupAutoscaleStatic specifies the static configuration for the data-plane group.
	//
	// +kubebuilder:validation:Optional
	ConfigurationDataPlaneGroupAutoscaleStatic *ConfigurationDataPlaneGroupAutoscaleStatic `json:"static,omitempty"`

	// ConfigurationDataPlaneGroupAutoscaleAutopilot specifies the autoscale configuration for the data-plane group.
	//
	// +kubebuilder:validation:Optional
	ConfigurationDataPlaneGroupAutoscaleAutopilot *ConfigurationDataPlaneGroupAutoscaleAutopilot `json:"autopilot,omitempty"`

	// Type of autoscaling to use.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ConfigurationDataPlaneGroupAutoscaleStatic;ConfigurationDataPlaneGroupAutoscaleAutopilot
	Type sdkkonnectcomp.ConfigurationDataPlaneGroupAutoscaleType `json:"type"`
}

// ConfigurationDataPlaneGroupAutoscaleAutopilot specifies the autoscale configuration for the data-plane group.
type ConfigurationDataPlaneGroupAutoscaleAutopilot struct {
	// Base number of requests per second that the deployment target should support.
	//
	// +kubebuilder:validation:Required
	BaseRps int64 `json:"base_rps"`

	// Max number of requests per second that the deployment target should support. If not set, this defaults to 10x base_rps.
	MaxRps *int64 `json:"max_rps,omitempty"`
}

// ConfigurationDataPlaneGroupAutoscaleStatic specifies the static configuration for the data-plane group.
type ConfigurationDataPlaneGroupAutoscaleStatic struct {
	// Instance type name to indicate capacity.
	//
	// +kubebuilder:validation:Required
	InstanceType sdkkonnectcomp.InstanceTypeName `json:"instance_type"`

	// Number of data-planes the deployment target will contain.
	//
	// +kubebuilder:validation:Required
	RequestedInstances int64 `json:"requested_instances"`
}

// KonnectCloudGatewayDataPlaneGroupConfigurationStatus defines the observed state of KonnectCloudGatewayDataPlaneGroupConfiguration.
// +apireference:kgo:include
type KonnectCloudGatewayDataPlaneGroupConfigurationStatus struct {
	KonnectEntityStatusWithControlPlaneRef `json:",inline"`

	// DataPlaneGroups is a list of deployed data-plane groups.
	DataPlaneGroups []KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup `json:"dataplane_groups,omitempty"`

	// Conditions describe the current conditions of the KonnectCloudGatewayDataPlaneGroupConfiguration.
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

// KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup defines the observed state of a deployed data-plane group.
type KonnectCloudGatewayDataPlaneGroupConfigurationStatusGroup struct {
	// ID is the ID of the deployed data-plane group.
	//
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty"`

	// CloudGatewayNetworkID is the ID of the cloud gateway network.
	//
	// +kubebuilder:validation:Optional
	CloudGatewayNetworkID *string `json:"cloud_gateway_network_id,omitempty"`

	// Name of cloud provider.
	//
	// +kubebuilder:validation:Optional
	Provider sdkkonnectcomp.ProviderName `json:"provider,omitempty"`

	// Region ID for cloud provider region.
	//
	// +kubebuilder:validation:Optional
	Region string `json:"region,omitempty"`

	// PrivateIPAddresses is a list of private IP addresses of the internal load balancer that proxies traffic to this data-plane group.
	PrivateIPAddresses []string `json:"private_ip_addresses,omitempty"`

	// EgressIPAddresses is a list of egress IP addresses for the network that this data-plane group runs on.
	EgressIPAddresses []string `json:"egress_ip_addresses,omitempty"`

	// State is the current state of the data plane group. Can be e.g. initializing, ready, terminating.
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty"`
}

// KonnectCloudGatewayDataPlaneGroupConfigurationList contains a list of KonnectCloudGatewayDataPlaneGroupConfiguration.
// +kubebuilder:object:root=true
// +apireference:kgo:include
type KonnectCloudGatewayDataPlaneGroupConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectCloudGatewayDataPlaneGroupConfiguration `json:"items"`
}
