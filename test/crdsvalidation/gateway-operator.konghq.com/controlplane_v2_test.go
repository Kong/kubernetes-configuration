package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	operatorv2alpha1 "github.com/kong/kubernetes-configuration/api/gateway-operator/v2alpha1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation/common"
)

func TestControlPlaneV2(t *testing.T) {
	validDataPlaneTarget := operatorv2alpha1.ControlPlaneDataPlaneTarget{
		Type: operatorv2alpha1.ControlPlaneDataPlaneTargetName,
		Name: "dataplane",
	}

	validControlPlaneOptions := operatorv2alpha1.ControlPlaneOptions{
		DataPlane: validDataPlaneTarget,
	}

	t.Run("extensions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "no extensions",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: validControlPlaneOptions,
					},
				},
			},
			{
				Name: "konnectExtension set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
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
			},
			{
				Name: "konnectExtension and DataPlaneMetricsExtension set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "konnect.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
								{
									Group: "gateway-operator.konghq.com",
									Kind:  "DataPlaneMetricsExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-metrics-extension",
									},
								},
							},
						},
					},
				},
			},
			{
				Name: "invalid extension",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							Extensions: []commonv1alpha1.ExtensionRef{
								{
									Group: "invalid.konghq.com",
									Kind:  "KonnectExtension",
									NamespacedRef: commonv1alpha1.NamespacedRef{
										Name: "my-konnect-extension",
									},
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Extension not allowed for ControlPlane"),
			},
		}.Run(t)
	})

	t.Run("feature gates", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "no feature gates",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: validControlPlaneOptions,
					},
				},
			},
			{
				Name: "feature gate set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:    "KongCustomEntity",
									Enabled: lo.ToPtr(true),
								},
							},
						},
					},
				},
			},
			{
				Name: "feature gate disabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:    "KongCustomEntity",
									Enabled: lo.ToPtr(false),
								},
							},
						},
					},
				},
			},
			{
				Name: "feature gate set and then removed",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:    "KongCustomEntity",
									Enabled: lo.ToPtr(true),
								},
							},
						},
					},
				},
				Update: func(cp *operatorv2alpha1.ControlPlane) {
					cp.Spec.ControlPlaneOptions = validControlPlaneOptions
				},
			},
		}.Run(t)
	})

	t.Run("controllers", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "no controller overrides specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: validControlPlaneOptions,
					},
				},
			},
			{
				Name: "controller overrides specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:    "GatewayAPI",
									Enabled: lo.ToPtr(true),
								},
							},
						},
					},
				},
			},
			{
				Name: "controller overrides specified - disabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:    "GatewayAPI",
									Enabled: lo.ToPtr(false),
								},
							},
						},
					},
				},
			},
			{
				Name: "controller overrides specified and then removed",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: validDataPlaneTarget,
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:    "GatewayAPI",
									Enabled: lo.ToPtr(true),
								},
							},
						},
					},
				},
				Update: func(cp *operatorv2alpha1.ControlPlane) {
					cp.Spec.ControlPlaneOptions = validControlPlaneOptions
				},
			},
		}.Run(t)
	})
}
