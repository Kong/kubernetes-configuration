package testcases

import (
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	"github.com/samber/lo"
)

var kongSNIAPISpec = testCasesGroup{
	Name: "kongSNIAPISpec",
	TestCases: []testCase{
		{
			Name: "weight must between 0 and 65535",
			KongSNI: configurationv1alpha1.KongSNI{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongSNISpec{
					CertificateRef: configurationv1alpha1.TargetRef{
						Name: "cert1",
					},
				},
			},
			ExpectedErrorMessage: lo.ToPtr("spec.name in body should be at least 1 chars long"),
		},
	},
}
