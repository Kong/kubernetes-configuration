<!-- This document is generated by KIC's 'generate.docs' make target, DO NOT EDIT -->

## Packages
- [konnect.konghq.com/v1alpha1](#konnectkonghqcomv1alpha1)


## konnect.konghq.com/v1alpha1

Package v1alpha1 contains API Schema definitions for the konnect.konghq.com v1alpha1 API group.

- [KonnectAPIAuthConfiguration](#konnectapiauthconfiguration)
- [KonnectCloudGatewayNetwork](#konnectcloudgatewaynetwork)
- [KonnectExtension](#konnectextension)
- [KonnectGatewayControlPlane](#konnectgatewaycontrolplane)
### KonnectAPIAuthConfiguration


KonnectAPIAuthConfiguration is the Schema for the Konnect configuration type.

<!-- konnect_api_auth_configuration description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectAPIAuthConfiguration`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectAPIAuthConfigurationSpec](#konnectapiauthconfigurationspec)_ | Spec is the specification of the KonnectAPIAuthConfiguration resource. |



### KonnectCloudGatewayNetwork


KonnectCloudGatewayNetwork is the Schema for the Konnect Network API.

<!-- konnect_cloud_gateway_network description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectCloudGatewayNetwork`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectCloudGatewayNetworkSpec](#konnectcloudgatewaynetworkspec)_ | Spec defines the desired state of KonnectCloudGatewayNetwork. |



### KonnectExtension


KonnectExtension is the Schema for the KonnectExtension API, and is intended to be referenced as
extension by the DataPlane, ControlPlane or GatewayConfiguration APIs.
If one of the above mentioned resources successfully refers a KonnectExtension, the underlying
deployment(s) spec gets customized to include the konnect-related configuration.

<!-- konnect_extension description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectExtension`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectExtensionSpec](#konnectextensionspec)_ | Spec is the specification of the KonnectExtension resource. |



### KonnectGatewayControlPlane


KonnectGatewayControlPlane is the Schema for the KonnectGatewayControlplanes API.

<!-- konnect_gateway_control_plane description placeholder -->

| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `konnect.konghq.com/v1alpha1`
| `kind` _string_ | `KonnectGatewayControlPlane`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[KonnectGatewayControlPlaneSpec](#konnectgatewaycontrolplanespec)_ | Spec defines the desired state of KonnectGatewayControlPlane. |



### Types

In this section you will find types that the CRDs rely on.
#### CertificateSecret


CertificateSecret contains the information to access the client certificate.



| Field | Description |
| --- | --- |
| `provisioning` _[ProvisioningMethod](#provisioningmethod)_ | Provisioning is the method used to provision the certificate. It can be either Manual or Automatic. In case manual provisioning is used, the certificate must be provided by the user. In case automatic provisioning is used, the certificate will be automatically generated by the system. |
| `secretRef` _[SecretRef](#secretref)_ | CertificateSecretRef is the reference to the Secret containing the client certificate. |


_Appears in:_
- [DataPlaneClientAuth](#dataplaneclientauth)

#### DataPlaneClientAuth


DataPlaneClientAuth contains the configuration for the client authentication for the DataPlane.
At the moment authentication is only supported through client certificate, but it might be extended in the future,
with e.g., token-based authentication.



| Field | Description |
| --- | --- |
| `certificateSecret` _[CertificateSecret](#certificatesecret)_ | CertificateSecret contains the information to access the client certificate. |


_Appears in:_
- [KonnectExtensionSpec](#konnectextensionspec)

#### DataPlaneClientAuthStatus


DataPlaneClientAuthStatus contains the status information related to the ClientAuth configuration.



| Field | Description |
| --- | --- |
| `certificateSecretRef` _[SecretRef](#secretref)_ | CertificateSecretRef is the reference to the Secret containing the client certificate. |


_Appears in:_
- [KonnectExtensionStatus](#konnectextensionstatus)

#### DataPlaneLabel


DataPlaneLabel contains the key-value pair of a label that will be applied to the Konnect DataPlane.



| Field | Description |
| --- | --- |
| `key` _string_ | Key is the key of the label. |
| `value` _string_ | Value is the value of the label. |


_Appears in:_
- [KonnectExtensionSpec](#konnectextensionspec)

#### KonnectAPIAuthConfigurationRef


KonnectAPIAuthConfigurationRef is a reference to a KonnectAPIAuthConfiguration resource.



| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the KonnectAPIAuthConfiguration resource. |


_Appears in:_
- [KonnectConfiguration](#konnectconfiguration)

#### KonnectAPIAuthConfigurationSpec


KonnectAPIAuthConfigurationSpec is the specification of the KonnectAPIAuthConfiguration resource.



| Field | Description |
| --- | --- |
| `type` _[KonnectAPIAuthType](#konnectapiauthtype)_ |  |
| `token` _string_ | Token is the Konnect token used to authenticate with the Konnect API. |
| `secretRef` _[SecretReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#secretreference-v1-core)_ | SecretRef is a reference to a Kubernetes Secret containing the Konnect token. This secret is required to have the konghq.com/credential label set to "konnect". |
| `serverURL` _string_ | ServerURL is the URL of the Konnect server. It can be either a full URL with an HTTPs scheme or just a hostname. Please refer to https://docs.konghq.com/konnect/network/ for the list of supported hostnames. |


_Appears in:_
- [KonnectAPIAuthConfiguration](#konnectapiauthconfiguration)



#### KonnectAPIAuthType
_Underlying type:_ `string`

KonnectAPIAuthType is the type of authentication used to authenticate with the Konnect API.





_Appears in:_
- [KonnectAPIAuthConfigurationSpec](#konnectapiauthconfigurationspec)

#### KonnectCloudGatewayNetworkSpec


KonnectCloudGatewayNetworkSpec defines the desired state of KonnectCloudGatewayNetwork.



| Field | Description |
| --- | --- |
| `name` _string_ | Specifies the name of the network on Konnect. |
| `cloud_gateway_provider_account_id` _string_ | Specifies the provider Account ID. |
| `region` _string_ | Region ID for cloud provider region. |
| `availability_zones` _string array_ | List of availability zones that the network is attached to. |
| `cidr_block` _string_ | CIDR block configuration for the network. |
| `state` _[NetworkCreateState](#networkcreatestate)_ | Initial state for creating a network. |
| `konnect` _[KonnectConfiguration](#konnectconfiguration)_ |  |


_Appears in:_
- [KonnectCloudGatewayNetwork](#konnectcloudgatewaynetwork)



#### KonnectConfiguration


KonnectConfiguration is the Schema for the KonnectConfiguration API.



| Field | Description |
| --- | --- |
| `authRef` _[KonnectAPIAuthConfigurationRef](#konnectapiauthconfigurationref)_ | APIAuthConfigurationRef is the reference to the API Auth Configuration that should be used for this Konnect Configuration. |


_Appears in:_
- [KonnectCloudGatewayNetworkSpec](#konnectcloudgatewaynetworkspec)
- [KonnectExtensionSpec](#konnectextensionspec)
- [KonnectGatewayControlPlaneSpec](#konnectgatewaycontrolplanespec)

#### KonnectEndpoints


KonnectEndpoints defines the Konnect endpoints for the control plane.



| Field | Description |
| --- | --- |
| `telemetry` _string_ | TelemetryEndpoint is the endpoint for telemetry. |
| `controlPlane` _string_ | ControlPlaneEndpoint is the endpoint for the control plane. |


_Appears in:_
- [KonnectGatewayControlPlaneStatus](#konnectgatewaycontrolplanestatus)

#### KonnectEntityStatus


KonnectEntityStatus represents the status of a Konnect entity.



| Field | Description |
| --- | --- |
| `id` _string_ | ID is the unique identifier of the Konnect entity as assigned by Konnect API. If it's unset (empty string), it means the Konnect entity hasn't been created yet. |
| `serverURL` _string_ | ServerURL is the URL of the Konnect server in which the entity exists. |
| `organizationID` _string_ | OrgID is ID of Konnect Org that this entity has been created in. |


_Appears in:_
- [KonnectCloudGatewayNetworkStatus](#konnectcloudgatewaynetworkstatus)
- [KonnectEntityStatusWithControlPlaneAndCertificateRefs](#konnectentitystatuswithcontrolplaneandcertificaterefs)
- [KonnectEntityStatusWithControlPlaneAndConsumerRefs](#konnectentitystatuswithcontrolplaneandconsumerrefs)
- [KonnectEntityStatusWithControlPlaneAndKeySetRef](#konnectentitystatuswithcontrolplaneandkeysetref)
- [KonnectEntityStatusWithControlPlaneAndServiceRefs](#konnectentitystatuswithcontrolplaneandservicerefs)
- [KonnectEntityStatusWithControlPlaneAndUpstreamRefs](#konnectentitystatuswithcontrolplaneandupstreamrefs)
- [KonnectEntityStatusWithControlPlaneRef](#konnectentitystatuswithcontrolplaneref)
- [KonnectGatewayControlPlaneStatus](#konnectgatewaycontrolplanestatus)













#### KonnectExtensionSpec


KonnectExtensionSpec defines the desired state of KonnectExtension.



| Field | Description |
| --- | --- |
| `controlPlaneRef` _[ControlPlaneRef](#controlplaneref)_ | ControlPlaneRef is a reference to a Konnect ControlPlane this KonnectExtension is associated with. |
| `dataPlaneClientAuth` _[DataPlaneClientAuth](#dataplaneclientauth)_ | DataPlaneClientAuth is the configuration for the client certificate authentication for the DataPlane. In case the ControlPlaneRef is of type KonnectID, it is required to set up the connection with the Konnect Platform. |
| `konnect` _[KonnectConfiguration](#konnectconfiguration)_ | KonnectConfiguration holds the information needed to setup the Konnect Configuration. |
| `dataPlaneLabels` _[DataPlaneLabel](#dataplanelabel) array_ | DataPlaneLabels is a set of labels that will be applied to the Konnect DataPlane. |


_Appears in:_
- [KonnectExtension](#konnectextension)



#### KonnectGatewayControlPlaneSpec


KonnectGatewayControlPlaneSpec defines the desired state of KonnectGatewayControlPlane.



| Field | Description |
| --- | --- |
| `name` _string_ | The name of the control plane. |
| `description` _string_ | The description of the control plane in Konnect. |
| `cluster_type` _[CreateControlPlaneRequestClusterType](#createcontrolplanerequestclustertype)_ | The ClusterType value of the cluster associated with the Control Plane. |
| `auth_type` _[AuthType](#authtype)_ | The auth type value of the cluster associated with the Runtime Group. |
| `cloud_gateway` _boolean_ | Whether this control-plane can be used for cloud-gateways. |
| `proxy_urls` _[ProxyURL](#proxyurl) array_ | Array of proxy URLs associated with reaching the data-planes connected to a control-plane. |
| `labels` _object (keys:string, values:string)_ | Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.<br /><br /> Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_". |
| `members` _[LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#localobjectreference-v1-core) array_ | Members is a list of references to the KonnectGatewayControlPlaneMembers that are part of this control plane group. Only applicable for ControlPlanes that are created as groups. |
| `konnect` _[KonnectConfiguration](#konnectconfiguration)_ |  |


_Appears in:_
- [KonnectGatewayControlPlane](#konnectgatewaycontrolplane)



#### ProvisioningMethod
_Underlying type:_ `string`

ProvisioningMethod is the type of the provisioning methods available to provision the certificate.





_Appears in:_
- [CertificateSecret](#certificatesecret)

#### SecretRef


SecretRef contains the reference to the Secret containing the Konnect Control Plane's cluster certificate.



| Field | Description |
| --- | --- |
| `name` _string_ | Name is the name of the Secret containing the Konnect Control Plane's cluster certificate. |


_Appears in:_
- [CertificateSecret](#certificatesecret)
- [DataPlaneClientAuthStatus](#dataplaneclientauthstatus)

