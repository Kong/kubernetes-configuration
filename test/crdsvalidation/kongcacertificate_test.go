package crdsvalidation_test

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

func TestKongCACertificate(t *testing.T) {
	t.Run("required fields validation", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongCACertificate]{
			{
				Name: "cert field is required",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.cert: Required value"),
			},
		}.Run(t)
	})

	t.Run("cp ref validation", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongCACertificate]{
			{
				Name: "konnectNamespacedRef reference is valid",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
							Cert: "cert",
						},
					},
				},
			},
			{
				Name: "not providing konnectNamespacedRef when type is konnectNamespacedRef yields an error",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
							Cert: "cert",
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectNamespacedRef, konnectNamespacedRef must be set"),
			},
			{
				Name: "not providing konnectID when type is konnectID yields an error",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectID,
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
							Cert: "cert",
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectID, konnectID must be set"),
			},
			{
				Name: "konnectNamespacedRef reference name cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongCACertificateStatus{
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
				Update: func(ks *configurationv1alpha1.KongCACertificate) {
					ks.Spec.ControlPlaneRef.KonnectNamespacedRef.Name = "new-konnect-control-plane"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
			{
				Name: "konnectNamespacedRef reference type cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongCACertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongCACertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongCACertificateStatus{
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
				Update: func(ks *configurationv1alpha1.KongCACertificate) {
					ks.Spec.ControlPlaneRef.Type = configurationv1alpha1.ControlPlaneRefKonnectID
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
		}.Run(t)

		t.Run("tags validation", func(t *testing.T) {
			CRDValidationTestCasesGroup[*configurationv1alpha1.KongCACertificate]{
				{
					Name: "up to 20 tags are allowed",
					TestObject: &configurationv1alpha1.KongCACertificate{
						ObjectMeta: commonObjectMeta,
						Spec: configurationv1alpha1.KongCACertificateSpec{
							ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
								Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
								KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
									Name: "test-konnect-control-plane",
								},
							},
							KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
								Cert: "cert",
								Tags: func() []string {
									var tags []string
									for i := range 20 {
										tags = append(tags, fmt.Sprintf("tag-%d", i))
									}
									return tags
								}(),
							},
						},
					},
				},
				{
					Name: "more than 20 tags are not allowed",
					TestObject: &configurationv1alpha1.KongCACertificate{
						ObjectMeta: commonObjectMeta,
						Spec: configurationv1alpha1.KongCACertificateSpec{
							ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
								Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
								KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
									Name: "test-konnect-control-plane",
								},
							},
							KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
								Cert: "cert",
								Tags: func() []string {
									var tags []string
									for i := range 21 {
										tags = append(tags, fmt.Sprintf("tag-%d", i))
									}
									return tags
								}(),
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("spec.tags: Too many: 21: must have at most 20 items"),
				},
				{
					Name: "tags entries must not be longer than 128 characters",
					TestObject: &configurationv1alpha1.KongCACertificate{
						ObjectMeta: commonObjectMeta,
						Spec: configurationv1alpha1.KongCACertificateSpec{
							ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
								Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
								KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
									Name: "test-konnect-control-plane",
								},
							},
							KongCACertificateAPISpec: configurationv1alpha1.KongCACertificateAPISpec{
								Cert: "cert",
								Tags: []string{
									lo.RandomString(129, lo.AlphanumericCharset),
								},
							},
						},
					},
					ExpectedErrorMessage: lo.ToPtr("tags entries must not be longer than 128 characters"),
				},
			}.Run(t)
		})
	})
}
