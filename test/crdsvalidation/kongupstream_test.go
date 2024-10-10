package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

func TestKongUpstream(t *testing.T) {
	t.Run("cp ref", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongUpstream]{
			{
				Name: "konnectNamespacedRef reference is valid",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
				},
			},
			{
				Name: "not providing konnectNamespacedRef when type is konnectNamespacedRef yields an error",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectNamespacedRef, konnectNamespacedRef must be set"),
			},
			{
				Name: "not providing konnectID when type is konnectID yields an error",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectID,
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectID, konnectID must be set"),
			},
			{
				Name: "providing namespace in konnectNamespacedRef yields an error",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name:      "test-konnect-control-plane",
								Namespace: "another-namespace",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.controlPlaneRef cannot specify namespace for namespaced resource"),
			},
			{
				Name: "konnectNamespacedRef reference name cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
					Status: configurationv1alpha1.KongUpstreamStatus{
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Programmed",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongUpstream) {
					ks.Spec.ControlPlaneRef.KonnectNamespacedRef.Name = "new-konnect-control-plane"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
			{
				Name: "konnectNamespacedRef reference type cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{},
					},
					Status: configurationv1alpha1.KongUpstreamStatus{
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Programmed",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongUpstream) {
					ks.Spec.ControlPlaneRef.Type = configurationv1alpha1.ControlPlaneRefKonnectID
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
		}.Run(t)
	})

	t.Run("required fields", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongUpstream]{
			{
				Name: "hash_fallback_header is required when hash_fallback is set to 'header'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback:       lo.ToPtr(sdkkonnectcomp.HashFallbackHeader),
							HashFallbackHeader: lo.ToPtr("X-Hash-Fallback"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_fallback_header is not provided when hash_fallback is set to 'header'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback: lo.ToPtr(sdkkonnectcomp.HashFallbackHeader),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_fallback_header is required when `hash_fallback` is set to `header`"),
			},
			{
				Name: "hash_fallback_query_arg is required when hash_fallback is set to 'query_arg'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback:         lo.ToPtr(sdkkonnectcomp.HashFallbackQueryArg),
							HashFallbackQueryArg: lo.ToPtr("arg"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_fallback_query_arg is not provided when hash_fallback is set to 'query_arg'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback: lo.ToPtr(sdkkonnectcomp.HashFallbackQueryArg),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_fallback_query_arg is required when `hash_fallback` is set to `query_arg`"),
			},
			{
				Name: "hash_fallback_uri_capture is required when hash_fallback is set to 'uri_capture'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback:           lo.ToPtr(sdkkonnectcomp.HashFallbackURICapture),
							HashFallbackURICapture: lo.ToPtr("arg"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_fallback_uri_capture is not provided when hash_fallback is set to 'uri_capture'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback: lo.ToPtr(sdkkonnectcomp.HashFallbackURICapture),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_fallback_uri_capture is required when `hash_fallback` is set to `uri_capture`"),
			},
			{
				Name: "hash_on_cookie and hash_on_cookie_path are required when hash_on is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn:           lo.ToPtr(sdkkonnectcomp.HashOnCookie),
							HashOnCookie:     lo.ToPtr("cookie"),
							HashOnCookiePath: lo.ToPtr("X-Hash-On-Cookie-Path"),
						},
					},
				},
			},
			{
				Name: "hash_on_cookie and hash_on_cookie_path are required when hash_fallback is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback:     lo.ToPtr(sdkkonnectcomp.HashFallbackCookie),
							HashOnCookie:     lo.ToPtr("cookie"),
							HashOnCookiePath: lo.ToPtr("X-Hash-On-Cookie-Path"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_on_cookie is not provided when hash_on is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn:           lo.ToPtr(sdkkonnectcomp.HashOnCookie),
							HashOnCookiePath: lo.ToPtr("X-Hash-On-Cookie-Path"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie is required when hash_on is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_cookie is not provided when hash_fallback is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback:     lo.ToPtr(sdkkonnectcomp.HashFallbackCookie),
							HashOnCookiePath: lo.ToPtr("X-Hash-On-Cookie-Path"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie is required when hash_fallback is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_cookie_path is not provided when hash_on is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn:       lo.ToPtr(sdkkonnectcomp.HashOnCookie),
							HashOnCookie: lo.ToPtr("cookie"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie_path is required when hash_on is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_cookie_path is not provided when hash_fallback is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback: lo.ToPtr(sdkkonnectcomp.HashFallbackCookie),
							HashOnCookie: lo.ToPtr("cookie"),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie_path is required when hash_fallback is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_cookie_path nor hash_on_cookie are provided when hash_fallback is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashFallback: lo.ToPtr(sdkkonnectcomp.HashFallbackCookie),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie_path is required when hash_fallback is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_cookie_path nor hash_on_cookie are provided when hash_on is set to 'cookie'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn: lo.ToPtr(sdkkonnectcomp.HashOnCookie),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("hash_on_cookie_path is required when hash_on is set to `cookie`."),
			},
			{
				Name: "validation fails when hash_on_header is not provided when hash_on is set to 'header'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn: lo.ToPtr(sdkkonnectcomp.HashOnHeader),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_on_header is required when hash_on is set to `header`"),
			},
			{
				Name: "hash_on_query_arg is required when hash_on is set to 'query_arg'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn:         lo.ToPtr(sdkkonnectcomp.HashOnQueryArg),
							HashOnQueryArg: lo.ToPtr("arg"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_on_query_arg is not provided when hash_on is set to 'query_arg'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn: lo.ToPtr(sdkkonnectcomp.HashOnQueryArg),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_on_query_arg is required when `hash_on` is set to `query_arg`"),
			},
			{
				Name: "hash_on_uri_capture is required when hash_on is set to 'uri_capture'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn:           lo.ToPtr(sdkkonnectcomp.HashOnURICapture),
							HashOnURICapture: lo.ToPtr("arg"),
						},
					},
				},
			},
			{
				Name: "validation fails when hash_on_uri_capture is not provided when hash_on is set to 'uri_capture'",
				TestObject: &configurationv1alpha1.KongUpstream{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongUpstreamSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongUpstreamAPISpec: configurationv1alpha1.KongUpstreamAPISpec{
							HashOn: lo.ToPtr(sdkkonnectcomp.HashOnURICapture),
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("Invalid value: \"object\": hash_on_uri_capture is required when `hash_on` is set to `uri_capture`"),
			},
		}.Run(t)
	})
}
