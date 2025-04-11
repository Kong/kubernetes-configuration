package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
)

func init() {
	SchemeBuilder.Register(&KonnectCloudGatewayTransitGateway{}, &KonnectCloudGatewayTransitGatewayList{})
}

// KonnectCloudGatewayTransitGateway is the Schema for the Konnect Transit Gateway API.
//
// +genclient
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:categories=kong
// +kubebuilder:object:root=true
// +kubebuilder:object:generate=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:finalizer
// +apireference:kgo:include
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`,description="State of the transit gateway in Konnect"
// +kubebuilder:validation:XValidation:rule="(!has(self.status) || !self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True')) ? true : (!has(self.spec.awsTransitGateway) ? true : oldSelf.spec.awsTransitGateway.name == self.spec.awsTransitGateway.name)",message="spec.awsTransitGateway.name is immutable when transit gateway is already Programmed"
// +kubebuilder:validation:XValidation:rule="(!has(self.status) || !self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True')) ? true : (!has(self.spec.azureTransitGateway) ? true : oldSelf.spec.azureTransitGateway.name == self.spec.azureTransitGateway.name)",message="spec.azureTransitGateway.name is immutable when transit gateway is already Programmed"
// +kong:channels=gateway-operator
type KonnectCloudGatewayTransitGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of KonnectCloudGatewayTransitGateway.
	Spec KonnectCloudGatewayTransitGatewaySpec `json:"spec"`

	// Status defines the observed state of KonnectCloudGatewayTransitGateway.
	Status KonnectCloudGatewayTransitGatewayStatus `json:"status,omitempty"`
}

// KonnectCloudGatewayTransitGatewaySpec defines the desired state of KonnectCloudGatewayTransitGateway.
//
// +kubebuilder:validation:XValidation:rule="self.networkRef.type == 'namespacedRef'", message = "only namespacedRef is supported currently"
type KonnectCloudGatewayTransitGatewaySpec struct {
	// NetworkRef is the schema for the NetworkRef type.
	//
	// +kubebuilder:validation:Required
	NetworkRef commonv1alpha1.ObjectRef `json:"networkRef"`
	// KonnectTransitGatewayAPISpec is the configuration of the transit gateway on Konnect side.
	KonnectTransitGatewayAPISpec `json:",inline"`
}

// TransitGatewayType defines the type of Konnect transit gateway.
type TransitGatewayType string

const (
	// TransitGatewayTypeAWSTransitGateway defines the the AWS transit gateway type.
	TransitGatewayTypeAWSTransitGateway TransitGatewayType = "awsTransitGateway"
	// TransitGatewayTypeAzureTransitGateway defines the Azure transit gateway type.
	TransitGatewayTypeAzureTransitGateway TransitGatewayType = "azureTransitGateway"
)

// KonnectTransitGatewayAPISpec specifies a transit gateway on the Konnect side.
// The type and all the types it referenced are mostly copied github.com/Kong/sdk-konnect-go/models/components.CreateTransitGatewayRequest.
//
// +kubebuilder:validation:XValidation:rule="self.type == oldSelf.type", message="spec.type is immutable"
// +kubebuilder:validation:XValidation:rule="self.type == 'awsTransitGateway' ? has(self.awsTransitGateway) : true", message="must set spec.awsTransitGateway when spec.type is 'awsTransitGateway'"
// +kubebuilder:validation:XValidation:rule="self.type != 'awsTransitGateway' ? !has(self.awsTransitGateway) : true", message="must not set spec.awsTransitGateway when spec.type is not 'awsTransitGateway'"
// +kubebuilder:validation:XValidation:rule="self.type == 'azureTransitGateway' ? has(self.azureTransitGateway) : true", message = "must set spec.azureTransitGateway when spec.type is 'azureTransitGateway'"
// +kubebuilder:validation:XValidation:rule="self.type != 'azureTransitGateway' ? !has(self.azureTransitGateway) : true", message = "must not set spec.azureTransitGateway when spec.type is not 'azureTransitGateway'"
// TODO: add more constraints on attachment_config based on type.
// REVIEW: should we set "immutable after programmed" on spec.*TransitGateway.name?
type KonnectTransitGatewayAPISpec struct {
	// Type is the type of the Konnect transit gateway.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=awsTransitGateway;azureTransitGateway
	Type TransitGatewayType `json:"type"`

	// AWSTransitGateway is the configuration of an AWS transit gateway.
	// Used when type is "AWS Transit Gateway".
	//
	// +kubebuilder:validation:Optional
	AWSTransitGateway *AWSTransitGateway `json:"awsTransitGateway,omitempty"`
	// AzureTransitGateway is the configuration of an Azure transit gateway.
	// Used when type is "Azure Transit Gateway".
	//
	// +kubebuilder:validation:Optional
	AzureTransitGateway *AzureTransitGateway `json:"azureTransitGateway,omitempty"`
}

// AWSTransitGateway is the configuration of an AWS transit gateway.
type AWSTransitGateway struct {
	// Human-readable name of the transit gateway.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=120
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway
	// attachment.
	//
	// +kubebuilder:validation:Optional
	DNSConfig []TransitGatewayDNSConfig `json:"dns_config,omitempty"`
	// CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning
	// network.
	//
	// +kubebuilder:validation:Required
	CIDRBlocks []string `json:"cidr_blocks"`
	// configuration to attach to AWS transit gateway on the AWS side.
	//
	// +kubebuilder:validation:Required
	AttachmentConfig AwsTransitGatewayAttachmentConfig `json:"attachment_config"`
}

// AzureTransitGateway is the configuration of an Azure transit gateway.
type AzureTransitGateway struct {
	// Human-readable name of the transit gateway.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=120
	// +kubebuilder:validation:Required
	Name string `json:"name"`
	// List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway
	// attachment.
	//
	// +kubebuilder:validation:Optional
	DNSConfig []TransitGatewayDNSConfig `json:"dns_config,omitempty"`
	// configuration to attach to Azure VNET peering gateway.
	//
	// +kubebuilder:validation:Required
	AttachmentConfig AzureVNETPeeringAttachmentConfig `json:"attachment_config"`
}

// TransitGatewayDNSConfig is the DNS configuration of a tansit gateway.
type TransitGatewayDNSConfig struct {
	// Remote DNS Server IP Addresses to connect to for resolving internal DNS via a transit gateway.
	RemoteDNSServerIPAddresses []string `json:"remote_dns_server_ip_addresses,omitempty"`
	// Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,
	// for a transit gateway.
	DomainProxyList []string `json:"domain_proxy_list,omitempty"`
}

// AwsTransitGatewayAttachmentConfig is the configuration to attach to a AWS transit gateway.
type AwsTransitGatewayAttachmentConfig struct {
	// AWS Transit Gateway ID to create attachment to.
	//
	// +kubebuilder:validation:Required
	TransitGatewayID string `json:"transit_gateway_id"`
	// Resource Share ARN to verify request to create transit gateway attachment.
	//
	// +kubebuilder:validation:Required
	RAMShareArn string `json:"ram_share_arn"`
}

// AzureVNETPeeringAttachmentConfig is the configuration to attach to a Azure VNET peering gateway.
type AzureVNETPeeringAttachmentConfig struct {
	// Tenant ID for the Azure VNET Peering attachment.
	//
	// +kubebuilder:validation:Required
	TenantID string `json:"tenant_id"`
	// Subscription ID for the Azure VNET Peering attachment.
	//
	// +kubebuilder:validation:Required
	SubscriptionID string `json:"subscription_id"`
	// Resource Group Name for the Azure VNET Peering attachment.
	//
	// +kubebuilder:validation:Required
	ResourceGroupName string `json:"resource_group_name"`
	// VNET Name for the Azure VNET Peering attachment.
	//
	// +kubebuilder:validation:Required
	VnetName string `json:"vnet_name"`
}

// KonnectCloudGatewayTransitGatewayStatus defines the current state of KonnectCloudGatewayTransitGateway.
type KonnectCloudGatewayTransitGatewayStatus struct {
	KonnectEntityStatusWithNetworkRef `json:",inline"`
	// State is the state of the transit gateway on Konnect side.
	State sdkkonnectcomp.TransitGatewayState `json:"state,omitempty"`
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

// KonnectCloudGatewayTransitGatewayList contains a list of KonnectCloudGatewayTransitGateway.
// +kubebuilder:object:root=true
// +apireference:kgo:include
type KonnectCloudGatewayTransitGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KonnectCloudGatewayTransitGateway `json:"items"`
}
