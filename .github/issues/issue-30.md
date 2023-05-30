---
code: null
id: MDU6SXNzdWU1MjMwOTUyODc=
databaseId: 523095287
number: 30
url: https://github.com/octopot/tablo/issues/30
title: "tools: Brewfile"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2019-11-14T20:46:07Z
updatedAt: 2019-11-14T20:53:43Z
lastEditedAt: null
closedAt: 2019-11-14T20:53:43Z
---

# tools: Brewfile

instead of

```
.PHONY: brew
brew:
	@brew list protobuf &>/dev/null || brew install protobuf
```
