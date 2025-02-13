package crdsvalidation_test

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation"
)

func TestKongService(t *testing.T) {
	obj := &configurationv1alpha1.KongService{
		TypeMeta: metav1.TypeMeta{
			Kind:       "KongService",
			APIVersion: configurationv1alpha1.GroupVersion.String(),
		},
		ObjectMeta: commonObjectMeta,
	}

	t.Run("cp ref", func(t *testing.T) {
		NewCRDValidationTestCasesGroupCPRefChange(t, obj, NotSupportedByKIC, ControlPlaneRefRequired).Run(t)
	})

	t.Run("cp ref, type=kic", func(t *testing.T) {
		NewCRDValidationTestCasesGroupCPRefChangeKICUnsupportedTypes(t, obj, EmptyControlPlaneRefNotAllowed).Run(t)
	})

	t.Run("konnect adopt", func(t *testing.T) {
		crdsvalidation.TestCasesGroup[*configurationv1alpha1.KongService]{
			{
				Name: "konnect adopt can be changed before getting programmed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KonnectOptions: &konnectv1alpha1.KonnectEntityOptions{
							Adopt: &konnectv1alpha1.KonnectAdoptOptions{
								ID: "abcddcba-0000-1111-9999-0123456789ab",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
							Host: "example.com",
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongService) {
					ks.Spec.KonnectOptions = &konnectv1alpha1.KonnectEntityOptions{
						Adopt: &konnectv1alpha1.KonnectAdoptOptions{
							ID: "abcddcba-0000-1111-9999-0123456789ac",
						},
					}
				},
			},
			{
				Name: "konnect adopt cannot be changed after programmed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KonnectOptions: &konnectv1alpha1.KonnectEntityOptions{
							Adopt: &konnectv1alpha1.KonnectAdoptOptions{
								ID: "abcddcba-0000-1111-9999-0123456789ab",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
							Host: "example.com",
						},
					},
					Status: configurationv1alpha1.KongServiceStatus{
						Konnect: &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{
							KonnectEntityStatus: konnectv1alpha1.KonnectEntityStatus{
								ID: "abcddcba-0000-1111-9999-0123456789ac",
							},
							ControlPlaneID: "cp-1",
						},
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Valid",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongService) {
					ks.Spec.KonnectOptions = &konnectv1alpha1.KonnectEntityOptions{
						Adopt: &konnectv1alpha1.KonnectAdoptOptions{
							ID: "abcddcba-0000-1111-9999-fd9876543211",
						},
					}
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("spec.konnect.adopt is immutable when an entity is already Programmed"),
			},
			{
				Name: "Cannot set konnect adopt after programmed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
							Host: "example.com",
						},
					},
					Status: configurationv1alpha1.KongServiceStatus{
						Konnect: &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{
							KonnectEntityStatus: konnectv1alpha1.KonnectEntityStatus{
								ID: "abcddcba-0000-1111-9999-fdecba987654",
							},
							ControlPlaneID: "cp-1",
						},
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Valid",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongService) {
					ks.Spec.KonnectOptions = &konnectv1alpha1.KonnectEntityOptions{
						Adopt: &konnectv1alpha1.KonnectAdoptOptions{
							ID: "abcddcba-0000-1111-9999-1234567890ab",
						},
					}
				},
				ExpectedUpdateErrorMessage: lo.ToPtr("Cannot add spec.konnect.adopt when an entity is already programmed"),
			},
			{
				Name: "konnect adopt can be removed after programmed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KonnectOptions: &konnectv1alpha1.KonnectEntityOptions{
							Adopt: &konnectv1alpha1.KonnectAdoptOptions{
								ID: "abcddcba-0000-1111-9999-0123456789ab",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
							Host: "example.com",
						},
					},
					Status: configurationv1alpha1.KongServiceStatus{
						Konnect: &konnectv1alpha1.KonnectEntityStatusWithControlPlaneRef{
							KonnectEntityStatus: konnectv1alpha1.KonnectEntityStatus{
								ID: "abcddcba-0000-1111-9999-fdecba987654",
							},
							ControlPlaneID: "cp-1",
						},
						Conditions: []metav1.Condition{
							{
								Type:               "Programmed",
								Status:             metav1.ConditionTrue,
								Reason:             "Valid",
								LastTransitionTime: metav1.Now(),
							},
						},
					},
				},
				Update: func(ks *configurationv1alpha1.KongService) {
					ks.Spec.KonnectOptions.Adopt = nil
				},
			},
		}.Run(t)
	})

	t.Run("tags validation", func(t *testing.T) {
		crdsvalidation.TestCasesGroup[*configurationv1alpha1.KongService]{
			{
				Name: "up to 20 tags are allowed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
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
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
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
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: commonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
								Name: "test-konnect-control-plane",
							},
						},
						KongServiceAPISpec: configurationv1alpha1.KongServiceAPISpec{
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
}
