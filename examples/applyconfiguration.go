/*
Copyright 2021 Kong, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package examples demonstrates how to use Kong apply configurations
// for server-side apply with structured merge-diff.
package examples

import (
	"context"
	"fmt"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/common/v1alpha1"
	applyv1alpha1 "github.com/kong/kubernetes-configuration/v2/pkg/configuration/v1alpha1"
)

// ExampleServerSideApply demonstrates how to use Kong apply configurations
// with Kubernetes server-side apply for structured merge-diff.
func ExampleServerSideApply(ctx context.Context, client kubernetes.Interface, namespace string) error {
	// Create a KongService apply configuration
	service := applyv1alpha1.KongService("example-service", namespace).
		WithLabels(map[string]string{
			"app":     "example",
			"version": "v1",
		}).
		WithAnnotations(map[string]string{
			"konghq.com/description": "Example service managed via server-side apply",
		}).
		WithSpec(applyv1alpha1.KongServiceSpec().
			WithHost("api.example.com").
			WithPort(443).
			WithProtocol(sdkkonnectcomp.ProtocolHTTPS).
			WithPath("/v1"))

	fmt.Printf("KongService apply configuration: %+v\n", service)

	// Create a KongUpstream apply configuration
	upstream := applyv1alpha1.KongUpstream("example-upstream", namespace).
		WithLabels(map[string]string{
			"app":     "example",
			"version": "v1",
		}).
		WithSpec(applyv1alpha1.KongUpstreamSpec().
			WithName("example-upstream").
			WithAlgorithm("round-robin").
			WithHashOn("none"))

	fmt.Printf("KongUpstream apply configuration: %+v\n", upstream)

	// Create a KongTarget apply configuration
	target := applyv1alpha1.KongTarget("example-target", namespace).
		WithLabels(map[string]string{
			"app":     "example",
			"version": "v1",
		}).
		WithSpec(applyv1alpha1.KongTargetSpec().
			WithUpstreamRef(commonv1alpha1.NameRef{Name: "example-upstream"}).
			WithTarget("backend.example.com:443").
			WithWeight(100))

	fmt.Printf("KongTarget apply configuration: %+v\n", target)

	// Create a KongRoute apply configuration
	route := applyv1alpha1.KongRoute("example-route", namespace).
		WithLabels(map[string]string{
			"app":     "example",
			"version": "v1",
		}).
		WithSpec(applyv1alpha1.KongRouteSpec().
			WithServiceRef(&applyv1alpha1.ServiceRef{Name: "example-service"}).
			WithHosts("api.example.com").
			WithPaths("/api/v1").
			WithMethods("GET", "POST").
			WithName("example-route").
			WithStripPath(false))

	fmt.Printf("KongRoute apply configuration: %+v\n", route)

	return nil
}

// ExamplePartialUpdate demonstrates how apply configurations enable
// partial updates with structured merge-diff.
func ExamplePartialUpdate() *applyv1alpha1.KongServiceApplyConfiguration {
	// You can create partial configurations that only update specific fields
	// This will only update the host and port, leaving other fields unchanged
	return applyv1alpha1.KongService("example-service", "default").
		WithSpec(applyv1alpha1.KongServiceSpec().
			WithHost("new-api.example.com").
			WithPort(8080))
}

// ExampleStatusUpdate demonstrates how to update only the status of a resource.
func ExampleStatusUpdate() *applyv1alpha1.KongServiceApplyConfiguration {
	// You can update only status fields
	return applyv1alpha1.KongService("example-service", "default").
		WithStatus(applyv1alpha1.KongServiceStatus().
			WithConditions(metav1.Condition{
				Type:               "Programmed",
				Status:             metav1.ConditionTrue,
				Reason:             "Synced",
				Message:            "Resource has been synced successfully",
				LastTransitionTime: metav1.Now(),
			}))
}

// ExampleComplexRoute demonstrates a more complex route configuration
// with multiple fields set using the fluent API.
func ExampleComplexRoute() *applyv1alpha1.KongRouteApplyConfiguration {
	return applyv1alpha1.KongRoute("complex-route", "production").
		WithLabels(map[string]string{
			"app":        "api-gateway",
			"component":  "routing",
			"managed-by": "kong-operator",
		}).
		WithAnnotations(map[string]string{
			"konghq.com/description": "Complex route with multiple features",
			"konghq.com/priority":    "high",
		}).
		WithSpec(applyv1alpha1.KongRouteSpec().
			WithServiceRef(&applyv1alpha1.ServiceRef{Name: "backend-service"}).
			WithHosts("api.prod.example.com", "api.example.com").
			WithPaths("/api/v2", "/api/v3").
			WithMethods("GET", "POST", "PUT", "DELETE").
			WithName("complex-api-route").
			WithStripPath(true).
			WithPreserveHost(true).
			WithRequestBuffering(true).
			WithResponseBuffering(true).
			WithTags("production", "api", "v2"))
}