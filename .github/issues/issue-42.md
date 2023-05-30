---
code: null
id: MDU6SXNzdWU1Mjg5NjI5Njk=
databaseId: 528962969
number: 42
url: https://github.com/octopot/tablo/issues/42
title: "problem with ssl cert"
issueType: null
labels:
  - "type:bug"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: CLOSED
stateReason: COMPLETED
duplicateOf: null
createdAt: 2019-11-26T20:46:28Z
updatedAt: 2019-11-27T20:45:28Z
lastEditedAt: null
closedAt: 2019-11-27T20:45:27Z
---

# problem with ssl cert

```
2019/11/26 20:41:51 [emerg] 24#24: cannot load certificate "/etc/nginx/ssl/xip.io.crt": BIO_new_file() failed (SSL: error:02001002:system library:fopen:No such file or directory:fopen('/etc/nginx/ssl/xip.io.crt','r') error:2006D080:BIO routines:BIO_new_file:no such file)
nginx: [emerg] cannot load certificate "/etc/nginx/ssl/xip.io.crt": BIO_new_file() failed (SSL: error:02001002:system library:fopen:No such file or directory:fopen('/etc/nginx/ssl/xip.io.crt','r') error:2006D080:BIO routines:BIO_new_file:no such file)
nginx: configuration file /etc/nginx/nginx.conf test failed
```
