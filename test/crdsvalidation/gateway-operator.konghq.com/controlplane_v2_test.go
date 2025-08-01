package crdsvalidation_test

import (
	"testing"
	"time"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/common/v1alpha1"
	operatorv2alpha1 "github.com/kong/kubernetes-configuration/v2/api/gateway-operator/v2alpha1"
	"github.com/kong/kubernetes-configuration/v2/test/crdsvalidation/common"
)

func TestControlPlaneV2(t *testing.T) {
	validDataPlaneTarget := operatorv2alpha1.ControlPlaneDataPlaneTarget{
		Type: operatorv2alpha1.ControlPlaneDataPlaneTargetRefType,
		Ref: &operatorv2alpha1.ControlPlaneDataPlaneTargetRef{
			Name: "dataplane-1",
		},
	}

	t.Run("extensions", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "no extensions",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
			},
			{
				Name: "konnectExtension set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
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
				Name: "konnectExtension and DataPlaneMetricsExtension set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
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
			{
				Name: "invalid extension",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
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
				ExpectedErrorMessage: lo.ToPtr("spec.dataplane.type: Unsupported value: \"\""),
			},
			{
				Name: "when dataplane.type is set to name, name must be specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
							Type: operatorv2alpha1.ControlPlaneDataPlaneTargetRefType,
						},
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Ref has to be provided when type is set to ref"),
			},
			{
				Name: "specifying dataplane ref name when type is ref passes",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: operatorv2alpha1.ControlPlaneDataPlaneTarget{
							Type: operatorv2alpha1.ControlPlaneDataPlaneTargetRefType,
							Ref: &operatorv2alpha1.ControlPlaneDataPlaneTargetRef{
								Name: "dataplane-1",
							},
						},
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
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
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
			},
			{
				Name: "feature gate set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:  "KongCustomEntity",
									State: operatorv2alpha1.FeatureGateStateEnabled,
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
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:  "KongCustomEntity",
									State: operatorv2alpha1.FeatureGateStateDisabled,
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
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name:  "KongCustomEntity",
									State: operatorv2alpha1.FeatureGateStateEnabled,
								},
							},
						},
					},
				},
				Update: func(cp *operatorv2alpha1.ControlPlane) {
					cp.Spec.ControlPlaneOptions = operatorv2alpha1.ControlPlaneOptions{}
				},
			},
			{
				Name: "cannot provide a feature gate with enabled unset",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							FeatureGates: []operatorv2alpha1.ControlPlaneFeatureGate{
								{
									Name: "KongCustomEntity",
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.featureGates[0].state: Unsupported value: \"\": supported values: \"enabled\", \"disabled\""),
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
						DataPlane:           validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
					},
				},
			},
			{
				Name: "controller overrides specified",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:  "GatewayAPI",
									State: operatorv2alpha1.ControllerStateEnabled,
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
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:  "GatewayAPI",
									State: operatorv2alpha1.ControllerStateDisabled,
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
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name:  "GatewayAPI",
									State: operatorv2alpha1.ControllerStateEnabled,
								},
							},
						},
					},
				},
				Update: func(cp *operatorv2alpha1.ControlPlane) {
					cp.Spec.ControlPlaneOptions = operatorv2alpha1.ControlPlaneOptions{}
				},
			},
			{
				Name: "cannot provide a controller with enabled unset",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Controllers: []operatorv2alpha1.ControlPlaneController{
								{
									Name: "GatewayAPI",
								},
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.controllers[0].state: Unsupported value: \"\": supported values: \"enabled\", \"disabled\""),
			},
		}.Run(t)
	})

	t.Run("translation", func(t *testing.T) {
		common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
			{
				Name: "combinedServicesFromDifferentHTTPRoutes set to enabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								CombinedServicesFromDifferentHTTPRoutes: lo.ToPtr(operatorv2alpha1.ControlPlaneCombinedServicesFromDifferentHTTPRoutesStateEnabled),
							},
						},
					},
				},
			},
			{
				Name: "combinedServicesFromDifferentHTTPRoutes set to disabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								CombinedServicesFromDifferentHTTPRoutes: lo.ToPtr(operatorv2alpha1.ControlPlaneCombinedServicesFromDifferentHTTPRoutesStateEnabled),
							},
						},
					},
				},
			},
			{
				Name: "combinedServicesFromDifferentHTTPRoutes set to disallowed value",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								CombinedServicesFromDifferentHTTPRoutes: lo.ToPtr(operatorv2alpha1.ControlPlaneCombinedServicesFromDifferentHTTPRoutesState("invalid")),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.translation.combinedServicesFromDifferentHTTPRoutes: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
			},
			{
				Name: "drainSupport set to enabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								DrainSupport: lo.ToPtr(operatorv2alpha1.ControlPlaneDrainSupportStateEnabled),
							},
						},
					},
				},
			},
			{
				Name: "drainSupport set to disabled",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								DrainSupport: lo.ToPtr(operatorv2alpha1.ControlPlaneDrainSupportStateDisabled),
							},
						},
					},
				},
			},
			{
				Name: "drainSupport set to disallowed value",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								DrainSupport: lo.ToPtr(operatorv2alpha1.ControlPlaneDrainSupportState("invalid")),
							},
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.translation.drainSupport: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
			},
			{
				Name: "both combinedServicesFromDifferentHTTPRoutes and drainSupport set",
				TestObject: &operatorv2alpha1.ControlPlane{
					ObjectMeta: common.CommonObjectMeta,
					Spec: operatorv2alpha1.ControlPlaneSpec{
						DataPlane: validDataPlaneTarget,
						ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
							Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
								CombinedServicesFromDifferentHTTPRoutes: lo.ToPtr(operatorv2alpha1.ControlPlaneCombinedServicesFromDifferentHTTPRoutesStateEnabled),
								DrainSupport:                            lo.ToPtr(operatorv2alpha1.ControlPlaneDrainSupportStateDisabled),
							},
						},
					},
				},
			},
		}.Run(t)

		t.Run("fallbackConfiguration", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "fallbackConfiguration.useLastValidConfig set to enabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
									FallbackConfiguration: &operatorv2alpha1.ControlPlaneFallbackConfiguration{
										UseLastValidConfig: lo.ToPtr(operatorv2alpha1.ControlPlaneFallbackConfigurationStateEnabled),
									},
								},
							},
						},
					},
				},
				{
					Name: "fallbackConfiguration.useLastValidConfig set to disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
									FallbackConfiguration: &operatorv2alpha1.ControlPlaneFallbackConfiguration{
										UseLastValidConfig: lo.ToPtr(operatorv2alpha1.ControlPlaneFallbackConfigurationStateDisabled),
									},
								},
							},
						},
					},
				},
				{
					Name: "fallbackConfiguration.useLastValidConfig set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Translation: &operatorv2alpha1.ControlPlaneTranslationOptions{
									FallbackConfiguration: &operatorv2alpha1.ControlPlaneFallbackConfiguration{
										UseLastValidConfig: lo.ToPtr(operatorv2alpha1.ControlPlaneFallbackConfigurationState("invalid")),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.translation.fallbackConfiguration.useLastValidConfig: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
				},
			}.Run(t)
		})

		t.Run("configDump", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "configDump.state and configDump.dumpsensitive set to enabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpStateEnabled,
									DumpSensitive: operatorv2alpha1.ConfigDumpStateEnabled,
								},
							},
						},
					},
				},
				{
					Name: "configDump.state and configDump.dumpSensitive set to disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpStateDisabled,
									DumpSensitive: operatorv2alpha1.ConfigDumpStateDisabled,
								},
							},
						},
					},
				},
				{
					Name: "configDump.state set to enabled and configDump.dumpSensitive set to disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpStateEnabled,
									DumpSensitive: operatorv2alpha1.ConfigDumpStateDisabled,
								},
							},
						},
					},
				},
				{
					Name: "configDump.state set to disabled and configDump.dumpSensitive set to enabled is invalid",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpStateDisabled,
									DumpSensitive: operatorv2alpha1.ConfigDumpStateEnabled,
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("Cannot enable dumpSensitive when state is disabled"),
				},
				{
					Name: "configDump.state set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpState("invalid"),
									DumpSensitive: operatorv2alpha1.ConfigDumpStateEnabled,
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr(`spec.configDump.state: Unsupported value: "invalid": supported values: "enabled", "disabled"`),
				},
				{
					Name: "configDump.dumpSensitive is set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ConfigDump: &operatorv2alpha1.ControlPlaneConfigDump{
									State:         operatorv2alpha1.ConfigDumpStateEnabled,
									DumpSensitive: operatorv2alpha1.ConfigDumpState("invalid"),
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr(`spec.configDump.dumpSensitive: Unsupported value: "invalid": supported values: "enabled", "disabled"`),
				},
			}.Run(t)
		})

		t.Run("objectFilters", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "objectFilters.secrets and objectFilters.configMaps are set",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ObjectFilters: &operatorv2alpha1.ControlPlaneObjectFilters{
									Secrets: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{"konghq.com/secret": "true"},
									},
									ConfigMaps: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{"konghq.com/configmap": "true"},
									},
								},
							},
						},
					},
				},
				{
					Name: "maximum items in matchLabels is 8",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ObjectFilters: &operatorv2alpha1.ControlPlaneObjectFilters{
									Secrets: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{
											"konghq.com/secret": "true",
											"label1":            "value1",
											"label2":            "value2",
											"label3":            "value3",
											"label4":            "value4",
											"label5":            "value5",
											"label6":            "value6",
											"label7":            "value7",
											"label8":            "value8",
										},
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.objectFilters.secrets.matchLabels: Too many: 9: must have at most 8 items"),
				},
				{
					Name: "key of objectFilters.*.matchLabels must have minimum length 1",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ObjectFilters: &operatorv2alpha1.ControlPlaneObjectFilters{
									Secrets: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{"konghq.com/secret": "true"},
									},
									ConfigMaps: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{"": "aaa"},
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("Minimum length of key in matchLabels is 1"),
				},
				{
					Name: "value of objectFilters.*.matchLabels must have maximum length 63",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								ObjectFilters: &operatorv2alpha1.ControlPlaneObjectFilters{
									Secrets: &operatorv2alpha1.ControlPlaneFilterForObjectType{
										MatchLabels: map[string]string{"konghq.com/secret": "this-is-a-very-very-long-label-which-is-longer-than-63-characters"},
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("Maximum length of value in matchLabels is 63"),
				},
			}.Run(t)
		})
	})

	t.Run("konnect", func(t *testing.T) {
		t.Run("basic configuration", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "no konnect configuration",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane:           validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{},
						},
					},
				},
				{
					Name: "konnect configuration with all options set",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									ConsumersSync: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectConsumersSyncStateEnabled),
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:              lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										InitialPollingPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
										PollingPeriod:        lo.ToPtr(metav1.Duration{Duration: 300 * time.Second}),
										StorageState:       lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
									},
									NodeRefreshPeriod:  lo.ToPtr(metav1.Duration{Duration: 60 * time.Second}),
									ConfigUploadPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
								},
							},
						},
					},
				},
			}.Run(t)
		})

		t.Run("consumersSync", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "consumersSync set to enabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									ConsumersSync: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectConsumersSyncStateEnabled),
								},
							},
						},
					},
				},
				{
					Name: "consumersSync set to disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									ConsumersSync: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectConsumersSyncStateDisabled),
								},
							},
						},
					},
				},
				{
					Name: "consumersSync set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									ConsumersSync: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectConsumersSyncState("invalid")),
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.konnect.consumersSync: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
				},
			}.Run(t)
		})

		t.Run("licensing", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "licensing set to enabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("initialPollingPeriod is required when licensing is enabled"),
				},
				{
					Name: "licensing set to disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:        lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateDisabled),
										StorageState: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateDisabled),
									},
								},
							},
						},
					},
				},
				{
					Name: "licensing with polling periods and storage",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:              lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										InitialPollingPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
										PollingPeriod:        lo.ToPtr(metav1.Duration{Duration: 300 * time.Second}),
										StorageState:       lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
									},
								},
							},
						},
					},
				},
				{
					Name: "licensing with storage disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:              lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										InitialPollingPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
										PollingPeriod:        lo.ToPtr(metav1.Duration{Duration: 300 * time.Second}),
										StorageState:       lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateDisabled),
									},
								},
							},
						},
					},
				},
				{
					Name: "licensing storage set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:        lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										StorageState: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingState("invalid")),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.konnect.licensing.storageState: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
				},
				{
					Name: "licensing enabled without initialPollingPeriod",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:       lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										PollingPeriod: lo.ToPtr(metav1.Duration{Duration: 300 * time.Second}),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("initialPollingPeriod is required when licensing is enabled"),
				},
				{
					Name: "licensing enabled without pollingPeriod",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:              lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
										InitialPollingPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("pollingPeriod is required when licensing is enabled"),
				},
				{
					Name: "storageState set when licensing is disabled",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State:        lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateDisabled),
										StorageState: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingStateEnabled),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("storageState can only be set to enabled when licensing is enabled"),
				},
				{
					Name: "licensing set to disallowed value",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									Licensing: &operatorv2alpha1.ControlPlaneKonnectLicensing{
										State: lo.ToPtr(operatorv2alpha1.ControlPlaneKonnectLicensingState("invalid")),
									},
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.konnect.licensing.state: Unsupported value: \"invalid\": supported values: \"enabled\", \"disabled\""),
				},
			}.Run(t)
		})

		t.Run("periods", func(t *testing.T) {
			common.TestCasesGroup[*operatorv2alpha1.ControlPlane]{
				{
					Name: "nodeRefreshPeriod set",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									NodeRefreshPeriod: lo.ToPtr(metav1.Duration{Duration: 60 * time.Second}),
								},
							},
						},
					},
				},
				{
					Name: "configUploadPeriod set",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									ConfigUploadPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
								},
							},
						},
					},
				},
				{
					Name: "both periods set",
					TestObject: &operatorv2alpha1.ControlPlane{
						ObjectMeta: common.CommonObjectMeta,
						Spec: operatorv2alpha1.ControlPlaneSpec{
							DataPlane: validDataPlaneTarget,
							ControlPlaneOptions: operatorv2alpha1.ControlPlaneOptions{
								Konnect: &operatorv2alpha1.ControlPlaneKonnectOptions{
									NodeRefreshPeriod:  lo.ToPtr(metav1.Duration{Duration: 60 * time.Second}),
									ConfigUploadPeriod: lo.ToPtr(metav1.Duration{Duration: 30 * time.Second}),
								},
							},
						},
					},
				},
			}.Run(t)
		})
	})
}
