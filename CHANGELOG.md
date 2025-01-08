# Table of Contents

<!---
Adding a new version? You'll need three changes:
* Add the ToC link, like "[v1.2.3](#v123)".
* Add the section header, like "## [v1.2.3]".
* Add the diff link, like "[v2.7.0]: https://github.com/kong/kubernetes-ingress-controller/compare/v1.2.2...v1.2.3".
--->
- [v1.0.4](#v104)
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.0](#v100)

## [v1.0.4]

[v1.0.4]: https://github.com/Kong/kubernetes-configuration/compare/v1.0.3...v1.0.4

### Changes

- Allowed `konnectID` as `ControlPlaneRef`'s `type` field value.
  [#214](https://github.com/Kong/kubernetes-configuration/pull/214)

> Release date: 2025-01-08

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
