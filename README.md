# kubernetes-configuration

This repository holds the API definitions for Kong's Kubernetes configuration.

> 👷 🚧 This is currently a work in progress which is heavily based on [Kong's Ingress Controller][kic] CRDs
> Before KIC starts using these CRDs this repo should contain only additive,
> non-breaking changes on top of KIC's types.

[kic]: https://github.com/Kong/kubernetes-ingress-controller

## Repository structure

- [`api/`][api] directory contains Go types that are the source for generating
  - [`pkg/clientset`][clientset]: Go clientsets for users who want to interact
    with Kong's Kubernetes configuration in Go
  - [`config/crd`][crd]: Kubernetes CRDs for Kong configuration
- [`test/`][test] directory contains Go tests
  - [`test/crdsvalidation`][testcrdsvalidation] directory contains Go tests which
    perform operations against a live Kubernetes cluster, testing [CEL][cel] rules
    set on API types
  - [`test/unit`][testunit] directory contains Go unit tests for generated Go types
- [`docs/`][docs] directory contains generated API reference markdown files

[api]: ./api/
[clientset]: ./pkg/clientset/
[crd]: ./config/crd
[docs]: ./docs/
[test]: ./test/
[testcrdsvalidation]: ./test/crdsvalidation
[testunit]: ./test/unit
[cel]: https://kubernetes.io/docs/reference/using-api/cel/

## Install CRDs

In order to install the CRDs from this repo:

```terminal
kustomize build github.com/kong/kubernetes-configuration/config/crd | kubectl apply -f -
```

## Generate code

In order to run code generation in this repo you can use `make generate`.

## Adding new CRDs

When you add a new CRD make sure to

- Add it to CRD [kustomization.yaml][crd_kustomization]

[crd_kustomization]: ./config/crd/kustomization.yaml
