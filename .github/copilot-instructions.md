# Kong Kubernetes Configuration AI Agent Instructions

## Project Overview

This repository contains API definitions for Kong's Kubernetes configuration system, supporting multiple products through a **channel-based architecture**. The codebase generates CRDs, Go clientsets, and documentation for three main channels:

- `ingress-controller` - Kong Ingress Controller CRDs
- `ingress-controller-incubator` - Experimental KIC CRDs  
- `gateway-operator` - Kong Gateway Operator CRDs

## Critical Architecture Patterns

### Channel System & Markers
**IMPORTANT**: CRDs are organized by channels using the `+kong:channels` marker. When creating new CRDs:

```go
// Example from api/configuration/v1alpha1/kongroute_types.go
// +kong:channels=gateway-operator
type KongRoute struct {
    // ...
}
```

Channels determine which kustomize manifests include the CRD (`config/crd/<channel>/`). A single CRD can belong to multiple channels by separating them with semicolons: `+kong:channels=ingress-controller;gateway-operator`.

### API Reference Markers
For documentation generation, use specific markers:
- `+apireference:kgo:include` - Include in Gateway Operator docs
- `+apireference:kic:include` - Include in Ingress Controller docs  
- `+apireference:kgo:exclude` - Exclude specific fields from KGO docs
- `+apireference:kic:exclude` - Exclude specific fields from KIC docs

### Multi-Version API Structure
The repository follows Kubernetes API versioning patterns:
```
api/
├── configuration/     # Core Kong configuration objects
│   ├── v1/           # Stable APIs
│   ├── v1alpha1/     # Alpha APIs
│   └── v1beta1/      # Beta APIs
├── gateway-operator/  # Gateway Operator specific
│   ├── v1alpha1/
│   ├── v1beta1/
│   └── v2beta1/
├── konnect/          # Kong Konnect integration
│   ├── v1alpha1/
│   └── v1alpha2/
├── incubator/        # Experimental APIs
│   └── v1alpha1/
└── common/           # Shared types
    └── v1alpha1/
```

## Essential Development Workflows

### Code Generation (Always Required)
**Run after ANY API changes**:
```bash
make generate
```

This runs multiple generators:
- `generate.crds` - CRD manifests via controller-tools
- `generate.deepcopy` - Runtime deepcopy methods
- `generate.clientsets` - Go client libraries
- `generate.applyconfigurations` - Apply configurations for server-side apply
- `generate.docs` - API reference documentation
- `generate.apitypes-funcs` - Helper functions

### Verification
```bash
make verify.generators  # Ensures generated code is up-to-date
make verify.diff       # Checks for uncommitted changes
```

### Testing Strategy
```bash
make test.unit              # Unit tests in test/unit/
make test.crds-validation   # Live cluster CEL validation tests
make test.conversion        # API version conversion tests
make test.samples           # Validate example manifests
```

### Tool Management
Uses `mise` for tool version management. Tools are installed to `bin/installs/` and accessed via shims in `bin/shims/`. If tool commands fail, run:
```bash
make tools  # Install controller-gen, kustomize, client-gen
```

## CRD Development Patterns

### Required Annotations
All CRDs should include these kubebuilder markers:
```go
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:categories=kong
// +kubebuilder:storageversion  // For the storage version
// +kubebuilder:subresource:status
// +kubebuilder:ac:generate=true  // For apply configuration generation
```

### Common Validation Patterns
Immutable control plane references:
```go
// +kubebuilder:validation:XValidation:rule="!has(oldSelf.spec.controlPlaneRef) || has(self.spec.controlPlaneRef)", message="controlPlaneRef is required once set"
// +kubebuilder:validation:XValidation:rule="(!has(self.spec.controlPlaneRef)) ? true : (!has(self.status) || !self.status.conditions.exists(c, c.type == 'Programmed' && c.status == 'True')) ? true : oldSelf.spec.controlPlaneRef == self.spec.controlPlaneRef", message="spec.controlPlaneRef is immutable when an entity is already Programmed"
```

### Status Patterns
Most CRDs include Konnect integration status:
```go
type ResourceStatus struct {
    // Konnect metadata
    Konnect *konnectv1alpha2.KonnectEntityStatusWithControlPlaneRef `json:"konnect,omitempty"`
    // Standard conditions
    Conditions []metav1.Condition `json:"conditions,omitempty"`
}
```

## Integration Points

### Konnect Integration
Resources reference Konnect entities via:
- `KonnectEntityRef` for direct Konnect entity references
- `ControlPlaneRef` for associating with control planes
- Status includes Konnect sync state

### Gateway API Integration
Gateway Operator APIs heavily integrate with Kubernetes Gateway API:
- `ControlPlane` manages Gateway resources
- `GatewayConfiguration` provides gateway-specific options
- `DataPlane` handles Kong proxy deployments

## File Naming Conventions

- `*_types.go` - API type definitions
- `*_conversion.go` - Version conversion functions  
- `gvrs.go` - GroupVersionResource helpers
- `zz_generated*.go` - Auto-generated code (never edit manually)

## Critical Scripts

- `scripts/crds-generator/` - Custom CRD generation with channel support
- `scripts/apitypes-funcs/` - Generates helper functions for types
- `scripts/apidocs-gen/` - API documentation generation
- `scripts/verify-diff.sh` - Ensures no uncommitted generated code

## When Adding New CRDs

1. Add to appropriate API package with correct markers
2. Add channel annotation: `+kong:channels=<channel-name>`
3. Add apply configuration marker: `+kubebuilder:ac:generate=true`
4. Add to unit tests in `test/unit/`
5. Add CRD validation tests in `test/crdsvalidation/`
6. Update `scripts/apitypes-funcs/supportedtypes.go` if helper functions needed
7. Run `make generate` and commit all generated files
8. Add sample manifests to `config/samples/`

## Apply Configurations

The repository generates apply configurations for server-side apply with structured merge-diff. These are located in `pkg/configuration/v1alpha1/` and enable:

- **Partial updates**: Only specified fields are updated
- **Conflict resolution**: Kubernetes manages field ownership automatically
- **Efficient operations**: Only modified fields sent to API server

Example usage:
```go
import applyv1alpha1 "github.com/kong/kubernetes-configuration/v2/pkg/configuration/v1alpha1"

service := applyv1alpha1.KongService("my-service", "default").
    WithSpec(applyv1alpha1.KongServiceSpec().
        WithHost("api.example.com").
        WithPort(80))
```

Add `+kubebuilder:ac:generate=true` marker to enable generation for new CRDs.

## Testing Against Live Clusters

CRD validation tests in `test/crdsvalidation/` require a live Kubernetes cluster to test CEL validation rules. Ensure `kubectl` is configured before running `make test.crds-validation`.