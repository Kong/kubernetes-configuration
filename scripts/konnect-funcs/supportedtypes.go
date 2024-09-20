package main

var supportedTypes = []supportedTypesT{
	{
		Package: "v1",
		Types: []templateDataT{
			{
				Type:              "KongConsumer",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "c",
			},
		},
	},
	{
		Package: "v1beta1",
		Types: []templateDataT{
			{
				Type:              "KongConsumerGroup",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "g",
			},
		},
	},
	{
		Package: "v1alpha1",
		Types: []templateDataT{
			{
				Type:              "KongCredentialBasicAuth",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndConsumerRefs",
				ReceiverName:      "c",
			},
			{
				Type:              "KongCACertificate",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "c",
			},
			{
				Type:              "KongPluginBinding",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "b",
			},
			{
				Type:              "KongService",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "s",
			},
			{
				Type:              "KongRoute",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndServiceRefs",
				ReceiverName:      "r",
			},
			{
				Type:              "KongUpstream",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "u",
			},
			{
				Type:              "KongTarget",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneAndUpstreamRefs",
				ReceiverName:      "t",
			},
			{
				Type:              "KongVault",
				KonnectStatusType: "konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef",
				ReceiverName:      "v",
			},
		},
	},
}
