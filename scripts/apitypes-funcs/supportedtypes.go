package main

import "github.com/samber/lo"

var supportedKonnectV1Alpha1TypesWithControlPlaneRef = []supportedTypesT{
	{
		PackageVersion: "v1alpha1",
		AdditionalImports: []string{
			`commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                       "KonnectCloudGatewayDataPlaneGroupConfiguration",
				KonnectStatusType:          "KonnectEntityStatusWithControlPlaneRef",
				KonnectStatusEmbedded:      true,
				GetKonnectStatusReturnType: "*KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
				ControlPlaneRefRequired:    true,
			},
			{
				Type:                     "KonnectExtension",
				ControlPlaneRefType:      "commonv1alpha1.ControlPlaneRef",
				ControlPlaneRefRequired:  true,
				ControlPlaneRefFieldPath: "Spec.Konnect.ControlPlane.ControlPlaneRef",
			},
		},
	},
}

var supportedKonnectTypesControlPlaneConfig = []supportedTypesT{
	{
		PackageVersion: "v1",
		AdditionalImports: []string{
			`commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"`,
			`konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                       "KongConsumer",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type: "KongPlugin",
			},
		},
	},
	{
		PackageVersion: "v1beta1",
		AdditionalImports: []string{
			`commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"`,
			`konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                       "KongConsumerGroup",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
		},
	},
	{
		PackageVersion: "v1alpha1",
		AdditionalImports: []string{
			`commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"`,
			`konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"`,
		},
		Types: []templateDataT{
			{
				Type:                       "KongKey",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndKeySetRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongKeySet",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongCredentialBasicAuth",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongCredentialAPIKey",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongCredentialJWT",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongCredentialACL",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongCredentialHMAC",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongCACertificate",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongCertificate",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongPluginBinding",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
				ControlPlaneRefRequired:    true,
			},
			{
				Type:                       "KongService",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongRoute",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongUpstream",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongTarget",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongVault",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
			{
				Type:                       "KongSNI",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndCertificateRefs",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
			},
			{
				Type:                       "KongDataPlaneClientCertificate",
				KonnectStatusType:          "*konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				GetKonnectStatusReturnType: "*konnectv1alpha1.KonnectEntityStatus",
				ControlPlaneRefType:        "commonv1alpha1.ControlPlaneRef",
			},
		},
	},
}

var supportedKonnectTypesStandalone = []supportedTypesT{
	{
		PackageVersion: "v1alpha1",
		Types: []templateDataT{
			{
				Type:                       "KonnectGatewayControlPlane",
				KonnectStatusType:          "*KonnectEntityStatus",
				KonnectStatusEmbedded:      true,
				GetKonnectStatusReturnType: "KonnectEntityStatus",
			},
			{
				Type: "KonnectAPIAuthConfiguration",
			},
			{
				Type:                       "KonnectCloudGatewayNetwork",
				KonnectStatusType:          "*KonnectEntityStatus",
				KonnectStatusEmbedded:      true,
				GetKonnectStatusReturnType: "KonnectEntityStatus",
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

var supportedKonnectPackageTypesWithList = func() []supportedTypesT {
	// Make sure that each template is generated once per package version.
	base := append(
		supportedKonnectTypesStandalone,
		supportedKonnectV1Alpha1TypesWithControlPlaneRef...,
	)

	m := make(map[string]supportedTypesT)
	for _, t := range base {
		v, ok := m[t.PackageVersion]
		if !ok {
			m[t.PackageVersion] = t
			continue
		}
		v.Types = append(m[t.PackageVersion].Types, t.Types...)
		m[t.PackageVersion] = v
	}

	return lo.Values(m)
}()

var supportedGatewayOperatorPackageTypesWithList = supportedGatewayOperatorTypes
