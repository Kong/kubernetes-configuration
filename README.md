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
- Add unit tests in [`test/unit`][testunit]
- Add CRD validation tests in [`test/crdsvalidation`][testcrdsvalidation]
- If this CRD is meant to have Konnect helpers functions generated for it,
  add it in [konnect-funcs supported type list][konnect_funcs_gen].
- Annotate the CRD and any new type it depends on with the right markers to make sure it will be included 
  in the generated documentation. See [available markers](#Available custom markers).

[crd_kustomization]: ./config/crd/kustomization.yaml
[konnect_funcs_gen]: ./scripts/konnect-funcs/supportedtypes.go

## How to release?

Currently in order to make a new release/tag available for users to use is to
create a new tag and push it to the repository.

## Available custom markers

| Name                        | Applies to | Meaning                                                                                                       |
|-----------------------------|------------|---------------------------------------------------------------------------------------------------------------|
| `+apireference:kgo:exclude` | Fields     | Any field annotated with this marker will be excluded from the [KGO's generated CRDs reference][kgo-crd-ref]. |
| `+apireference:kgo:include` | Types      | Any type annotated with this marker will be included in the [KGO's generated CRDs reference][kgo-crd-ref].    |

[kgo-crd-ref]: https://github.com/Kong/gateway-operator/blob/main/docs/api-reference.md
