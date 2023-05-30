---
code: null
id: MDU6SXNzdWU1MDk4OTgzOTQ=
databaseId: 509898394
number: 9
url: https://github.com/octopot/tablo/issues/9
title: "update deps and improve makefile"
issueType: null
labels:
  - "type:task"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2019-10-21T11:20:51Z
updatedAt: 2019-10-22T20:01:25Z
lastEditedAt: 2019-10-22T20:01:12Z
closedAt: 2019-10-22T20:01:25Z
---

# update deps and improve makefile

- [x] goreleaser and godownloader
- [x] try to move them up (tools)
- [x] test-check
```
.PHONY: test-check
test-check:
	@go test -run=^hack $(PACKAGES)
```
- [x] test-full
```
.PHONY: test-full
test-full: test-with-coverage test-integration

.PHONY: test-integration
test-integration:
	@go test -tags=integration -run=^Example ./test/... 2> /dev/null
	@go test -tags=integration \
	          -cover \
	          -covermode count \
	          -coverpkg $(shell go list -m)/... \
	          -coverprofile integro.out \
	          -run=^TestService \
	          ./test/api
```
- [x] implement test-smoke
- [x] update env
```
.PHONY: env
env:
	@echo "GO111MODULE: $(shell go env GO111MODULE)"
	@echo "GOFLAGS:     $(shell go env GOFLAGS)"
	@echo "GOPRIVATE:   $(shell go env GOPRIVATE)"
	@echo "GONOPROXY:   $(shell go env GONOPROXY)"
	@echo "GONOSUMDB:   $(shell go env GONOSUMDB)"
	@echo "GOPROXY:     $(shell go env GOPROXY)"
	@echo "GOSUMDB:     $(shell go env GOSUMDB)"
```
- [x] update
```
.PHONY: update
update:
	@GOPRIVATE=$(GOPRIVATE) go get -mod= -u ./...
```
