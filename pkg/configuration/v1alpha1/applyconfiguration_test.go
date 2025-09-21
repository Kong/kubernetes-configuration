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

package v1alpha1

import (
	"encoding/json"
	"testing"

	sdkkonnectcomp "github.com/Kong/sdk-konnect-go/models/components"

	commonv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/common/v1alpha1"
	configurationv1alpha1 "github.com/kong/kubernetes-configuration/v2/api/configuration/v1alpha1"
)

func TestKongRouteApplyConfiguration(t *testing.T) {
	tests := []struct {
		name string
		fn   func() *KongRouteApplyConfiguration
	}{
		{
			name: "basic kong route",
			fn: func() *KongRouteApplyConfiguration {
				return KongRoute("test-route", "test-namespace").
					WithSpec(KongRouteSpec().
						WithHosts("example.com").
						WithPaths("/test"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got == nil {
				t.Fatal("ApplyConfiguration should not be nil")
			}

			// Test that it can be marshaled to JSON
			data, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("failed to marshal ApplyConfiguration: %v", err)
			}

			// Test that it has the correct API version and kind
			var obj map[string]interface{}
			if err := json.Unmarshal(data, &obj); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			if got, want := obj["apiVersion"], "configuration.konghq.com/v1alpha1"; got != want {
				t.Errorf("apiVersion = %v, want %v", got, want)
			}
			if got, want := obj["kind"], "KongRoute"; got != want {
				t.Errorf("kind = %v, want %v", got, want)
			}
		})
	}
}

func TestKongServiceApplyConfiguration(t *testing.T) {
	tests := []struct {
		name string
		fn   func() *KongServiceApplyConfiguration
	}{
		{
			name: "basic kong service",
			fn: func() *KongServiceApplyConfiguration {
				return KongService("test-service", "test-namespace").
					WithSpec(KongServiceSpec().
						WithHost("backend.example.com").
						WithPort(80).
						WithProtocol(sdkkonnectcomp.ProtocolHTTP))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got == nil {
				t.Fatal("ApplyConfiguration should not be nil")
			}

			// Test that it can be marshaled to JSON
			data, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("failed to marshal ApplyConfiguration: %v", err)
			}

			// Test that it has the correct API version and kind
			var obj map[string]interface{}
			if err := json.Unmarshal(data, &obj); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			if got, want := obj["apiVersion"], "configuration.konghq.com/v1alpha1"; got != want {
				t.Errorf("apiVersion = %v, want %v", got, want)
			}
			if got, want := obj["kind"], "KongService"; got != want {
				t.Errorf("kind = %v, want %v", got, want)
			}
		})
	}
}

func TestKongUpstreamApplyConfiguration(t *testing.T) {
	tests := []struct {
		name string
		fn   func() *KongUpstreamApplyConfiguration
	}{
		{
			name: "basic kong upstream",
			fn: func() *KongUpstreamApplyConfiguration {
				return KongUpstream("test-upstream", "test-namespace").
					WithSpec(KongUpstreamSpec().
						WithName("test-upstream"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got == nil {
				t.Fatal("ApplyConfiguration should not be nil")
			}

			// Test that it can be marshaled to JSON
			data, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("failed to marshal ApplyConfiguration: %v", err)
			}

			// Test that it has the correct API version and kind
			var obj map[string]interface{}
			if err := json.Unmarshal(data, &obj); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			if got, want := obj["apiVersion"], "configuration.konghq.com/v1alpha1"; got != want {
				t.Errorf("apiVersion = %v, want %v", got, want)
			}
			if got, want := obj["kind"], "KongUpstream"; got != want {
				t.Errorf("kind = %v, want %v", got, want)
			}
		})
	}
}

func TestKongTargetApplyConfiguration(t *testing.T) {
	tests := []struct {
		name string
		fn   func() *KongTargetApplyConfiguration
	}{
		{
			name: "basic kong target",
			fn: func() *KongTargetApplyConfiguration {
				return KongTarget("test-target", "test-namespace").
					WithSpec(KongTargetSpec().
						WithUpstreamRef(commonv1alpha1.NameRef{Name: "test-upstream"}).
						WithTarget("backend.example.com:8080"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if got == nil {
				t.Fatal("ApplyConfiguration should not be nil")
			}

			// Test that it can be marshaled to JSON
			data, err := json.Marshal(got)
			if err != nil {
				t.Fatalf("failed to marshal ApplyConfiguration: %v", err)
			}

			// Test that it has the correct API version and kind
			var obj map[string]interface{}
			if err := json.Unmarshal(data, &obj); err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}

			if got, want := obj["apiVersion"], "configuration.konghq.com/v1alpha1"; got != want {
				t.Errorf("apiVersion = %v, want %v", got, want)
			}
			if got, want := obj["kind"], "KongTarget"; got != want {
				t.Errorf("kind = %v, want %v", got, want)
			}
		})
	}
}

// TestApplyConfigurationInterface verifies that our apply configurations implement the expected interface
func TestApplyConfigurationInterface(t *testing.T) {
	route := KongRoute("test-route", "test-namespace")
	service := KongService("test-service", "test-namespace")
	upstream := KongUpstream("test-upstream", "test-namespace")
	target := KongTarget("test-target", "test-namespace")

	// These should not panic - they implement IsApplyConfiguration()
	route.IsApplyConfiguration()
	service.IsApplyConfiguration()
	upstream.IsApplyConfiguration()
	target.IsApplyConfiguration()
}

// TestStructuredMergeDiff verifies that the apply configurations support structured merge-diff
func TestStructuredMergeDiff(t *testing.T) {
	// Test that we can create a complete configuration
	route := KongRoute("test-route", "test-namespace").
		WithAnnotations(map[string]string{
			"test": "annotation",
		}).
		WithLabels(map[string]string{
			"app": "test",
		}).
		WithSpec(KongRouteSpec().
			WithHosts("example.com").
			WithPaths("/api").
			WithMethods("GET", "POST").
			WithName("test-route").
			WithStripPath(true))

	// Marshal and verify structure
	data, err := json.Marshal(route)
	if err != nil {
		t.Fatalf("failed to marshal apply configuration: %v", err)
	}

	var obj configurationv1alpha1.KongRoute
	if err := json.Unmarshal(data, &obj); err != nil {
		t.Fatalf("failed to unmarshal into KongRoute: %v", err)
	}

	// Verify that the data was properly set
	if obj.APIVersion != "configuration.konghq.com/v1alpha1" {
		t.Errorf("APIVersion = %v, want configuration.konghq.com/v1alpha1", obj.APIVersion)
	}
	if obj.Kind != "KongRoute" {
		t.Errorf("Kind = %v, want KongRoute", obj.Kind)
	}
	if obj.Name != "test-route" {
		t.Errorf("Name = %v, want test-route", obj.Name)
	}
	if obj.Namespace != "test-namespace" {
		t.Errorf("Namespace = %v, want test-namespace", obj.Namespace)
	}
}