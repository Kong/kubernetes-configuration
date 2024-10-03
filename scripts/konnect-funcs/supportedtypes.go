package main

var supportedTypesControlPlaneConfig = []supportedTypesT{
	{
		Package: "v1",
		Types: []templateDataT{
			{
				Type:              "KongConsumer",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
		},
	},
	{
		Package: "v1beta1",
		Types: []templateDataT{
			{
				Type:              "KongConsumerGroup",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
		},
	},
	{
		Package: "v1alpha1",
		Types: []templateDataT{
			{
				Type:              "KongKey",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndKeySetRef",
			},
			{
				Type:              "KongKeySet",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongCredentialBasicAuth",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
			},
			{
				Type:              "KongCredentialAPIKey",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
			},
			{
				Type:              "KongCredentialJWT",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
			},
			{
				Type:              "KongCredentialACL",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
			},
			{
				Type:              "KongCredentialHMAC",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
			},
			{
				Type:              "KongCACertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongCertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongPluginBinding",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongService",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongRoute",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs",
			},
			{
				Type:              "KongUpstream",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongTarget",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs",
			},
			{
				Type:              "KongVault",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
			{
				Type:              "KongSNI",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndCertificateRefs",
			},
			{
				Type:              "KongDataPlaneClientCertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
			},
		},
	},
}

var supportedTypesStandalone = []supportedTypesT{
	{
		Package: "v1alpha1",
		Types: []templateDataT{
			{
				Type:              "KonnectGatewayControlPlane",
				KonnectStatusType: "*KonnectEntityStatus",
			},
			{
				Type: "KonnectAPIAuthConfiguration",
			},
		},
	},
}
