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
VERSION ?= $(shell head -1 VERSION)

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

GOTESTSUM_VERSION = $(shell $(YQ) -r '.gotestsum' < $(TOOLS_VERSIONS_FILE))
GOTESTSUM = $(PROJECT_DIR)/bin/installs/gotestsum/$(GOTESTSUM_VERSION)/bin/gotestsum
.PHONY: gotestsum
gotestsum: mise yq ## Download gotestsum locally if necessary.
	@$(MISE) plugin install --yes -q gotestsum https://github.com/pmalek/mise-gotestsum.git
	@$(MISE) install -q gotestsum@$(GOTESTSUM_VERSION)

GOLANGCI_LINT_VERSION = $(shell $(YQ) -r '.golangci-lint' < $(TOOLS_VERSIONS_FILE))
GOLANGCI_LINT = $(PROJECT_DIR)/bin/installs/golangci-lint/$(GOLANGCI_LINT_VERSION)/bin/golangci-lint
.PHONY: golangci-lint
golangci-lint: mise yq ## Download golangci-lint locally if necessary.
	@$(MISE) plugin install --yes -q golangci-lint
	@$(MISE) install -q golangci-lint@$(GOLANGCI_LINT_VERSION)

ACTIONLINT_VERSION = $(shell $(YQ) -r '.actionlint' < $(TOOLS_VERSIONS_FILE))
ACTIONLINT = $(PROJECT_DIR)/bin/installs/actionlint/$(ACTIONLINT_VERSION)/bin/actionlint
.PHONY: download.actionlint
download.actionlint: mise yq ## Download actionlint locally if necessary.
	@$(MISE) plugin install --yes -q actionlint
	@$(MISE) install -q actionlint@$(ACTIONLINT_VERSION)

SHELLCHECK_VERSION = $(shell $(YQ) -r '.shellcheck' < $(TOOLS_VERSIONS_FILE))
SHELLCHECK = $(PROJECT_DIR)/bin/installs/shellcheck/$(SHELLCHECK_VERSION)/bin/shellcheck
.PHONY: download.shellcheck
download.shellcheck: mise yq ## Download shellcheck locally if necessary.
	@$(MISE) plugin install --yes -q shellcheck
	@$(MISE) install -q shellcheck@$(SHELLCHECK_VERSION)

# ------------------------------------------------------------------------------
# Verify steps
# ------------------------------------------------------------------------------

.PHONY: verify.diff
verify.diff:
	@$(PROJECT_DIR)/scripts/verify-diff.sh $(PROJECT_DIR)

.PHONY: verify.generators
verify.generators: generate verify.diff

# ------------------------------------------------------------------------------
# Build - Generators
# ------------------------------------------------------------------------------

CRD_GEN_PATHS ?= ./api/configuration/...;./api/konnect/...
CRD_INCUBATOR_GEN_PATHS ?= ./api/incubator/...
CRD_OPTIONS ?= "+crd:allowDangerousTypes=true"

API_DIR ?= api

.PHONY: generate
generate: generate.crds generate.deepcopy generate.clientsets generate.docs generate.apitypes-funcs

.PHONY: generate.apitypes-funcs
generate.apitypes-funcs:
	go run ./scripts/apitypes-funcs

.PHONY: generate.crds
generate.crds: controller-gen ## Generate WebhookConfiguration and CustomResourceDefinition objects.
	# Use gotypesalias=0 as a workaround for https://github.com/kubernetes-sigs/controller-tools/issues/1088.
	VERSION=$(VERSION) GODEBUG=gotypesalias=0 go run ./scripts/crds-generator

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
		--input konnect/v1alpha1 \
		--input incubator/v1alpha1 \
		--output-dir pkg/ \
		--output-pkg $(REPO_URL)/pkg/

.PHONY: generate.docs
generate.docs: generate.apidocs

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

# Define a constant list of channels
CHANNELS := ingress-controller ingress-controller-incubator gateway-operator

# Install all CRDs into the K8s cluster specified in ~/.kube/config.
.PHONY: install
install: generate.crds kustomize
	@for channel in $(CHANNELS); do \
		$(KUSTOMIZE) build config/crd/$$channel | kubectl apply -f -; \
	done

GOLANGCI_LINT_CONFIG ?= $(PROJECT_DIR)/.golangci.yaml
.PHONY: lint
lint: golangci-lint
	$(GOLANGCI_LINT) run -v --config $(GOLANGCI_LINT_CONFIG) $(GOLANGCI_LINT_FLAGS)

.PHONY: lint.actions
lint.actions: download.actionlint download.shellcheck
	$(ACTIONLINT) -shellcheck $(SHELLCHECK) ./.github/workflows/*

.PHONY: test.samples
test.samples: kustomize
	find ./config/samples -not -name "kustomization.*" -type f | sort | xargs -I{} bash -c "kubectl apply -f {}; kubectl delete -f {}"

GOTESTSUM_FORMAT ?= standard-verbose

.PHONY: test
test: test.unit test.crds-validation

.PHONY: test.pretty
test.pretty: test.unit.pretty test.crds-validation.pretty

UNIT_TEST_PATHS := ./test/unit/... ./pkg/...
CRDS_VALIDATION_TEST_PATHS := ./test/crdsvalidation/...

.PHONY: _test.unit
_test.unit: gotestsum
	GOTESTSUM_FORMAT=$(GOTESTSUM_FORMAT) \
		$(GOTESTSUM) -- $(GOTESTFLAGS) \
		-race \
		$(UNIT_TEST_PATHS)

.PHONY: _test.crds-validation
_test.crds-validation: gotestsum
	GOTESTSUM_FORMAT=$(GOTESTSUM_FORMAT) \
		$(GOTESTSUM) -- $(GOTESTFLAGS) \
		-race \
		$(CRDS_VALIDATION_TEST_PATHS)

.PHONY: test.unit
test.unit:
	@$(MAKE) _test.unit GOTESTFLAGS="$(GOTESTFLAGS)"

.PHONY: test.crds-validation
test.crds-validation:
	@$(MAKE) _test.crds-validation GOTESTFLAGS="$(GOTESTFLAGS)"

.PHONY: test.unit.pretty
test.unit.pretty:
	@$(MAKE) _test.unit GOTESTSUM_FORMAT=testname \
		GOTESTFLAGS="$(GOTESTFLAGS)" \
		UNIT_TEST_PATHS="$(UNIT_TEST_PATHS)"

.PHONY: test.crds-validation.pretty
test.crds-validation.pretty:
	@$(MAKE) _test.crds-validation GOTESTSUM_FORMAT=testname \
		GOTESTFLAGS="$(GOTESTFLAGS)" \
		UNIT_TEST_PATHS="$(UNIT_TEST_PATHS)"
