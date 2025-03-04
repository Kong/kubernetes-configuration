package crdsvalidation_test

import (
	"testing"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v1beta1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation"
	"github.com/samber/lo"
)

func TestGatewayConfiguration(t *testing.T) {
	t.Run("extensions", func(t *testing.T) {
		crdsvalidation.TestCasesGroup[*konnectv1alpha1.GatewayConfiguration]{
			{
				Name: "no extensions",
				TestObject: &konnectv1alpha1.GatewayConfiguration{
					ObjectMeta: commonObjectMeta,
					Spec:       konnectv1alpha1.GatewayConfigurationSpec{},
				},
			},
			{
				Name: "valid konnectExtension",
				TestObject: &konnectv1alpha1.GatewayConfiguration{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.GatewayConfigurationSpec{
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
				TestObject: &konnectv1alpha1.GatewayConfiguration{
					ObjectMeta: commonObjectMeta,
					Spec: konnectv1alpha1.GatewayConfigurationSpec{
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
		}.Run(t)
	})
}
