BIN         = $(shell basename $(shell go list -m))
BINPATH     = $(PWD)/bin
COMMIT      = $(shell git rev-parse --verify HEAD)
DATE        = $(shell date +%Y-%m-%dT%T%Z)
GO111MODULE = on
GOFLAGS     = -mod=vendor
GOPROXY     = https://proxy.golang.org,https://gocenter.io,direct
MODULE      = $(shell go list -m)
PACKAGES    = $(shell go list ./...)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/\{0,1\}||g")
SHELL       = /bin/bash -euo pipefail
TIMEOUT     = 1s

export PATH := $(BINPATH):$(PATH)

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
	@echo "PACKAGES:    $(PACKAGES)"
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
generate: generate-proto generate-go

.PHONY: generate-proto
generate-proto:
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --go_out=logtostderr=true:./internal/generated/api \
	        --twirp_out=./internal/generated/api \
	        $(shell ls api/protobuf-spec/v1 | grep .proto | awk '{print "v1/"$$1}')

# https://github.com/grpc-ecosystem/grpc-gateway/issues/837
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=tablo:./api/openapi-spec \
	        desc.proto
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --swagger_out=logtostderr=true,allow_merge=true,merge_file_name=v1:./api/openapi-spec \
	        $(shell ls api/protobuf-spec/v1 | grep .proto | awk '{print "v1/"$$1}')
	@swagger mixin \
	        -o api/openapi-spec/swagger.json \
	        api/openapi-spec/tablo.swagger.json \
	        api/openapi-spec/v1.swagger.json &>/dev/null
	@rm api/openapi-spec/*.swagger.json

.PHONY: generate-go
generate-go:
	@go generate $(PACKAGES)

.PHONY: refresh
refresh: deps-update deps generate format test-with-coverage


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-smoke
test-smoke:
	@go test -run=^smoke -tags=integration -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)


.PHONY: build
build:
	@go build -o bin/$(BIN) -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)" .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh
