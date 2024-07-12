# ------------------------------------------------------------------------------
# Configuration - Make
# ------------------------------------------------------------------------------

# Some sensible Make defaults: https://tech.davis-hansson.com/p/make/
SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c

# ------------------------------------------------------------------------------
# Configuration - Repository
# ------------------------------------------------------------------------------

MAKEFLAGS += --no-print-directory
REPO_URL ?= github.com/kong/kubernetes-configuration
# NOTE: this is not neeed now before v2, but will be needed in the future.
#GO_MOD_MAJOR_VERSION ?= $(subst $(REPO_URL)/,,$(shell go list -m))
REPO_INFO ?= $(shell git config --get remote.origin.url)

ifndef COMMIT
  COMMIT := $(shell git rev-parse --short HEAD)
endif

# ------------------------------------------------------------------------------
# Configuration - Golang
# ------------------------------------------------------------------------------

export GO111MODULE=on

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# ------------------------------------------------------------------------------
# Configuration - Tooling
# ------------------------------------------------------------------------------

PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
TOOLS_VERSIONS_FILE = .tools_versions.yaml

MISE := $(shell which mise)
.PHONY: mise
mise:
	@mise -V >/dev/null || (echo "mise - https://github.com/jdx/mise - not found. Please install it." && exit 1)

.PHONY: tools
tools: controller-gen kustomize client-gen

export MISE_DATA_DIR = $(PROJECT_DIR)/bin/

# NOTE: mise targets use -q to silence the output.
# Users can use MISE_VERBOSE=1 MISE_DEBUG=1 to get more verbose output.

.PHONY: mise-plugin-install
mise-plugin-install: mise
	@$(MISE) plugin install --yes -q $(DEP) $(URL)

.PHONY: mise-install
mise-install: mise
	@$(MISE) install -q $(DEP_VER)

CONTROLLER_GEN_VERSION = $(shell yq -ojson -r '.controller-tools' < $(TOOLS_VERSIONS_FILE))
CONTROLLER_GEN = $(PROJECT_DIR)/bin/installs/kube-controller-tools/$(CONTROLLER_GEN_VERSION)/bin/controller-gen
.PHONY: controller-gen
controller-gen: mise ## Download controller-gen locally if necessary.
	@$(MAKE) mise-plugin-install DEP=kube-controller-tools
	@$(MAKE) mise-install DEP_VER=kube-controller-tools@$(CONTROLLER_GEN_VERSION)

KUSTOMIZE_VERSION = $(shell yq -ojson -r '.kustomize' < $(TOOLS_VERSIONS_FILE))
KUSTOMIZE = $(PROJECT_DIR)/bin/installs/kustomize/$(KUSTOMIZE_VERSION)/bin/kustomize
.PHONY: kustomize
kustomize: mise ## Download kustomize locally if necessary.
	@$(MAKE) mise-plugin-install DEP=kustomize
	@$(MAKE) mise-install DEP_VER=kustomize@$(KUSTOMIZE_VERSION)

CLIENT_GEN_VERSION = $(shell yq -ojson -r '.kube-code-generator' < $(TOOLS_VERSIONS_FILE))
CLIENT_GEN = $(PROJECT_DIR)/bin/installs/kube-code-generator/$(CLIENT_GEN_VERSION)/bin/client-gen
.PHONY: client-gen
client-gen: mise ## Download client-gen locally if necessary.
	@$(MAKE) mise-plugin-install DEP=kube-code-generator
	@$(MAKE) mise-install DEP_VER=kube-code-generator@$(CLIENT_GEN_VERSION)

YQ_VERSION = $(shell yq -ojson -r '.yq' < $(TOOLS_VERSIONS_FILE))
YQ = $(PROJECT_DIR)/bin/installs/yq/$(YQ_VERSION)/bin/yq
.PHONY: yq
yq: mise # Download yq locally if necessary.
	@$(MAKE) mise-plugin-install DEP=yq
	@$(MAKE) mise-install DEP_VER=yq@$(YQ_VERSION)

CRD_REF_DOCS_VERSION = $(shell yq -ojson -r '.crd-ref-docs' < $(TOOLS_VERSIONS_FILE))
CRD_REF_DOCS = $(PROJECT_DIR)/bin/crd-ref-docs
.PHONY: crd-ref-docs
crd-ref-docs: ## Download crd-ref-docs locally if necessary.
	GOBIN=$(PROJECT_DIR)/bin go install -v \
		github.com/elastic/crd-ref-docs@v$(CRD_REF_DOCS_VERSION)

# ------------------------------------------------------------------------------
# Verify steps
# ------------------------------------------------------------------------------

.PHONY: verify.repo
verify.repo:
	./scripts/verify-repo.sh

.PHONY: verify.diff
verify.diff:
	./scripts/verify-diff.sh

.PHONY: verify.versions
verify.versions:
	./scripts/verify-versions.sh $(TAG)

.PHONY: verify.manifests
verify.manifests: verify.repo manifests verify.diff

.PHONY: verify.generators
verify.generators: verify.repo generate verify.diff

# ------------------------------------------------------------------------------
# Build - Generators
# ------------------------------------------------------------------------------

CRD_GEN_PATHS ?= ./api/configuration/...
CRD_INCUBATOR_GEN_PATHS ?= ./api/incubator/...
CRD_OPTIONS ?= "+crd:allowDangerousTypes=true"

API_DIR ?= api

.PHONY: generate
generate: generate.crds generate.deepcopy generate.clientsets generate.docs

.PHONY: generate.crds
generate.crds: controller-gen ## Generate WebhookConfiguration and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=kong-ingress webhook paths="$(CRD_INCUBATOR_GEN_PATHS)" output:crd:artifacts:config=config/crd/incubator
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=kong-ingress webhook paths="$(CRD_GEN_PATHS)" output:crd:artifacts:config=config/crd/bases

.PHONY: generate.deepcopy
generate.deepcopy: controller-gen
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="$(CRD_GEN_PATHS)"

# this will generate the custom typed clients needed for end-users implementing logic in Go to use our API types.
.PHONY: generate.clientsets
generate.clientsets: client-gen
	$(CLIENT_GEN) \
		--go-header-file ./hack/boilerplate.go.txt \
		--logtostderr \
		--clientset-name clientset \
		--input-base $(REPO_URL)/$(API_DIR)/ \
		--input configuration/v1 \
		--input configuration/v1alpha1 \
		--input configuration/v1beta1 \
		--input incubator/v1alpha1 \
		--output-dir pkg/ \
		--output-pkg $(REPO_URL)/pkg/

.PHONY: generate.docs
generate.docs: generate.apidocs # generate.cli-arguments-docs

.PHONY: generate.apidocs
generate.apidocs: crd-ref-docs
	./scripts/apidocs-gen/generate.sh $(CRD_REF_DOCS)

# NOTE(pmalek): We can't generate CLI args docs because the tool at 
# https://github.com/Kong/kubernetes-ingress-controller/blob/513db87cbf94ce66207f74365b502e8cde841357/scripts/cli-arguments-docs-gen/main.go
#
# relies on an internal KIC package. We could solve this by exportint the internal/manager package.
# .PHONY: generate.cli-arguments
# generate.cli-arguments-docs:
# 	go run ./scripts/cli-arguments-docs-gen/main.go > ./docs/cli-arguments.md
