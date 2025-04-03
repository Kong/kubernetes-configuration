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
// +kubebuilder:validation:XValidation:rule="self.type == oldSelf.type", message="spec.type is immutable"
// +kubebuilder:validation:XValidation:rule="self.type == 'AWS Transit Gateway' ? has(self.awsTransitGateway) : true", message="must set spec.awsTransitGateway when spec.type is 'AWS Transit Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type != 'AWS Transit Gateway' ? !has(self.awsTransitGateway) : true", message="must not set spec.awsTransitGateway when spec.type is not 'AWS Transit Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type == 'AWS VPC Peering Gateway' ? has(self.awsVPCPeeringGateway) : true", message = "must set spec.awsVPCPeeringGateway when spec.type is 'AWS VPC Peering Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type != 'AWS VPC Peering Gateway' ? !has(self.awsVPCPeeringGateway) : true", message = "must not set spec.awsVPCPeeringGateway when spec.type is not 'AWS VPC Peering Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type == 'Azure Transit Gateway' ? has(self.azureTransitGateway) : true", message = "must set spec.azureTransitGateway when spec.type is 'Azure Transit Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type != 'Azure Transit Gateway' ? !has(self.azureTransitGateway) : true", message = "must not set spec.azureTransitGateway when spec.type is not 'Azure Transit Gateway'"
// +kubebuilder:validation:XValidation:rule="self.type == 'AWS Transit Gateway' ? (has(self.awsTransitGateway) && self.awsTransitGateway.attachment_config.kind == 'aws-transit-gateway-attachment') : true",message="must set spec.awsTransitGateway.attachment_config.kind to 'aws-transit-gateway-attachment' for AWS transit gateway type"
// +kubebuilder:validation:XValidation:rule="self.type == 'AWS VPC Peering Gateway' ? (has(self.awsVPCPeeringGateway) && self.awsVPCPeeringGateway.attachment_config.kind == 'aws-vpc-peering-attachment') : true",message="must set spec.awsVPCPeeringGateway.attachment_config.kind to 'aws-vpc-peering-attachment' for AWS VPC peering gateway type"
// +kubebuilder:validation:XValidation:rule="self.type == 'Azure Transit Gateway' ? (has(self.azureTransitGateway) && self.azureTransitGateway.attachment_config.kind == 'azure-vnet-peering-attachment') : true",message="must set spec.azureTransitGateway.attachment_config.kind to 'azure-vnet-peering-attachment' for Azure transit gateway type"
// TODO: add more constraints on attachment_config based on type
type KonnectCloudGatewayTransitGatewaySpec struct {
	// NetworkRef is the schema for the NetworkRef type.
	//
	// +kubebuilder:validation:Required
	NetworkRef commonv1alpha1.ObjectRef `json:"networkRef"`
	// KonnectTransitGatewayAPISpec is the configuration of the transit gateway on Konnect side.
	KonnectTransitGatewayAPISpec `json:",inline"`
}

// KonnectTransitGatewayAPISpec specifies a transit gateway on the Konnect side.
// The type and all the types it referenced are mostly copied github.com/Kong/sdk-konnect-go/models/components.CreateTransitGatewayRequest.
//
// TODO: add necessary comments for types/fields if they were missing in sdk-konnect-go.
type KonnectTransitGatewayAPISpec struct {
	// Type is the type of the Konnect transit gateway.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=AWS Transit Gateway;AWS VPC Peering Gateway;Azure Transit Gateway
	Type sdkkonnectcomp.CreateTransitGatewayRequestType `json:"type"`

	// AWSTransitGateway is the configuration of an AWS transit gateway.
	// Used when type is "AWS Transit Gateway".
	//
	// +kubebuilder:validation:Optional
	AWSTransitGateway *AWSTransitGateway `json:"awsTransitGateway,omitempty"`
	// AWSVPCPeeringGateway is the configuration of an AWS VPC peering gateway.
	// Used when type is "AWS VPC Peering Gateway".
	//
	// +kubebuilder:validation:Optional
	AWSVPCPeeringGateway *AWSVPCPeeringGateway `json:"awsVPCPeeringGateway,omitempty"`
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

// AWSVPCPeeringGateway is the configuration of an AWS VPC peering gateway.
type AWSVPCPeeringGateway struct {
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
	// configuration to attach to AWS VPC peering gateway on AWS side.
	//
	// +kubebuilder:validation:Required
	AttachmentConfig AwsVpcPeeringGatewayAttachmentConfig `json:"attachment_config"`
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
	RemoteDNSServerIPAddresses []string `json:"remote_dns_server_ip_addresses"`
	// Internal domain names to proxy for DNS resolution from the listed remote DNS server IP addresses,
	// for a transit gateway.
	DomainProxyList []string `json:"domain_proxy_list"`
}

// AwsTransitGatewayAttachmentConfig is the configuration to attach to a AWS transit gateway.
type AwsTransitGatewayAttachmentConfig struct {
	// +kubebuilder:validation:Required
	Kind sdkkonnectcomp.AWSTransitGatewayAttachmentType `json:"kind"`
	// AWS Transit Gateway ID to create attachment to.
	//
	// +kubebuilder:validation:Required
	TransitGatewayID string `json:"transit_gateway_id"`
	// Resource Share ARN to verify request to create transit gateway attachment.
	//
	// +kubebuilder:validation:Required
	RAMShareArn string `json:"ram_share_arn"`
}

// AwsVpcPeeringGatewayAttachmentConfig is the configuration to attach to a AWS VPC peering gateway.
type AwsVpcPeeringGatewayAttachmentConfig struct {
	// +kubebuilder:validation:Required
	Kind sdkkonnectcomp.AWSVPCPeeringAttachmentConfig `json:"kind"`
	// +kubebuilder:validation:Required
	PeerAccountID string `json:"peer_account_id"`
	// +kubebuilder:validation:Required
	PeerVpcID string `json:"peer_vpc_id"`
	// +kubebuilder:validation:Required
	PeerVpcRegion string `json:"peer_vpc_region"`
}

// AzureVNETPeeringAttachmentConfig is the configuration to attach to a Azure VNET peering gateway.
type AzureVNETPeeringAttachmentConfig struct {
	Kind sdkkonnectcomp.AzureVNETPeeringAttachmentType `json:"kind"`
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
