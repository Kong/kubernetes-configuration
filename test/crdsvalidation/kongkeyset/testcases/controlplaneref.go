package testcases

import (
	"github.com/samber/lo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

var cpRef = testCasesGroup{
	Name: "controlPlaneRef",
	TestCases: []testCase{
		{
			Name: "konnectNamespacedRef reference is valid",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
							Name: "test-konnect-control-plane",
						},
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
		},
		{
			Name: "not providing konnectNamespacedRef when type is konnectNamespacedRef yields an error",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
			ExpectedErrorMessage: lo.ToPtr("when type is konnectNamespacedRef, konnectNamespacedRef must be set"),
		},
		{
			Name: "not providing konnectID when type is konnectID yields an error",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectID,
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
			ExpectedErrorMessage: lo.ToPtr("when type is konnectID, konnectID must be set"),
		},
		{
			Name: "konnectNamespacedRef reference name cannot be changed when an entity is Programmed",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
							Name: "test-konnect-control-plane",
						},
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
			KongKeySetStatus: &configurationv1alpha1.KongKeySetStatus{
				Conditions: []metav1.Condition{
					{
						Type:               "Programmed",
						Status:             metav1.ConditionTrue,
						Reason:             "Programmed",
						LastTransitionTime: metav1.Now(),
					},
				},
			},
			Update: func(ks *configurationv1alpha1.KongKeySet) {
				ks.Spec.ControlPlaneRef.KonnectNamespacedRef.Name = "new-konnect-control-plane"
			},
			ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
		},
		{
			Name: "konnectNamespacedRef reference type cannot be changed when an entity is Programmed",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
							Name: "test-konnect-control-plane",
						},
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
			KongKeySetStatus: &configurationv1alpha1.KongKeySetStatus{
				Conditions: []metav1.Condition{
					{
						Type:               "Programmed",
						Status:             metav1.ConditionTrue,
						Reason:             "Programmed",
						LastTransitionTime: metav1.Now(),
					},
				},
			},
			Update: func(ks *configurationv1alpha1.KongKeySet) {
				ks.Spec.ControlPlaneRef.Type = configurationv1alpha1.ControlPlaneRefKonnectID
			},
			ExpectedUpdateErrorMessage: lo.ToPtr("spec.controlPlaneRef is immutable when an entity is already Programmed"),
		},
		{
			Name: "konnectNamespaced reference cannot set namespace as it's not supported yet",
			KongKeySet: configurationv1alpha1.KongKeySet{
				ObjectMeta: commonObjectMeta,
				Spec: configurationv1alpha1.KongKeySetSpec{
					ControlPlaneRef: &configurationv1alpha1.ControlPlaneRef{
						Type: configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef,
						KonnectNamespacedRef: &configurationv1alpha1.KonnectNamespacedRef{
							Name:      "test-konnect-control-plane",
							Namespace: "default",
						},
					},
					KongKeySetAPISpec: configurationv1alpha1.KongKeySetAPISpec{
						Name: "name",
					},
				},
			},
			ExpectedErrorMessage: lo.ToPtr("spec.controlPlaneRef cannot specify namespace for namespaced resource - it's not supported yet"),
		},
	},
}
