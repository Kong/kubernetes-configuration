# Table of Contents

<!---
Adding a new version? You'll need three changes:
* Add the ToC link, like "[v1.2.3](#v123)".
* Add the section header, like "## [v1.2.3]".
* Add the diff link, like "[v2.7.0]: https://github.com/kong/kubernetes-ingress-controller/compare/v1.2.2...v1.2.3".
--->
- [v2.0.0-alpha.3](#v200-alpha3)
- [v2.0.0-alpha.0](#v200-alpha0)
- [v1.5.2](#v152)
- [v1.5.1](#v151)
- [v1.5.0](#v150)
- [v1.4.2](#v142)
- [v1.4.1](#v141)
- [v1.4.0](#v140)
- [v1.3.2](#v132)
- [v1.3.1](#v131)
- [v1.3.0](#v130)
- [v1.2.0](#v120)
- [v1.1.0](#v110)
- [v1.0.6](#v106)
- [v1.0.5](#v105)
- [v1.0.4](#v104)
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.0](#v100)

## Unreleased

### Changes

- Implemented conversion functions between `KonnectGatewayControlPlane` `v1alpha1` and
  `v1alpha2`, enabling seamless migration between API versions.
  To prevent import cycles, `v1alpha1` now imports required types from `v1alpha2`.
  [#550](https://github.com/Kong/kubernetes-configuration/pull/550)

### Added

- Added optional `ingressClassName` pointer field to the following CRDs to allow association with a specific `IngressClass` and ingress controller:
  - `KongConsumer` (v1, in KongConsumerSpec)
  - `KongConsumerGroup` (v1beta1, in KongConsumerGroupSpec)
  - `KongClusterPlugin` (v1)
  - `KongVault` (v1alpha1, in KongVaultSpec)
  - `KongCustomEntity` (v1alpha1, in KongCustomEntitySpec)

## [v2.0.0-alpha.3]

[v2.0.0-alpha.3]: https://github.com/Kong/kubernetes-configuration/compare/v2.0.0-alpha.0...v2.0.0-alpha.3

### Breaking Changes

- `ControlPlane` `v2alpha1` has been replaced by `ControlPlane` `v2beta1`
  `GatewayConfiguration` `v2alpha1` has been replaced by `GatewayConfiguration` `v2beta1`
  [#548](https://github.com/Kong/kubernetes-configuration/pull/548)
- `KonnectGatewayControlPlane v1alpha2` has been introduced.
  The `CreateControlPlaneRequest` fields (`name, description, clusterType, authType, cloudGateway, proxyUrls, labels`) have been moved from the top level of `spec` into a new structured field: `spec.createControlPlaneRequest`. The old flat field layout is no longer supported in `v1alpha2`.
  *Action required*:
  - Update any manifests or code that reference these fields to use the new nested structure.
  [#502](https://github.com/Kong/kubernetes-configuration/pull/502)
- Removed `all` CRD categories from all CRDs.
  Added `konnect` category to all Konnect CRDs.
  [#541](https://github.com/Kong/kubernetes-configuration/pull/541)
- Removed `KongIngress`, `TCPIngress`, `UDPIngress`.
  [#542](https://github.com/Kong/kubernetes-configuration/pull/542)

### Changes

- Added test cases to cover CEL rules.
  [#538](https://github.com/Kong/kubernetes-configuration/pull/538)
- The `ControlPlane` provisioned conditions reasons have been renamed.
  The reason for the condition status true is now `Provisioned`, while the reason
  related to the provisioning non completed yet has been renamed to `ProvisioningInProgress`.
  [#546](https://github.com/Kong/kubernetes-configuration/pull/546)

### Added

- Added `spec.konnect` for `ControlPlane` `v2alpha1` allows configuration
  of Konnect-related options:
  - `consumersSync.enabled`: Configure consumer synchronization with Konnect (enabled/disabled)
  - `licensing.state`: Enable or disable Konnect licensing
  - `licensing.initialPollingPeriod`: Set initial polling period for license checks
  - `licensing.pollingPeriod`: Set regular polling period for license checks
  - `licensing.storageState`: Configure whether to store licenses fetched from Konnect to Secrets locally (enabled/disabled)
  - `nodeRefreshPeriod`: Configure refresh period for node information in Konnect
  - `configUploadPeriod`: Configure period for uploading configuration to Konnect
  Also added CEL validation rules for `ControlPlane` `v2alpha1` Konnect licensing configuration:
  - `initialPollingPeriod` and `pollingPeriod` can only be set to `enabled` when `licensing.state` is set to `enabled`
  - `storageState` can only be set to `enabled` when `licensing.state` is set to `enabled`
  [#535](https://github.com/Kong/kubernetes-configuration/pull/535)

## [v2.0.0-alpha.0]

[v2.0.0-alpha.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.5.2...v2.0.0-alpha.0

### Breaking Changes

- `KonnectExtension` `v1alpha2` has been introduced as the API does not allow anymore to
  reference Konnect Gateway ControlPlanes via plain KonnectID.
  Use Mirror `KonnectGatewayControlPlane`s instead.
  [#449](https://github.com/Kong/kubernetes-configuration/pull/449)
  [#452](https://github.com/Kong/kubernetes-configuration/pull/452)
- Extensions can only be set at the `GatewayConfiguration` spec level now or
  in the `ControlPlane` and `DataPlane` spec fields.
  `GatewayConfiguration` does not have `ControlPlane` and `DataPlane` extension
  API fields anymore.
  [#470](https://github.com/Kong/kubernetes-configuration/pull/470)

### Cleanups

- Unsuported `gateway-operator.konghq.com` `KonnectExtension` marked as deprecated
  ( it's already been since v1.5.0 but now the CRD is marked as such ).
  The only API version affected is `v1alpha1` which is now marked as unserved.
  Users should migrate to `v1alpha2` version of `konnect.konghq.com/KonnectExtension` API instead.
  [#450](https://github.com/Kong/kubernetes-configuration/pull/450)
- `DataPlaneMetricsExtension` is not marked as EE only anymore.
  [#456](https://github.com/Kong/kubernetes-configuration/pull/456)

### Added

- Added `GatewayConfiguration` `v2alpha1` API version.
  This is now the storage version for `GatewayConfiguration`.
  [#462](https://github.com/Kong/kubernetes-configuration/pull/462)
- Added `ControlPlane` `v2alpha1` API version.
  This is now the storage version for `ControlPlane`.
  [#441](https://github.com/Kong/kubernetes-configuration/pull/441)
  [#454](https://github.com/Kong/kubernetes-configuration/pull/454)
- `ControlPlane` (and by extension `GatewayConfiguration`) now allows configuration
  of the following fields:
  - under `spec.gatewayDiscovery`:
    - `readinessCheckTimeout` and `readinessCheckInterval`
    [#503](https://github.com/Kong/kubernetes-configuration/pull/503)
  - under `spec.cache`:
    - `initSyncDuration`
    [#512](https://github.com/Kong/kubernetes-configuration/pull/512)
  - under `spec.dataplaneSync`:
    - `interval` and `timeout`
    [#513](https://github.com/Kong/kubernetes-configuration/pull/513)
  - under `spec.configDump`:
    - `enabled` and `dumpSensitive`
    [#518](https://github.com/Kong/kubernetes-configuration/pull/518)
  - under `spec.objectFilters`:
    - `secrets` and `configMaps` to constrain watched `Secret`s and `ConfigMap`s
    [#534](https://github.com/Kong/kubernetes-configuration/pull/534)

## [v1.5.2]

[v1.5.2]: https://github.com/Kong/kubernetes-configuration/compare/v1.5.1...v1.5.2

### Fixes

- Fix `KongUpstreamPolicy` CRD CEL validation rules on `hash_on`.
  [#498](https://github.com/Kong/kubernetes-configuration/pull/498)

## [v1.5.1]

[v1.5.1]: https://github.com/Kong/kubernetes-configuration/compare/v1.5.0...v1.5.1

### Cleanups

- Update the CRDs metadata with up to date version.
  [#495](https://github.com/Kong/kubernetes-configuration/pull/495)

## [v1.5.0]

[v1.5.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.4.2...v1.5.0

### Cleanups

- Marked `KongIngress`, `TCPIngress`, and `UDPIngress` as deprecated.
  You can use Gateway API resources instead.
  See the migration guides for [KongIngress](https://developer.konghq.com/kubernetes-ingress-controller/migrate/kongingress/)
  and [TCPIngress/UDPIngress](https://developer.konghq.com/kubernetes-ingress-controller/migrate/ingress-to-gateway/) for more details.
  [#464](https://github.com/Kong/kubernetes-configuration/pull/464)

### Added

- Added sticky sessions support for `KongUpstreamPolicy`.
  It corresponds to the `sticky session` feature of Kong Gateway.
  In KIC, we have added a `drain support` feature.
  When both are enabled, sticky session will continue to be bound to Kong Pod
  when it's marked as terminating.
  When sticky session is enabled but drain support is not, then sticky session
  will stop routing traffic to Kong Pods when they are marked as terminating.
  [#463](https://github.com/Kong/kubernetes-configuration/pull/463)

## [v1.4.2]

[v1.4.2]: https://github.com/Kong/kubernetes-configuration/compare/v1.4.1...v1.4.2

### Fixes

- Fix `DataPlane` CEL validation rule during blue green rollout
  [#439](https://github.com/Kong/kubernetes-configuration/pull/439)

## [v1.4.1]

[v1.4.1]: https://github.com/Kong/kubernetes-configuration/compare/v1.4.0...v1.4.1

### Changes

- Clarified the comment in `DataPlane` CRD on the allowed values for `Service` types.
  [#429](https://github.com/Kong/kubernetes-configuration/pull/429)

## [v1.4.0]

[v1.4.0]: https://github.com/Kong/kubernetes-configuration/compare/v1.3.2...v1.4.0

### Added

- Added `KonnectCloudGatewayTransitGateway` to support Konnect transit gateways.
  [#375](https://github.com/Kong/kubernetes-configuration/pull/375)
- Allow setting `DataPlane`'s `NodePort` port number
  [#401](https://github.com/Kong/kubernetes-configuration/pull/401)
- Added `scale` subresource to `DataPlane` CRD.
  [#402](https://github.com/Kong/kubernetes-configuration/pull/402)
- Added support for setting `PodDisruptionBudget` in `GatewayConfiguration`'s `DataPlane` options.
  [#405](https://github.com/Kong/kubernetes-configuration/pull/405)
- Add methods `SetServiceRef` and `GetServiceRef` to `KongRoute` object.
  [#404](https://github.com/Kong/kubernetes-configuration/pull/404)
- Added `WatchNamespaceGrant` CRD.
  [#403](https://github.com/Kong/kubernetes-configuration/pull/403)
- Add `watchNamespaces` field to `GatewayConfiguration`'s `ControlPlane` options.
  [#416](https://github.com/Kong/kubernetes-configuration/pull/416)

### Changes

- Move `KongObjectRef` to `common/v1alpha1` and rename to `NameRef`.
  [#381](https://github.com/Kong/kubernetes-configuration/pull/381)
- Add Type and KonnectID fields to the `KonnectGatewayControlPlane` CRD.
  [#387](https://github.com/Kong/kubernetes-configuration/pull/387)
  [#395](https://github.com/Kong/kubernetes-configuration/pull/395)
- Change validation rules for `KongRoute` to allow its migration from serviceless
  (`KonnectGatewayControlPlane` bound) to `KongService` bound and vice versa.
  [#386](https://github.com/Kong/kubernetes-configuration/pull/386)

## [v1.3.2]

[v1.3.2]: https://github.com/Kong/kubernetes-configuration/compare/v1.3.1...v1.3.2

### Changes

- Removed `namespace` field validation on `KonnectConfigurationDataPlaneGroup` `networkRef` field.
  Using CEL reserved keywords is only available in Kubernetes 1.32+.
  Added CRD validation test for all supported Kubernetes versions.
  [#394](https://github.com/Kong/kubernetes-configuration/pull/394)

## [v1.3.1]

[v1.3.1]: https://github.com/Kong/kubernetes-configuration/compare/v1.3.0...v1.3.1

### Fixes

- Fix rollout in progress validation for `DataPlane`s using blue green deployments.
  [#373](https://github.com/Kong/kubernetes-configuration/pull/373)

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
