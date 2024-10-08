package crdsvalidation_test

import (
	"testing"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

func TestKongDataPlaneClientCertificate(t *testing.T) {
	t.Run("cp ref", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongDataPlaneClientCertificate]{
			{
				Name: "konnectNamespacedRef reference is valid",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
				},
			},
			{
				Name: "not providing konnectNamespacedRef when type is konnectNamespacedRef yields an error",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectNamespacedRef, konnectNamespacedRef must be set"),
			},
			{
				Name: "not providing konnectID when type is konnectID yields an error",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectID,
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("when type is konnectID, konnectID must be set"),
			},
			{
				Name: "konnectNamespacedRef reference name cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongDataPlaneClientCertificateStatus{
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
				Update: func(ks *configurationv1alpha1.KongDataPlaneClientCertificate) {
					ks.Spec.ControlPlaneRef.KonnectNamespacedRef.Name = "new-konnect-control-plane"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
			{
				Name: "konnectNamespacedRef reference type cannot be changed when an entity is Programmed",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongDataPlaneClientCertificateStatus{
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
				Update: func(ks *configurationv1alpha1.KongDataPlaneClientCertificate) {
					ks.Spec.ControlPlaneRef.Type = configurationv1alpha1.ControlPlaneRefKonnectID
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
			},
			{
				Name: "konnectNamespaced reference cannot set namespace as it's not supported yet",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name:      "test-konnect-control-plane",
								Namespace: "default",
							},
						},
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
				},
				ExpectedErrorMessage: lo.ToPtr("spec.controlPlaneRef cannot specify namespace for namespaced resource - it's not supported yet"),
			},
		}.Run(t)
	})

	t.Run("spec", func(t *testing.T) {
		CRDValidationTestCasesGroup[*configurationv1alpha1.KongDataPlaneClientCertificate]{
			{
				Name: "valid KongDataPlaneClientCertificate",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
				},
			},
			{
				Name: "cert is required",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
				},
				ExpectedErrorMessage: lo.ToPtr("spec.cert in body should be at least 1 chars long"),
			},
			{
				Name: "cert can be altered before programmed",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongDataPlaneClientCertificateStatus{
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionFalse,
								Reason:             "Pending",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(k *configurationv1alpha1.KongDataPlaneClientCertificate) {
					k.Spec.Cert = "cert2"
				},
			},
			{
				Name: "cert becomes immutable after programmed",
				TestObject: &configurationv1alpha1.KongDataPlaneClientCertificate{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongDataPlaneClientCertificateSpec{
						KongDataPlaneClientCertificateAPISpec: configurationv1alpha1.KongDataPlaneClientCertificateAPISpec{
							Cert: "cert",
						},
					},
					Status: configurationv1alpha1.KongDataPlaneClientCertificateStatus{
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
				Update: func(k *configurationv1alpha1.KongDataPlaneClientCertificate) {
					k.Spec.Cert = "cert2"
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.cert is immutable when an entity is already Programmed"),
			},
		}.Run(t)
	})

}
