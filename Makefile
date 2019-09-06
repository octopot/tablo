BIN         = $(shell basename $(shell go list -m))
BINPATH     = $(PWD)/bin
COMMIT      = $(shell git rev-parse --verify HEAD)
DATE        = $(shell date +%Y-%m-%dT%T%Z)
GO111MODULE = on
GOFLAGS     = -mod=vendor
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
	@echo "GO111MODULE: $(GO111MODULE)"
	@echo "GOFLAGS:     $(GOFLAGS)"
	@echo "MODULE:      $(MODULE)"
	@echo "PACKAGES:    $(PACKAGES)"
	@echo "PATH:        $(PATH)"
	@echo "PATHS:       $(PATHS)"
	@echo "SHELL:       $(SHELL)"
	@echo "TIMEOUT:     $(TIMEOUT)"


.PHONY: deps
deps: tools
	@go mod tidy && go mod vendor && go mod verify

.PHONY: format
format:
	@goimports -local $(dir $(shell go list -m)) -ungroup -w $(PATHS)

.PHONY: generate
generate: generate-go generate-proto

.PHONY: generate-go
generate-go:
	@go generate $(PACKAGES)

.PHONY: generate-proto
generate-proto:
	@protoc --proto_path=./api/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec \
	        --proto_path=./third_party/protobuf-spec/googleapis \
	        --go_out=logtostderr=true:./internal/generated/api \
	        --twirp_out=./internal/generated/api \
	        v1/service.proto

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
	        v1/service.proto
	@swagger mixin \
	        -o api/openapi-spec/swagger.json \
	        api/openapi-spec/tablo.swagger.json \
	        api/openapi-spec/v1.swagger.json
	@rm api/openapi-spec/*.swagger.json

.PHONY: tools
tools:
	@cd tools && make

.PHONY: update
update:
	@go get -mod= -u

.PHONY: refresh
refresh: update deps generate format test-with-coverage


.PHONY: test
test:
	@go test -race -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-with-coverage
test-with-coverage:
	@go test -cover -timeout $(TIMEOUT) $(PACKAGES) | column -t | sort -r

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@go test -cover -covermode count -coverprofile c.out -timeout $(TIMEOUT) $(PACKAGES)

.PHONY: test-smoke
test-smoke:
	@echo not implemented yet


.PHONY: build
build:
	@go build -o bin/$(BIN) -ldflags "-s -w -X main.commit=$(COMMIT) -X main.date=$(DATE)" .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh
