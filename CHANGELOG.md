# Table of Contents

<!---
Adding a new version? You'll need three changes:
* Add the ToC link, like "[v1.2.3](#v123)".
* Add the section header, like "## [v1.2.3]".
* Add the diff link, like "[v2.7.0]: https://github.com/kong/kubernetes-ingress-controller/compare/v1.2.2...v1.2.3".
--->
- [v1.2.0](#v120)
- [v1.1.0](#v110)
- [v1.0.6](#v106)
- [v1.0.5](#v105)
- [v1.0.4](#v104)
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.0](#v100)

## Unreleased

## [v1.3.0]

[v1.3.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.2.0...v1.3.0

### Fixes

- Proper validation of fields `dataPlaneOptions` and `controlPlaneOptions` for `GatewayConfiguration`
  [#359](https://github.com/Kong/kubernetes-configuration/pull/359)

### Changes

- Add `watchNamespaces` spec field to `ControlPlane` to allow watching only specific namespaces.
  [#358](https://github.com/Kong/kubernetes-configuration/pull/358)
- Support `NodePort` as ingress service type for `DataPlane`
  [#367](https://github.com/Kong/kubernetes-configuration/pull/367)
- Add support for network ref of type `namespacedRef` in `KonnectCloudGatewayDataPlaneGroupConfiguration`
  [#370](https://github.com/Kong/kubernetes-configuration/pull/370)

## [v1.2.0]

[v1.2.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.1.0...v1.2.0

### Added

- Migrate KGO CRDs to this repo.
  [#274](https://github.com/Kong/kubernetes-configuration/pull/274)
  [#282](https://github.com/Kong/kubernetes-configuration/pull/282)
- Added `KonnectCloudGatewayNetwork` CRD.
  [#268](https://github.com/Kong/kubernetes-configuration/pull/268)
- Added `GatewayConfiguration` extension point.
  [#300](https://github.com/Kong/kubernetes-configuration/pull/300)
- Added Endpoints to `KonnectGatewayControlPlane`'s status.
  [#299](https://github.com/Kong/kubernetes-configuration/pull/299)
- Added `konnect.konghq.com/KonnectExtension` CRD. `gateway-operator.konghq.com/KonnectExtension`
  has been deprecated.
  [#291](https://github.com/Kong/kubernetes-configuration/pull/291)
  [#315](https://github.com/Kong/kubernetes-configuration/pull/315)
  [#317](https://github.com/Kong/kubernetes-configuration/pull/317)
- Added `KonnectExtension` status conditions.
  [#301](https://github.com/Kong/kubernetes-configuration/pull/301)
- Added `KonnectCloudGatewayDataPlaneGroupConfiguration`.
  [#307](https://github.com/Kong/kubernetes-configuration/pull/307)
- Migrate KGO conditions to this repo.
  [#323](https://github.com/Kong/kubernetes-configuration/pull/323)
  [#337](https://github.com/Kong/kubernetes-configuration/pull/337)
- Disallowed `konnectID` as `ControlPlaneRef`'s `type` field value for Konnect entities that do not support it yet:
  - `KongCACertificate`
  - `KongCertificate`
  - `KongVault`
  - `KongDataPlaneClientCertificate`
  - `KongKey`
  - `KongKeySet`
  - `KongPluginBinding`
  - `KongRoute`
  - `KongService`
  - `KongUpstream`
  - `KongConsumerGroup`
  [#326](https://github.com/Kong/kubernetes-configuration/pull/326)
- Added `kong` category to Kong CRDs.
  [#336](https://github.com/Kong/kubernetes-configuration/pull/336)

### Changes

- Set `GenerateEmbeddedObjectMeta` to `true` when generating CRDs to align with
  existing KGO CRDs.
  [#298](https://github.com/Kong/kubernetes-configuration/pull/298)

## [v1.1.0]

[v1.1.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.6...v1.1.0

### Changes

- Fix `KongRoute`'s `spec.headers` field type.
  [#243](https://github.com/Kong/kubernetes-configuration/pull/243)
- Add a `scope` field to `KongPluginBindingSpec` to allow setting the scope of
  the plugin binding. The default value (`OnlyTargets`) is aligned with the previous
  default behavior - the plugin will only be applied to the targets specified in the
  `targets` field. A new alternative is `GlobalInControlPlane` that will make the
  plugin apply globally in a control plane.
  [#236](https://github.com/Kong/kubernetes-configuration/pull/236)
- Make `KongPluginBinding`'s `spec.controlPlaneRef` field required as that's expected
  by KGO.
  [#238](https://github.com/Kong/kubernetes-configuration/pull/238)

## [v1.0.6]

[v1.0.6]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.5...v1.0.6

### Changes

- Implemented `Stringer` interface for `configuraionv1alpha1.ControlPlaneRef` type.
  [#230](https://github.com/Kong/kubernetes-configuration/pull/230)

## [v1.0.5]

[v1.0.5]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.4...v1.0.5

> Release date: 2025-01-10

### Changes

- Exported CRD validation test suite types.
  [#220](https://github.com/Kong/kubernetes-configuration/pull/220)

## [v1.0.4]

[v1.0.4]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.3...v1.0.4

> Release date: 2025-01-09

### Changes

- Allowed `konnectID` as `ControlPlaneRef`'s `type` field value.
  [#214](https://github.com/Kong/kubernetes-configuration/pull/214)

## [v1.0.3]

[v1.0.3]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.2...v1.0.3

> Release date: 2025-01-07

### Changes

This release didn't include any user-facing changes.

## [v1.0.2]

[v1.0.2]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.0...v1.0.2

> Release date: 2025-01-07

### Changes

- Disallowed `konnectID` as `ControlPlaneRef`'s `type` field value. This is not
  supported by KGO yet.
  [#198](https://github.com/Kong/kubernetes-configuration/pull/198)
- Added version information to CRDs. All CRDs now will be annotated with the
  `kubernetes-configuration.konghq.com/version` annotation with the value of the
  version of the CRDs bundle.
  [#203](https://github.com/Kong/kubernetes-configuration/pull/203)

## [v1.0.0]

[v1.0.0]: https://github.com/kong/kubernetes-configuration/compare/ecf9b7bd62bfb92327a6ddd9aeaec9f73fc13a72...v1.0.0

> Release date: 2024-12-12

This is an initial stable release of the Kong Kubernetes configuration module.
It includes CRDs for the Kong Ingress Controller and the Kong Gateway Operator
in the `ingress-controller`, `ingress-controller-incubator` and `gateway-operator`
channels.

Go bindings for the CRDs are available in the [`api/`][api] directory and the
[`pkg/clientset/`][clientset] directory contains the clientset for interacting
with the CRDs.

[api]: ./api/
[clientset]: ./pkg/clientset/
