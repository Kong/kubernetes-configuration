package testcases

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
)

// testCase is a test case related to KongRoute validation.
type testCase struct {
	Name                       string
	KongRoute                  configurationv1alpha1.KongRoute
	KongRouteStatus            *configurationv1alpha1.KongRouteStatus
	Update                     func(*configurationv1alpha1.KongRoute)
	ExpectedErrorMessage       *string
	ExpectedUpdateErrorMessage *string
}

// testCasesGroup is a group of test cases related to KongRoute validation.
// The grouping is done by a common name.
type testCasesGroup struct {
	Name      string
	TestCases []testCase
}

// TestCases is a collection of all test cases groups related to KongRoute validation.
var TestCases = []testCasesGroup{}

func init() {
	TestCases = append(TestCases,
		cpRef,
		serviceRef,
		protocols,
	)
}

var commonObjectMeta = metav1.ObjectMeta{
	GenerateName: "test-kongroute-",
	Namespace:    "default",
}
