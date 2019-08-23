BIN         = $(shell basename $(shell go list -m))
GO111MODULE = on
GOFLAGS     = -mod=vendor
MODULE      = $(shell go list -m)
PACKAGES    = $(shell go list ./...)
PATHS       = $(shell go list ./... | sed -e "s|$(shell go list -m)/||g")
SHELL       = /bin/bash -euo pipefail
TIMEOUT     = 1s


.DEFAULT_GOAL = test-with-coverage


.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: update
update:
	@go get -mod= -u


.PHONY: format
format:
	@goimports -local $(dirname $(MODULE)) -ungroup -w $(PATHS)

.PHONY: generate
generate: go-generate proto-generate

.PHONY: go-generate
go-generate:
	@go generate ./...

.PHONY: proto-generate
proto-generate:
	@protoc --proto_path=./api/protobuf-spec \
	        --twirp_out=./internal/generated/api/v1 \
	        --go_out=./internal/generated/api/v1 \
	        service.proto

.PHONY: refresh
refresh: generate format


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
	@go build -o bin/$(BIN) .

.PHONY: dist
dist:
	@godownloader .goreleaser.yml > .github/install.sh

.PHONY: install
install:
	@go build -o $(GOPATH)/bin/$(BIN) .
