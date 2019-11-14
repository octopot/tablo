BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=
GOPROXY     = https://proxy.golang.org,https://gocenter.io,direct
SHELL       = /bin/bash -euo pipefail

export PATH := $(BINPATH):$(PATH)

.DEFAULT_GOAL = install

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(shell go env GOFLAGS)"
	@echo "GOPRIVATE:   $(shell go env GOPRIVATE)"
	@echo "GOPROXY:     $(shell go env GOPROXY)"
	@echo "GONOPROXY:   $(shell go env GONOPROXY)"
	@echo "GOSUMDB:     $(shell go env GOSUMDB)"
	@echo "GONOSUMDB:   $(shell go env GONOSUMDB)"
	@echo "PATH:        $(PATH)"
	@echo "SHELL:       $(SHELL)"


.PHONY: brew
brew:
	@brew bundle

.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: deps-update
deps-update:
	@go get -u all

.PHONY: egg
egg:
	@go get -d -u github.com/kamilsk/egg
	@go build -o $(BINPATH)/egg github.com/kamilsk/egg


.PHONY: build
build:
	@GOFLAGS='-mod=vendor' egg --build
	@mv bin/* $(BINPATH) && rm -rf bin

.PHONY: install
install: brew egg deps build
