package testcases

import (
	"github.com/samber/lo"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

var certificateRef = testCasesGroup{
	Name: "certificateRef",
	TestCases: []testCase{
		{
			Name: "certificate ref is immutable",
			KongSNI: configurationv1alpha1.KongSNI{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongSNISpec{
					CertificateRef: configurationv1alpha1.TargetRef{
						Name: "cert1",
					},
					KongSNIAPISpec: configurationv1alpha1.KongSNIAPISpec{
						Name: "example.com",
					},
				},
			},
			Update: func(sni *configurationv1alpha1.KongSNI) {
				sni.Spec.CertificateRef = configurationv1alpha1.TargetRef{
					Name: "cert-2",
				}
			},
			ExpectedUpdateErrorMessage: lo.ToPtr("spec.certificateRef is immutable"),
		},
	},
}
