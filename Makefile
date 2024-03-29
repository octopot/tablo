# sourced by https://github.com/octomation/makefiles

.DEFAULT_GOAL = test-with-coverage
GIT_HOOKS     = post-merge pre-commit pre-push
GO_VERSIONS   = 1.13 1.14

SHELL := /bin/bash -euo pipefail # `explain set -euo pipefail`

OS   = $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH = $(shell uname -m | tr '[:upper:]' '[:lower:]')

GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPRIVATE   = go.octolab.net
GOPROXY     = direct
LOCAL       = $(MODULE)
MODULE      = `GO111MODULE=on go list -m $(GOFLAGS)`
PACKAGES    = `GO111MODULE=on go list $(GOFLAGS) ./...`
PATHS       = $(shell echo $(PACKAGES) | sed -e "s|$(MODULE)/||g" | sed -e "s|$(MODULE)|$(PWD)/*.go|g")
TIMEOUT     = 1s

ifeq (, $(PACKAGES))
	PACKAGES = $(MODULE)
endif

ifeq (, $(PATHS))
	PATHS = .
endif

export GO111MODULE := $(GO111MODULE)
export GOFLAGS     := $(GOFLAGS)
export GOPRIVATE   := $(GOPRIVATE)
export GOPROXY     := $(GOPROXY)

.PHONY: go-env
go-env:
	@echo "GO111MODULE: `go env GO111MODULE`"
	@echo "GOFLAGS:     $(strip `go env GOFLAGS`)"
	@echo "GOPRIVATE:   $(strip `go env GOPRIVATE`)"
	@echo "GOPROXY:     $(strip `go env GOPROXY`)"
	@echo "LOCAL:       $(LOCAL)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATHS:       $(strip $(PATHS))"
	@echo "TIMEOUT:     $(TIMEOUT)"

.PHONY: deps-check
deps-check:
	@go mod verify
	@if command -v egg > /dev/null; then \
		egg deps check license; \
		egg deps check version; \
	fi

.PHONY: deps-clean
deps-clean:
	@go clean -modcache

.PHONY: deps-fetch
deps-fetch:
	@go mod download
	@if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi

.PHONY: deps-tidy
deps-tidy:
	@go mod tidy
	@if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi

.PHONY: deps-update
deps-update: selector = '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}'
deps-update:
	@if command -v egg > /dev/null; then \
		packages="`egg deps list`"; \
	else \
		packages="`go list -f $(selector) -m -mod=readonly all`"; \
	fi; \
	if [[ "`go version`" == *1.1[1-3]* ]]; then \
		go get -d -mod= -u $$packages; \
	else \
		go get -d -u $$packages; \
	fi; \
	if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi

.PHONY: deps-update-all
deps-update-all:
	@if [[ "`go version`" == *1.1[1-3]* ]]; then \
		go get -d -mod= -u ./...; \
	else \
		go get -d -u ./...; \
	fi; \
	if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi

.PHONY: go-fmt
go-fmt:
	@if command -v goimports > /dev/null; then \
		goimports -local $(LOCAL) -ungroup -w $(PATHS); \
	else \
		gofmt -s -w $(PATHS); \
	fi

.PHONY: go-generate
go-generate:
	@go generate $(PACKAGES)

.PHONY: lint
lint:
	@if command -v golangci-lint > /dev/null; then \
		`which time` -- golangci-lint run ./...; \
	else \
		go vet $(PACKAGES); \
	fi

.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-clean
test-clean:
	@go clean -testcache

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

BINARY  = $(BINPATH)/$(shell basename $(MAIN))
BINPATH = $(PWD)/bin/$(OS)/$(ARCH)
COMMIT  = $(shell git rev-parse --verify HEAD)
DATE    = $(shell date +%Y-%m-%dT%T%Z)
LDFLAGS = -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)"
MAIN    = $(MODULE)

export GOBIN := $(BINPATH)
export PATH  := $(BINPATH):$(PATH)

.PHONY: build-env
build-env:
	@echo "BINARY:      $(BINARY)"
	@echo "BINPATH:     $(BINPATH)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "DATE:        $(DATE)"
	@echo "GOBIN:       `go env GOBIN`"
	@echo "LDFLAGS:     $(LDFLAGS)"
	@echo "MAIN:        $(MAIN)"
	@echo "PATH:        $$PATH"

.PHONY: build
build:
	@go build -o $(BINARY) $(LDFLAGS) $(MAIN)

.PHONY: build-clean
build-clean:
	@rm -f $(BINARY)

.PHONY: install
install:
	@go install $(LDFLAGS) $(MAIN)

.PHONY: install-clean
install-clean:
	@go clean -cache

.PHONY: dist-check
dist-check:
	@goreleaser --snapshot --skip-publish --rm-dist

.PHONY: dist-dump
dist-dump:
	@godownloader .goreleaser.yml > bin/install

TOOLFLAGS = -mod=

.PHONY: tools-env
tools-env:
	@echo "GOBIN:       `go env GOBIN`"
	@echo "TOOLFLAGS:   $(TOOLFLAGS)"

.PHONY: toolset
toolset:
	@( \
		GOFLAGS=$(TOOLFLAGS); \
		cd tools; \
		go mod tidy; \
		go mod download; \
		if [[ "`go env GOFLAGS`" =~ -mod=vendor ]]; then go mod vendor; fi; \
		go generate tools.go; \
	)

ifdef GIT_HOOKS

.PHONY: hooks
hooks:
	@ls .git/hooks | grep -v .sample | sed 's|.*|.git/hooks/&|' | xargs rm -f || true
	@for hook in $(GIT_HOOKS); do cp githooks/$$hook .git/hooks/; done

define hook_tpl
.PHONY: $(1)
$(1):
	@githooks/$(1)
endef

render_hook_tpl = $(eval $(call hook_tpl,$(hook)))
$(foreach hook,$(GIT_HOOKS),$(render_hook_tpl))

endif

ifdef GO_VERSIONS

define go_tpl
.PHONY: go$(1)
go$(1):
	@docker run \
		--rm -it \
		-v $(PWD):/src \
		-w /src \
		golang:$(1) bash
endef

render_go_tpl = $(eval $(call go_tpl,$(version)))
$(foreach version,$(GO_VERSIONS),$(render_go_tpl))

endif


.PHONY: init
init: deps test lint hooks

.PHONY: clean
clean: build-clean deps-clean install-clean test-clean

.PHONY: deps
deps: deps-fetch toolset

.PHONY: env
env: go-env build-env tools-env

.PHONY: format
format: go-fmt

.PHONY: generate
generate: go-generate format

.PHONY: refresh
refresh: deps-tidy update deps generate format test build

.PHONY: update
update: deps-update
