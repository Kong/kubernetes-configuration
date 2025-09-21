# Kong Apply Configurations

This directory contains generated apply configurations for Kong custom resources. Apply configurations enable server-side apply with structured merge-diff, providing efficient and conflict-free updates to Kubernetes resources.

## Overview

Apply configurations are generated Go types that mirror the structure of Kong custom resources but provide fluent "With" methods for building resource specifications. They are designed to work with Kubernetes server-side apply, which offers several benefits:

- **Structured Merge-Diff**: Only the fields you specify are updated, leaving other fields unchanged
- **Conflict Resolution**: Kubernetes manages field ownership and resolves conflicts automatically  
- **Efficient Updates**: Only modified fields are sent to the API server
- **Partial Updates**: You can update specific parts of a resource without specifying the entire structure

## Supported Resources

The following Kong resources have apply configurations available:

- **KongRoute**: Route configuration for Kong Gateway
- **KongService**: Service configuration for Kong Gateway  
- **KongUpstream**: Upstream configuration for load balancing
- **KongTarget**: Target endpoints for upstreams

## Usage

### Basic Example

```go
import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    
    applyv1alpha1 "github.com/kong/kubernetes-configuration/v2/pkg/configuration/v1alpha1"
)

// Create a service configuration
service := applyv1alpha1.KongService("my-service", "default").
    WithSpec(applyv1alpha1.KongServiceSpec().
        WithHost("api.example.com").
        WithPort(80).
        WithProtocol("http"))

// The apply configuration can be used with server-side apply
// This would typically be done through a Kong-specific client that supports apply
```

### Partial Updates

Apply configurations excel at partial updates:

```go
// Only update the host and port, leave everything else unchanged
update := applyv1alpha1.KongService("my-service", "default").
    WithSpec(applyv1alpha1.KongServiceSpec().
        WithHost("new-api.example.com").
        WithPort(443))
```

### Status Updates

You can also update just the status:

```go
statusUpdate := applyv1alpha1.KongService("my-service", "default").
    WithStatus(applyv1alpha1.KongServiceStatus().
        WithConditions(metav1.Condition{
            Type:               "Programmed",
            Status:             metav1.ConditionTrue,
            Reason:             "Synced",
            Message:            "Service configured successfully",
            LastTransitionTime: metav1.Now(),
        }))
```

### Complex Configurations

For more complex scenarios:

```go
route := applyv1alpha1.KongRoute("api-route", "production").
    WithLabels(map[string]string{
        "app":        "api-gateway",
        "component":  "routing",
    }).
    WithAnnotations(map[string]string{
        "konghq.com/description": "Main API route",
    }).
    WithSpec(applyv1alpha1.KongRouteSpec().
        WithServiceRef(&applyv1alpha1.ServiceRef{Name: "backend-service"}).
        WithHosts("api.example.com").
        WithPaths("/api/v1").
        WithMethods("GET", "POST", "PUT", "DELETE").
        WithStripPath(true))
```

## Server-Side Apply Integration

Apply configurations are designed to work with server-side apply. When used with a Kubernetes client that supports apply operations, they provide:

1. **Field Management**: Kubernetes tracks which client owns which fields
2. **Three-Way Merge**: Combines your desired state, current state, and last applied configuration
3. **Conflict Detection**: Automatic detection and resolution of conflicting updates
4. **Atomic Updates**: All or nothing updates to ensure consistency

## Generation

Apply configurations are automatically generated from the Kong API types using `applyconfiguration-gen`. To regenerate them:

```bash
make generate.applyconfigurations
```

This is typically included in the main generation target:

```bash
make generate
```

## Structured Merge-Diff Parser

The `internal` package contains a structured merge-diff parser that enables proper field tracking and conflict resolution. This parser is automatically generated and should not be modified manually.

## Testing

Apply configurations include comprehensive tests to ensure they:

- Can be marshaled to/from JSON correctly
- Have the proper API version and kind set
- Support all expected "With" methods
- Work with structured merge-diff operations

Run tests with:

```bash
go test ./pkg/configuration/v1alpha1/
```