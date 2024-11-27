# Table of Contents

<!---
Adding a new version? You'll need three changes:
* Add the ToC link, like "[v1.2.3](#v123)".
* Add the section header, like "## [v1.2.3]".
* Add the diff link, like "[v2.7.0]: https://github.com/kong/kubernetes-ingress-controller/compare/v1.2.2...v1.2.3".
--->
- [v1.0.0-rc.0](#v100-rc0)

## Unreleased

> Release date: TBA

## [v1.0.0-rc.0]

[v1.0.0-rc.0]: https://github.com/kong/kubernetes-configuration/compare/ecf9b7bd62bfb92327a6ddd9aeaec9f73fc13a72...v1.0.0-rc.0

> Release date: 2024-11-26

This is an initial stable release of the Kong Kubernetes configuration module.
It includes CRDs for the Kong Ingress Controller and the Kong Gateway Operator
in the `ingress-controller`, `ingress-controller-incubator` and `gateway-operator`
channels.

Go bindings for the CRDs are available in the [`api/`][api] directory and the
[`pkg/clientset/`][clientset] directory contains the clientset for interacting
with the CRDs.

[api]: ./api/
[clientset]: ./pkg/clientset/
