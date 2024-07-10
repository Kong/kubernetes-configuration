# kubernetes-configuration

This repository holds the API definitions for Kong's Kubernetes configuration.

<<<<<<< HEAD
> 👷 🚧 This is currently a work in progress which is heavily based on [Kong's Ingress Controller][kic] CRDs
=======
> This is currently a work in progress which is heavily based on [Kong's Ingress Controller][kic] CRDs
>>>>>>> 09eb63c (chore: first generation)
> Before KIC starts using these CRDs this repo should contain only additive,
> non-breaking changes on top of KIC's types.

[kic]: https://github.com/Kong/kubernetes-ingress-controller

## Repository structure

- [`api/`][api] directory contains Go types that are the source for generating
  - [`pkg/clientset`][clientset]: Go clientsets for users who want to interact with Kong's Kubernetes configuration in Go
  - [`config/crd`][crd]: Kubernetes CRDs for Kong configuration

[api]: ./api/
[clientset]: ./pkg/clientset/
[crd]: ./config/crd

## Install CRDs

In order to install the CRDs from this repo:

```terminal
kustomize build github.com/kong/kubernetes-configuration/config/crd | kubectl apply -f -
```

## Adding new CRDs

When you add a new CRD make sure to

- Add it to CRD [kustomization.yaml][crd_kustomization]

[crd_kustomization]: ./config/crd/kustomization.yaml
