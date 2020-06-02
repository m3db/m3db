SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# Grab necessary submodules, in case the repo was cloned without --recursive
$(SELF_DIR)/.ci/common.mk:
	git submodule update --init --recursive

include $(SELF_DIR)/.ci/common.mk

SHELL=/bin/bash -o pipefail
GOPATH=$(shell eval $$(go env | grep GOPATH) && echo $$GOPATH)

auto_gen             := scripts/auto-gen.sh
process_coverfile    := scripts/process-cover.sh
gopath_prefix        := $(GOPATH)/src
gopath_bin_path      := $(GOPATH)/bin
m3_package           := github.com/m3db/m3
m3_package_path      := $(gopath_prefix)/$(m3_package)
mockgen_package      := github.com/golang/mock/mockgen
retool_path          := $(m3_package_path)/_tools
retool_bin_path      := $(retool_path)/bin
combined_bin_paths   := $(retool_bin_path):$(gopath_bin_path)
retool_src_prefix    := $(m3_package_path)/_tools/src
retool_package       := github.com/twitchtv/retool
metalint_check       := .ci/metalint.sh
metalint_config      := .metalinter.json
metalint_exclude     := .excludemetalint
mocks_output_dir     := generated/mocks
mocks_rules_dir      := generated/mocks
proto_output_dir     := generated/proto
proto_rules_dir      := generated/proto
assets_output_dir    := generated/assets
assets_rules_dir     := generated/assets
thrift_output_dir    := generated/thrift/rpc
thrift_rules_dir     := generated/thrift
vendor_prefix        := vendor
bad_trace_dep        := go.etcd.io/etcd/vendor/golang.org/x/net/trace
bad_prom_vendor_dir  := github.com/prometheus/prometheus/vendor
cache_policy         ?= recently_read
genny_target         ?= genny-all

BUILD                     := $(abspath ./bin)
VENDOR                    := $(m3_package_path)/$(vendor_prefix)
GO_BUILD_LDFLAGS_CMD      := $(abspath ./scripts/go-build-ldflags.sh)
GO_BUILD_LDFLAGS          := $(shell $(GO_BUILD_LDFLAGS_CMD) LDFLAG)
GO_BUILD_COMMON_ENV       := CGO_ENABLED=0
LINUX_AMD64_ENV           := GOOS=linux GOARCH=amd64 $(GO_BUILD_COMMON_ENV)
GO_RELEASER_DOCKER_IMAGE  := goreleaser/goreleaser:v0.117.2
GO_RELEASER_WORKING_DIR   := /go/src/github.com/m3db/m3
GOMETALINT_VERSION        := v2.0.5

# Retool will look for tools.json in the nearest parent git directory if not
# explicitly told the current dir. Allow setting the base dir so that tools can
# be built inside of other external repos.
ifdef RETOOL_BASE_DIR
	retool_base_args := -base-dir $(RETOOL_BASE_DIR)
endif

export NPROC := 2 # Maximum package concurrency for unit tests.

SERVICES :=     \
	m3dbnode      \
	m3coordinator \
	m3aggregator  \
	m3query       \
	m3collector   \
	m3em_agent    \
	m3nsch_server \
	m3nsch_client \
	m3comparator  \
	r2ctl         \

SUBDIRS :=    \
	x           \
	cluster     \
	msg         \
	metrics     \
	cmd         \
	collector   \
	dbnode      \
	query       \
	m3em        \
	m3nsch      \
	m3ninx      \
	aggregator  \
	ctl         \
	# Disabled during kubeval dependency issue https://github.com/m3db/m3/issues/2220
	# kube        \

TOOLS :=               \
	read_ids             \
	read_index_ids       \
	read_data_files      \
	read_index_files     \
	read_index_segments  \
	clone_fileset        \
	dtest                \
	verify_data_files    \
	verify_index_files   \
	carbon_load          \
	docs_test            \
	m3ctl                \

.PHONY: setup
setup:
	mkdir -p $(BUILD)

.PHONY: install-vendor-m3
install-vendor-m3:
	[ -d $(VENDOR) ] || make install-vendor
	# See comment for "install-vendor-m3-remove-bad-dep" why required and the TODO.
	make install-vendor-m3-remove-bad-dep
	# See comment for "install-vendor-m3-remove-prometheus-vendor-dir" why required.
	make install-vendor-m3-remove-prometheus-vendor-dir

# Some deps were causing panics when using GRPC and etcd libraries were used.
# See issue: https://github.com/etcd-io/etcd/issues/9357
# TODO: Move M3 to go mod to avoid the issue entirely instead of this hack
# (which is bad and we should feel bad).
# $ go test -v
# panic: /debug/requests is already registered. You may have two independent
# copies of golang.org/x/net/trace in your binary, trying to maintain separate
# state. This may involve a vendored copy of golang.org/x/net/trace.
#
# goroutine 1 [running]:
# github.com/m3db/m3/vendor/go.etcd.io/etcd/vendor/golang.org/x/net/trace.init.0()
#         /Users/r/go/src/github.com/m3db/m3/vendor/go.etcd.io/etcd/vendor/golang.org/x/net/trace/trace.go:123 +0x1cd
# exit status 2
# FAIL    github.com/m3db/m3/src/query/remote     0.024s
.PHONY: install-vendor-m3-remove-bad-dep
install-vendor-m3-remove-bad-dep:
	([ -d $(VENDOR)/$(bad_trace_dep) ] && rm -rf $(VENDOR)/$(bad_trace_dep)) || (echo "No bad trace dep" > /dev/null)

# Note: Prometheus has an entire copy of all vendored code which makes
# it impossible to pass sub-dependencies on it since you'll get errors like:
#   have MustRegister(... vendor/github.com/prometheus/client_golang/prometheus.Collector)
#   want MustRegister(... vendor/github.com/prometheus/prometheus/vendor/github.com/prometheus/client_golang/prometheus
# Even if you have the same deps as prometheus you can't pass the dep types to
# it since it depends on the concrete subdirectory vendored code import path.
# Therefore we delete the vendored code and make it rely on our own dependencies
# we install.
.PHONY: install-vendor-m3-remove-prometheus-vendor-dir
install-vendor-m3-remove-prometheus-vendor-dir:
	([ -d $(VENDOR)/$(bad_prom_vendor_dir) ] && rm -rf $(VENDOR)/$(bad_prom_vendor_dir)) || (echo "No bad prom vendor dir" > /dev/null)

.PHONY: docker-dev-prep
docker-dev-prep:
	mkdir -p ./bin/config

	# Hacky way to find all configs and put into ./bin/config/
	find ./src | fgrep config | fgrep ".yml" | xargs -I{} cp {} ./bin/config/
	find ./src | fgrep config | fgrep ".yaml" | xargs -I{} cp {} ./bin/config/

define SERVICE_RULES

.PHONY: $(SERVICE)
$(SERVICE): setup
ifeq ($(SERVICE),m3ctl)
	@echo "Building $(SERVICE) dependencies"
	make build-ui-ctl-statik-gen
endif
	@echo Building $(SERVICE)
	[ -d $(VENDOR) ] || make install-vendor-m3
	$(GO_BUILD_COMMON_ENV) go build -ldflags '$(GO_BUILD_LDFLAGS)' -o $(BUILD)/$(SERVICE) ./src/cmd/services/$(SERVICE)/main/.

.PHONY: $(SERVICE)-linux-amd64
$(SERVICE)-linux-amd64:
	$(LINUX_AMD64_ENV) make $(SERVICE)

.PHONY: $(SERVICE)-docker-dev
$(SERVICE)-docker-dev: clean-build $(SERVICE)-linux-amd64
	make docker-dev-prep

	# Build development docker image
	docker build -t $(SERVICE):dev -t quay.io/m3dbtest/$(SERVICE):dev-$(USER) -f ./docker/$(SERVICE)/development.Dockerfile ./bin

.PHONY: $(SERVICE)-docker-dev-push
$(SERVICE)-docker-dev-push: $(SERVICE)-docker-dev
	docker push quay.io/m3dbtest/$(SERVICE):dev-$(USER)
	@echo "Pushed quay.io/m3dbtest/$(SERVICE):dev-$(USER)"

endef

$(foreach SERVICE,$(SERVICES),$(eval $(SERVICE_RULES)))

define TOOL_RULES

.PHONY: $(TOOL)
$(TOOL): setup
	@echo Building $(TOOL)
	go build -o $(BUILD)/$(TOOL) ./src/cmd/tools/$(TOOL)/main/.

.PHONY: $(TOOL)-linux-amd64
$(TOOL)-linux-amd64:
	$(LINUX_AMD64_ENV) make $(TOOL)

endef

$(foreach TOOL,$(TOOLS),$(eval $(TOOL_RULES)))

.PHONY: services services-linux-amd64
services: $(SERVICES)
services-linux-amd64:
	$(LINUX_AMD64_ENV) make services

.PHONY: tools tools-linux-amd64
tools: $(TOOLS)
tools-linux-amd64:
	$(LINUX_AMD64_ENV) make tools

.PHONY: all
all: metalint test-ci-unit test-ci-integration services tools
	@echo Made all successfully

.PHONY: install-tools
install-tools:
	@echo "Installing retool dependencies"
	GOBIN=$(retool_bin_path) go install github.com/fossas/fossa-cli/cmd/fossa
	GOBIN=$(retool_bin_path) go install github.com/golang/mock/mockgen
	GOBIN=$(retool_bin_path) go install github.com/google/go-jsonnet/cmd/jsonnet
	GOBIN=$(retool_bin_path) go install github.com/m3db/build-tools/utilities/genclean
	GOBIN=$(retool_bin_path) go install github.com/m3db/tools/update-license
	GOBIN=$(retool_bin_path) go install github.com/mauricelam/genny
	GOBIN=$(retool_bin_path) go install github.com/mjibson/esc
	GOBIN=$(retool_bin_path) go install github.com/pointlander/peg
	GOBIN=$(retool_bin_path) go install github.com/rakyll/statik

.PHONY: install-gometalinter
install-gometalinter:
	@mkdir -p $(retool_bin_path)
	./scripts/install-gometalinter.sh -b $(retool_bin_path) -d $(GOMETALINT_VERSION)

.PHONY: check-for-goreleaser-github-token
check-for-goreleaser-github-token:
  ifndef GITHUB_TOKEN
		echo "Usage: make GITHUB_TOKEN=\"<TOKEN>\" release"
		exit 1
  endif

.PHONY: release
release: check-for-goreleaser-github-token
	@echo Releasing new version
	$(GO_BUILD_LDFLAGS_CMD) ECHO > $(BUILD)/release-vars.env
	docker run -e "GITHUB_TOKEN=$(GITHUB_TOKEN)" --env-file $(BUILD)/release-vars.env -v $(PWD):$(GO_RELEASER_WORKING_DIR) -w $(GO_RELEASER_WORKING_DIR) $(GO_RELEASER_DOCKER_IMAGE) release --rm-dist

.PHONY: release-snapshot
release-snapshot: check-for-goreleaser-github-token
	@echo Creating snapshot release
	docker run -e "GITHUB_TOKEN=$(GITHUB_TOKEN)" -v $(PWD):$(GO_RELEASER_WORKING_DIR) -w $(GO_RELEASER_WORKING_DIR) $(GO_RELEASER_DOCKER_IMAGE) --snapshot --rm-dist

.PHONY: docs-container
docs-container:
	docker run --rm hello-world >/dev/null
	docker build -t m3db-docs docs

# NB(schallert): if updating this target, be sure to update the commands used in
# the .buildkite/docs_push.sh. We can't share the make targets because our
# Makefile assumes its running under bash and the container is alpine (ash
# shell).
.PHONY: docs-build
docs-build: docs-container
	docker run -v $(PWD):/m3db --rm m3db-docs "mkdocs build -e docs/theme -t material"

.PHONY: docs-serve
docs-serve: docs-container
	docker run -v $(PWD):/m3db -p 8000:8000 -it --rm m3db-docs "mkdocs serve -e docs/theme -t material -a 0.0.0.0:8000"

.PHONY: docs-deploy
docs-deploy: docs-container
	docker run -v $(PWD):/m3db --rm -v $(HOME)/.ssh/id_rsa:/root/.ssh/id_rsa:ro -it m3db-docs "mkdocs build -e docs/theme -t material && mkdocs gh-deploy --force --dirty"

.PHONY: docs-validate
docs-validate: docs_test
	./bin/docs_test

.PHONY: docs-test
docs-test:
	@echo "--- Documentation validate test"
	make docs-validate
	@echo "--- Documentation build test"
	make docs-build

.PHONY: docker-integration-test
docker-integration-test:
	@echo "--- Running Docker integration test"
	./scripts/docker-integration-tests/run.sh


.PHONY: docker-compatibility-test
docker-compatibility-test:
	@echo "--- Running Prometheus compatibility test"
	./scripts/comparator/run.sh

.PHONY: site-build
site-build:
	@echo "Building site"
	@./scripts/site-build.sh

# Generate configs in config/
.PHONY: config-gen
config-gen: install-tools
	@echo "--- Generating configs"
	$(retool_bin_path)/jsonnet -S $(m3_package_path)/config/m3db/local-etcd/m3dbnode_cmd.jsonnet > $(m3_package_path)/config/m3db/local-etcd/generated.yaml
	$(retool_bin_path)/jsonnet -S $(m3_package_path)/config/m3db/clustered-etcd/m3dbnode_cmd.jsonnet > $(m3_package_path)/config/m3db/clustered-etcd/generated.yaml

SUBDIR_TARGETS := \
	mock-gen        \
	thrift-gen      \
	proto-gen       \
	asset-gen       \
	genny-gen       \
	license-gen     \
	all-gen         \
	metalint

.PHONY: test-ci-unit
test-ci-unit: test-base
	$(process_coverfile) $(coverfile)

.PHONY: test-ci-big-unit
test-ci-big-unit: test-big-base
	$(process_coverfile) $(coverfile)

.PHONY: test-ci-integration
test-ci-integration:
	INTEGRATION_TIMEOUT=4m TEST_SERIES_CACHE_POLICY=$(cache_policy) make test-base-ci-integration
	$(process_coverfile) $(coverfile)

define SUBDIR_RULES

# Temporarily remove kube validation until we fix a dependency issue with
# kubeval (one of its depenencies depends on go1.13).
# https://github.com/m3db/m3/issues/2220
#
# We override the rules for `*-gen-kube` to just generate the kube manifest
# bundle.
# ifeq ($(SUBDIR), kube)

# Builds the single kube bundle from individual manifest files.
# all-gen-kube: install-tools
# 	@echo "--- Generating kube bundle"
# 	@./kube/scripts/build_bundle.sh
# 	find kube -name '*.yaml' -print0 | PATH=$(combined_bin_paths):$(PATH) xargs -0 kubeval -v=1.12.0

# else

.PHONY: mock-gen-$(SUBDIR)
mock-gen-$(SUBDIR): install-tools
	@echo "--- Generating mocks $(SUBDIR)"
	@[ ! -d src/$(SUBDIR)/$(mocks_rules_dir) ] || \
		PATH=$(combined_bin_paths):$(PATH) PACKAGE=$(m3_package) $(auto_gen) src/$(SUBDIR)/$(mocks_output_dir) src/$(SUBDIR)/$(mocks_rules_dir)

.PHONY: thrift-gen-$(SUBDIR)
thrift-gen-$(SUBDIR): install-tools
	@echo "--- Generating thrift files $(SUBDIR)"
	@[ ! -d src/$(SUBDIR)/$(thrift_rules_dir) ] || \
		PATH=$(combined_bin_paths):$(PATH) PACKAGE=$(m3_package) $(auto_gen) src/$(SUBDIR)/$(thrift_output_dir) src/$(SUBDIR)/$(thrift_rules_dir)

.PHONY: proto-gen-$(SUBDIR)
proto-gen-$(SUBDIR): install-tools
	@echo "--- Generating protobuf files $(SUBDIR)"
	@[ ! -d src/$(SUBDIR)/$(proto_rules_dir) ] || \
		PATH=$(combined_bin_paths):$(PATH) PACKAGE=$(m3_package) $(auto_gen) src/$(SUBDIR)/$(proto_output_dir) src/$(SUBDIR)/$(proto_rules_dir)

.PHONY: asset-gen-$(SUBDIR)
asset-gen-$(SUBDIR): install-tools
	@echo "--- Generating asset files $(SUBDIR)"
	@[ ! -d src/$(SUBDIR)/$(assets_rules_dir) ] || \
		PATH=$(combined_bin_paths):$(PATH) PACKAGE=$(m3_package) $(auto_gen) src/$(SUBDIR)/$(assets_output_dir) src/$(SUBDIR)/$(assets_rules_dir)

.PHONY: genny-gen-$(SUBDIR)
genny-gen-$(SUBDIR): install-tools
	@echo "--- Generating genny files $(SUBDIR)"
	@[ ! -f $(SELF_DIR)/src/$(SUBDIR)/generated-source-files.mk ] || \
		PATH=$(combined_bin_paths):$(PATH) GO111MODULE=on make -f $(SELF_DIR)/src/$(SUBDIR)/generated-source-files.mk $(genny_target)
	@PATH=$(combined_bin_paths):$(PATH) GO111MODULE=on bash -c "source ./scripts/auto-gen-helpers.sh && gen_cleanup_dir '*_gen.go' $(SELF_DIR)/src/$(SUBDIR)/ && gen_cleanup_dir '*_gen_test.go' $(SELF_DIR)/src/$(SUBDIR)/"

.PHONY: license-gen-$(SUBDIR)
license-gen-$(SUBDIR): install-tools
	@echo "--- Updating license in files $(SUBDIR)"
	@find $(SELF_DIR)/src/$(SUBDIR) -name '*.go' | PATH=$(combined_bin_paths):$(PATH) xargs -I{} update-license {}

.PHONY: all-gen-$(SUBDIR)
# NB(prateek): order matters here, mock-gen needs to be after proto/thrift because we sometimes
# generate mocks for thrift/proto generated code. Similarly, license-gen needs to be last because
# we make header changes.
all-gen-$(SUBDIR): thrift-gen-$(SUBDIR) proto-gen-$(SUBDIR) asset-gen-$(SUBDIR) genny-gen-$(SUBDIR) mock-gen-$(SUBDIR) license-gen-$(SUBDIR)

.PHONY: test-$(SUBDIR)
test-$(SUBDIR):
	@echo testing $(SUBDIR)
	SRC_ROOT=./src/$(SUBDIR) make test-base
	gocov convert $(coverfile) | gocov report

.PHONY: test-xml-$(SUBDIR)
test-xml-$(SUBDIR):
	@echo test-xml $(SUBDIR)
	SRC_ROOT=./src/$(SUBDIR) make test-base-xml

.PHONY: test-html-$(SUBDIR)
test-html-$(SUBDIR):
	@echo test-html $(SUBDIR)
	SRC_ROOT=./src/$(SUBDIR) make test-base-html

.PHONY: test-integration-$(SUBDIR)
test-integration-$(SUBDIR):
	@echo test-integration $(SUBDIR)
	SRC_ROOT=./src/$(SUBDIR) make test-base-integration

# Usage: make test-single-integration name=<test_name>
.PHONY: test-single-integration-$(SUBDIR)
test-single-integration-$(SUBDIR):
	SRC_ROOT=./src/$(SUBDIR) make test-base-single-integration name=$(name)

.PHONY: test-ci-unit-$(SUBDIR)
test-ci-unit-$(SUBDIR):
	@echo "--- test-ci-unit $(SUBDIR)"
	SRC_ROOT=./src/$(SUBDIR) make test-base
	@echo "--- uploading coverage report"
	$(codecov_push) -f $(coverfile) -F $(SUBDIR)

.PHONY: test-ci-big-unit-$(SUBDIR)
test-ci-big-unit-$(SUBDIR):
	@echo "--- test-ci-big-unit $(SUBDIR)"
	SRC_ROOT=./src/$(SUBDIR) make test-big-base
	@echo "--- uploading coverage report"
	$(codecov_push) -f $(coverfile) -F $(SUBDIR)

.PHONY: test-ci-integration-$(SUBDIR)
test-ci-integration-$(SUBDIR):
	@echo "--- test-ci-integration $(SUBDIR)"
	SRC_ROOT=./src/$(SUBDIR) PANIC_ON_INVARIANT_VIOLATED=true INTEGRATION_TIMEOUT=4m TEST_SERIES_CACHE_POLICY=$(cache_policy) make test-base-ci-integration
	@echo "--- uploading coverage report"
	$(codecov_push) -f $(coverfile) -F $(SUBDIR)

.PHONY: metalint-$(SUBDIR)
metalint-$(SUBDIR): install-gometalinter install-linter-badtime install-linter-importorder
	@echo "--- metalinting $(SUBDIR)"
	@(PATH=$(combined_bin_paths):$(PATH) $(metalint_check) \
		$(metalint_config) $(metalint_exclude) src/$(SUBDIR))

# endif kubeval
# endif

endef

# generate targets for each SUBDIR in SUBDIRS based on the rules specified above.
$(foreach SUBDIR,$(SUBDIRS),$(eval $(SUBDIR_RULES)))

define SUBDIR_TARGET_RULE
.PHONY: $(SUBDIR_TARGET)
$(SUBDIR_TARGET): $(foreach SUBDIR,$(SUBDIRS),$(SUBDIR_TARGET)-$(SUBDIR))
endef

# generate targets across SUBDIRS for each SUBDIR_TARGET. i.e. generate rules
# which allow `make all-gen` to invoke `make all-gen-dbnode all-gen-coordinator ...`
# NB: we skip metalint explicity as the default target below requires less invocations
# of metalint and finishes faster.
$(foreach SUBDIR_TARGET, $(filter-out metalint,$(SUBDIR_TARGETS)), $(eval $(SUBDIR_TARGET_RULE)))

.PHONY: build-ui-ctl
build-ui-ctl:
ifeq ($(shell ls ./src/ctl/ui/build 2>/dev/null),)
	# Need to use subshell output of set-node-version as cannot
	# set side-effects of nvm to later commands
	@echo "Building UI components, if npm install or build fails try: npm cache clean"
	make node-yarn-run \
		node_version="6" \
		node_cmd="cd $(m3_package_path)/src/ctl/ui && yarn install && yarn build"
	# If we've installed nvm locally, remove it due to some cleanup permissions
	# issue we run into in CI.
	rm -rf .nvm
else
	@echo "Skip building UI components, already built, to rebuild first make clean"
endif
	# Move public assets into public subdirectory so that it can
	# be included in the single statik package built from ./ui/build
	rm -rf ./src/ctl/ui/build/public
	cp -r ./src/ctl/public ./src/ctl/ui/build/public

.PHONY: build-ui-ctl-statik-gen
build-ui-ctl-statik-gen: build-ui-ctl-statik license-gen-ctl

.PHONY: build-ui-ctl-statik
build-ui-ctl-statik: build-ui-ctl install-tools
	mkdir -p ./src/ctl/generated/ui
	$(retool_bin_path)/statik -m -f -src ./src/ctl/ui/build -dest ./src/ctl/generated/ui -p statik

.PHONY: node-yarn-run
node-yarn-run:
	make node-run \
		node_version="$(node_version)" \
		node_cmd="(yarn --version 2>&1 >/dev/null || npm install -g yarn@^1.17.0) && $(node_cmd)"

.PHONY: node-run
node-run:
ifneq ($(shell command -v nvm 2>/dev/null),)
	@echo "Using nvm to select node version $(node_version)"
	nvm use $(node_version) && bash -c "$(node_cmd)"
else
	mkdir .nvm
	# Install nvm locally
	NVM_DIR=$(SELF_DIR)/.nvm PROFILE=/dev/null scripts/install_nvm.sh
	bash -c "source $(SELF_DIR)/.nvm/nvm.sh; nvm install 6"
	bash -c "source $(SELF_DIR)/.nvm/nvm.sh && nvm use 6 && $(node_cmd)"
endif

.PHONY: metalint
metalint: install-gometalinter install-linter-badtime install-linter-importorder
	@echo "--- metalinting src/"
	@(PATH=$(retool_bin_path):$(PATH) $(metalint_check) \
		$(metalint_config) $(metalint_exclude) $(m3_package_path)/src/)

# Tests that all currently generated types match their contents if they were regenerated
.PHONY: test-all-gen
test-all-gen: all-gen
	@test "$(shell git --no-pager diff --exit-code --shortstat 2>/dev/null)" = "" || (git --no-pager diff --text --exit-code && echo "Check git status, there are dirty files" && exit 1)
	@test "$(shell git status --exit-code --porcelain 2>/dev/null | grep "^??")" = "" || (git status --exit-code --porcelain && echo "Check git status, there are untracked files" && exit 1)

# Runs a fossa license report
.PHONY: fossa
fossa: install-tools
	PATH=$(combined_bin_paths):$(PATH) fossa analyze --verbose --no-ansi

# Waits for the result of a fossa test and exits success if pass or fail if fails
.PHONY: fossa-test
fossa-test: fossa
	PATH=$(combined_bin_paths):$(PATH) fossa test

.PHONY: clean-build
clean-build:
	@rm -rf $(BUILD)

.PHONY: clean
clean: clean-build
	@rm -f *.html *.xml *.out *.test
	@rm -rf $(VENDOR)
	@rm -rf ./src/ctl/ui/build

.DEFAULT_GOAL := all
