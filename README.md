# kubernetes-configuration

This repository holds the API definitions for Kong's Kubernetes configuration.

## Repository structure

- [`api/`][api] directory contains Go types that are the source for generating
  - [`pkg/clientset`][clientset]: Go clientsets for users who want to interact
    with Kong's Kubernetes configuration in Go
- [`config/crd`][crd]: Kubernetes CRDs for all supported [channels]
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
[channels]: #channels

## Channels

This repository supports multiple _channels_ of CRDs. Each channel is an independent collection
of CRDs that is meant to be used by a designated product or project. Manifests for each channel
are stored in a separate directory under `config/crd/<channel>` (each has a generated `kustomize.yaml` file as well).

The following channels are supported:
- `ingress-controller` - CRDs for [Kong Ingress Controller][kic]
- `ingress-controller-incubator` - experimental CRDs for [Kong Ingress Controller][kic]
- `gateway-operator` - CRDs for [Kong Gateway Operator][kgo]

A single CRD can be included in multiple channels. See [available custom markers](#available-custom-markers) for more details.

[kic]: https://github.com/kong/kubernetes-ingress-controller
[kgo]: https://github.com/kong/gateway-operator

## Install CRDs

In order to install the CRDs from this repo, you can use the following command, replacing 
`<channel>` with one of the supported [channel names](#channels).

```terminal
kustomize build github.com/kong/kubernetes-configuration/config/crd/<channel> | kubectl apply -f -
```

For example, to install the CRDs for the Kong Ingress Controller, you can run:

```terminal
kustomize build github.com/kong/kubernetes-configuration/config/crd/ingress-controller | kubectl apply -f -
```

## Generate code

In order to run code generation in this repo you can use `make generate`.

## Adding new CRDs

When you add a new CRD make sure to

- Annotate the CRD with `+kong:channels` marker to specify the channels that the CRD should be published to.
- Add unit tests in [`test/unit`][testunit]
- Add CRD validation tests in [`test/crdsvalidation`][testcrdsvalidation]
- If you want `GetItems() []T` helper to be generated for your type,
  add it to [supported type list][apitypes_funcs_gen_list].
- If this CRD is meant to have Konnect helpers functions generated for it,
  add it in [supported type list][apitypes_funcs_gen].
- Annotate the CRD and any new type it depends on with the right markers to make sure it will be included 
  in the generated documentation. See [available markers](#Available custom markers).

[crd_kustomization]: ./config/crd/kustomization.yaml
[apitypes_funcs_gen]: ./scripts/apitypes-funcs/supportedtypes.go
[apitypes_funcs_gen_list]: ./scripts/apitypes-funcs/supportedtypes.go#L112-114

## How to release?

- Make sure a changelog is updated with the new version, the release date, and all the changes.
- Trigger a [release workflow]. This will create a tag a release in GitHub.

[release workflow]: ./.github/workflows/release.yaml

## Available custom markers

| Name                        | Applies to | Meaning                                                                                                                                                                                                                                                                |
|-----------------------------|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `+kong:channels`            | Types      | Any root object type annotated with this marker will be included in a [CRD channel](#channels) passed as a marker value (e.g. `+kong:channels=ingress-controller;gateway-operator` will include the CRD both in `ingress-controller` and `gateway-operator` channels). |
| `+apireference:kgo:exclude` | Fields     | Any field annotated with this marker will be excluded from the [KGO's generated CRDs reference][kgo-crd-ref].                                                                                                                                                          |
| `+apireference:kgo:include` | Types      | Any type annotated with this marker will be included in the [KGO's generated CRDs reference][kgo-crd-ref].                                                                                                                                                             |
| `+apireference:kic:exclude` | Fields     | Any type annotated with this marker will be excluded from the [KIC's generated CRDs reference][kic-crd-ref].                                                                                                                                                           |
| `+apireference:kic:include` | Types      | Any type annotated with this marker will be included in the [KIC's generated CRDs reference][kic-crd-ref].                                                                                                                                                             |

### Why do we need separate markers for API reference and channels?

Channels are used to group CRDs into logical sets that are meant to be used by a specific product or project.
API reference markers are used to control which types and fields are included in the generated API reference documentation.
While the channels are enough to be defined on a root object type, the API reference markers need to be defined on
each type or field that should be included/excluded in the generated API reference documentation.

Currently, we don't have a way to automatically generate API reference documentation based only on channels,
so we need separate markers for this purpose.

[kgo-crd-ref]: https://github.com/Kong/gateway-operator/blob/main/docs/api-reference.md
[kic-crd-ref]: https://github.com/kong/kubernetes-ingress-controller/blob/main/docs/api-reference.md
