package configuration_test

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/api/common/v1alpha1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	"github.com/kong/kubernetes-configuration/test/crdsvalidation/common"
)

func TestKongService(t *testing.T) {
	obj := &configurationv1alpha1.KongService{
		TypeMeta: metav1.TypeMeta{
			Kind:       "KongService",
			APIVersion: configurationv1alpha1.GroupVersion.String(),
		},
		ObjectMeta: common.CommonObjectMeta,
	}

	t.Run("cp ref", func(t *testing.T) {
		common.NewCRDValidationTestCasesGroupCPRefChange(t, obj, common.NotSupportedByKIC, common.ControlPlaneRefRequired).Run(t)
	})

	t.Run("cp ref, type=kic", func(t *testing.T) {
		common.NewCRDValidationTestCasesGroupCPRefChangeKICUnsupportedTypes(t, obj, common.EmptyControlPlaneRefNotAllowed).Run(t)
	})

	t.Run("tags validation", func(t *testing.T) {
		common.TestCasesGroup[*configurationv1alpha1.KongService]{
			{
				Name: "up to 20 tags are allowed",
				TestObject: &configurationv1alpha1.KongService{
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &commonv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &commonv1alpha1.KonnectNamespacedRef{
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
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &commonv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &commonv1alpha1.KonnectNamespacedRef{
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
					ObjectMeta: common.CommonObjectMeta,
					Spec: configurationv1alpha1.KongServiceSpec{
						ControlPlaneRef: &commonv1alpha1.ControlPlaneRef{
							Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
							KonnectNamespacedRef: &commonv1alpha1.KonnectNamespacedRef{
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
