package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	operatorv1beta1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1beta1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation/common"
)

func TestGatewayConfiguration(t *testing.T) {
	t.Run("extensions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv1beta1.GatewayConfiguration]{
			{
				Name: "no extensions",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec:       operatorv1beta1.GatewayConfigurationSpec{},
				},
			},
			{
				Name: "valid konnectExtension at the gatewayConfiguration level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						Extensions: []commonv1alpha1.ExtensionRef{
							{
								Group: "konnect.konghq.com",
								Kind:  "KonnectExtension",
								NamespacedRef: commonv1alpha1.NamespacedRef{
									Name: "my-konnect-extension",
								},
							},
						},
					},
				},
			},
			{
				Name: "invalid konnectExtension",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						Extensions: []commonv1alpha1.ExtensionRef{
							{
								Group: "wrong.konghq.com",
								Kind:  "wrongExtension",
								NamespacedRef: commonv1alpha1.NamespacedRef{
									Name: "my-konnect-extension",
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Extension not allowed for GatewayConfiguration"),
			},
			{
				Name: "konnectExtension at the DataPlane level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						DataPlaneOptions: &operatorv1beta1.GatewayConfigDataPlaneOptions{
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "konnect.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("KonnectExtension must be set at the Gateway level"),
			},
			{
				Name: "konnectExtension at the ControlPlane level",
				TestObject: &operatorv1beta1.GatewayConfiguration{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv1beta1.GatewayConfigurationSpec{
						ControlPlaneOptions: &operatorv1beta1.ControlPlaneOptions{
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "konnect.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("KonnectExtension must be set at the Gateway level"),
			},
		}.Run(t)
	})
}
