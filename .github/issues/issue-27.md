---
code: null
id: MDU6SXNzdWU1MjA2MzQ5ODg=
databaseId: 520634988
number: 27
url: https://github.com/octopot/tablo/issues/27
title: "improve storage engine layer"
issueType: null
labels:
  - "type:refactoring"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2019-11-10T18:03:41Z
updatedAt: 2019-11-10T18:03:41Z
lastEditedAt: null
closedAt: null
---

# improve storage engine layer

```
		// TODO:debt
		//  - health check db.PingContext
		//  - better connection configuration through dsn
		//    - db.SetConnMaxLifetime
		//    - db.SetMaxIdleConns
		//    - db.SetMaxOpenConns
		//  - unset connection configuration from dsn
		//    - schema.Query().Del()
		//    - schema.RawQuery
```
