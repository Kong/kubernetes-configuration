package main

var supportedKonnectTypesControlPlaneConfig = []supportedTypesT{
	{
		PackageVersion: "v1",
		AdditionalImports: []string{
			`configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                "KongConsumer",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "configurationv1alpha1.ControlPlaneRef",
			},
		},
	},
	{
		PackageVersion: "v1beta1",
		AdditionalImports: []string{
			`configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                "KongConsumerGroup",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "configurationv1alpha1.ControlPlaneRef",
			},
		},
	},
	{
		PackageVersion: "v1alpha1",
		Types: []templateDataT{
			{
				Type:                "KongKey",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndKeySetRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:                "KongKeySet",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
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
				Type:                "KongCACertificate",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:                "KongCertificate",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:                    "KongPluginBinding",
				KonnectStatusType:       "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType:     "ControlPlaneRef",
				ControlPlaneRefRequired: true,
			},
			{
				Type:                "KongService",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:                "KongRoute",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:                "KongUpstream",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:              "KongTarget",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs",
			},
			{
				Type:                "KongVault",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
			{
				Type:              "KongSNI",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndCertificateRefs",
			},
			{
				Type:                "KongDataPlaneClientCertificate",
				KonnectStatusType:   "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ControlPlaneRefType: "ControlPlaneRef",
			},
		},
	},
}

var supportedKonnectTypesStandalone = []supportedTypesT{
	{
		PackageVersion: "v1alpha1",
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

var supportedGatewayOperatorTypes = []supportedTypesT{
	{
		PackageVersion: "v1alpha1",
		Types: []templateDataT{
			{
				Type: "AIGateway",
			},
			{
				Type: "KongPluginInstallation",
			},
			{
				Type: "KonnectExtension",
			},
		},
	},
	{
		PackageVersion: "v1beta1",
		Types: []templateDataT{
			{
				Type: "DataPlane",
			},
			{
				Type: "ControlPlane",
			},
		},
	},
}

var supportedConfigurationPackageTypesWithList = supportedKonnectTypesControlPlaneConfig

var supportedKonnectPackageTypesWithList = supportedKonnectTypesStandalone

var supportedGatewayOperatorPackageTypesWithList = supportedGatewayOperatorTypes
