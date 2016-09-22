The domain transform is run on URLs, and returns the domain.

For example:
```json
[
  "https://golang.org/pkg/net/url/#URL.EscapedPath",
  "https://connectordb.github.io",
  "https://github.com/connectordb/connectordb"
]
```
gives:
```json
[
  "golang.org",
  "connectordb.github.io",
  "github.com"
]
```
