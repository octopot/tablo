BIN         = $(shell basename $(shell go list -m))
BINPATH     = $(PWD)/bin
COMMIT      = $(shell git rev-parse --verify HEAD)
DATE        = $(shell date +%Y-%m-%dT%T%Z)
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPROXY     = https://proxy.golang.org,https://gocenter.io,direct
MODULE      = $(shell go list -m)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
SHELL       = /bin/bash -euo pipefail
TIMEOUT     = 1s

export PATH := $(BINPATH):$(PATH)

include env/make/go.mk
include env/make/protobuf.mk
include env/make/docker.mk
include env/make/compose.mk

.DEFAULT_GOAL = test-with-coverage

.PHONY: env
env:
	@echo "BIN:         $(BIN)"
	@echo "BINPATH:     $(BINPATH)"
	@echo "COMMIT:      $(COMMIT)"
	@echo "DATE:        $(DATE)"
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(shell go env GOFLAGS)"
	@echo "GOPRIVATE:   $(shell go env GOPRIVATE)"
	@echo "GOPROXY:     $(shell go env GOPROXY)"
	@echo "GONOPROXY:   $(shell go env GONOPROXY)"
	@echo "GOSUMDB:     $(shell go env GOSUMDB)"
	@echo "GONOSUMDB:   $(shell go env GONOSUMDB)"
	@echo "MODULE:      $(MODULE)"
	@echo "PATH:        $(PATH)"
	@echo "PATHS:       $(PATHS)"
	@echo "SHELL:       $(SHELL)"
	@echo "TIMEOUT:     $(TIMEOUT)"


.PHONY: deps
deps: deps-main deps-tools

.PHONY: deps-main
deps-main:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: deps-tools
deps-tools:
	@cd tools && make

.PHONY: deps-update
deps-update:
	@go get -mod= -u all


.PHONY: format
format:
	@goimports -local $(dir $(shell go list -m)) -ungroup -w $(PATHS)

.PHONY: generate
generate: protobuf-generate generate-go

.PHONY: generate-go
generate-go:
	@go generate $(MODULE)/...

.PHONY: refresh
refresh: deps-update deps generate format test-with-coverage


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(MODULE)/...

.PHONY: test-full
test-full: test-with-coverage test-integration

.PHONY: test-integration
test-integration:
	@go test -cover \
	         -covermode count \
	         -coverprofile i.out \
	         -tags=integration \
	         ./test/... 2> /dev/null

.PHONY: test-smoke
test-smoke:
	@go test -run=^TestSmoke -tags=integration -timeout 100ms $(MODULE)/...

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(MODULE)/... | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(MODULE)/...


.PHONY: build
build:
	@go build -o bin/$(BIN) -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)" .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh
