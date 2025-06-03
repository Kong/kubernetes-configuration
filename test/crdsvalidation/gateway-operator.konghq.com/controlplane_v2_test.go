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
		Name: lo.ToPtr("dataplane"),
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

	t.Run("dataplane", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "missing dataplane causes an error",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.dataplane.type: Required value"),
			},
			{
				Name: "when dataplane.type is set to name, name must be specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetName,
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Name has to be provided when type is set to name"),
			},
			{
				Name: "when dataplane.type is set to url, url must be specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetURL,
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("URL has to be provided when type is set to url"),
			},
			{
				Name: "specifying dataplane name when type is name passes",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetName,
								Name: lo.ToPtr("dataplane"),
							},
						},
					},
				},
			},
			{
				Name: "specifying dataplane url when type is url passes: https",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetURL,
								URL:  lo.ToPtr("https://dataplane.example.com:8444/admin"),
							},
						},
					},
				},
			},
			{
				Name: "specifying dataplane url when type is url passes: http, no port",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetURL,
								URL:  lo.ToPtr("http://dataplane.example.com/admin"),
							},
						},
					},
				},
			},
			{
				Name: "dataplane url must be a valid URL, otherwise it fails",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
								Type: operatorv2alpha1.ControlPlaneDataPlaneTargetURL,
								URL:  lo.ToPtr("not-a-valid-url"),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("invalid: spec.dataplane.url: Invalid value: \"not-a-valid-url\": spec.dataplane.url in body should match '^https?://[a-zA-Z0-9.-]+(:[0-9]+)?(/.*)?$'"),
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
