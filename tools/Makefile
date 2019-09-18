BINPATH     = $(shell dirname $(PWD))/bin
GO111MODULE = on
GOFLAGS     = -mod=
GOPROXY     = https://proxy.golang.org
SHELL       = /bin/bash -euo pipefail

export PATH := $(BINPATH):$(PATH)

.DEFAULT_GOAL = install

.PHONY: env
env:
	@echo "BINPATH:     $(BINPATH)"
	@echo "GO111MODULE: $(GO111MODULE)"
	@echo "GOFLAGS:     $(GOFLAGS)"
	@echo "PATH:        $(PATH)"
	@echo "SHELL:       $(SHELL)"


.PHONY: brew
brew:
	@brew list protobuf &>/dev/null || brew install protobuf

.PHONY: build
build:
	@GOFLAGS='-mod=vendor' egg --build
	@mv bin/* $(BINPATH) && rm -rf bin

.PHONY: deps
deps:
	@go mod tidy && go mod vendor && go mod verify

.PHONY: egg
egg:
	@go get -d -u github.com/kamilsk/egg
	@go build -o $(BINPATH)/egg github.com/kamilsk/egg

.PHONY: install
install: brew egg deps build

.PHONY: update
update:
	@go get -tags=tools -u
