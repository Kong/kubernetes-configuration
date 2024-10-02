package main

var supportedTypesControlPlaneConfig = []supportedTypesT{
	{
		Package: "v1",
		Types: []templateDataT{
			{
				Type:              "KongConsumer",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
		},
	},
	{
		Package: "v1beta1",
		Types: []templateDataT{
			{
				Type:              "KongConsumerGroup",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
		},
	},
	{
		Package: "v1alpha1",
		Types: []templateDataT{
			{
				Type:              "KongKey",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndKeySetRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongKeySet",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCredentialBasicAuth",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCredentialAPIKey",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCredentialJWT",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCredentialACL",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCACertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongCertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongPluginBinding",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongService",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongRoute",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongUpstream",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongTarget",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongVault",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				HasKonnectStatus:  true,
			},
			{
				Type:              "KongSNI",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndCertificateRefs",
				HasKonnectStatus:  true,
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
				HasKonnectStatus:  true,
			},
			{
				Type:              "KonnectAPIAuthConfiguration",
				KonnectStatusType: "*KonnectEntityStatus",
			},
		},
	},
}
