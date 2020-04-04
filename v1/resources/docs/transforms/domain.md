The domain transform is run on URLs, and returns the domain.

For example:

```json
[
  "https://golang.org/pkg/net/url/#URL.EscapedPath",
  "https://heedy.io",
  "https://github.com/heedy/heedy"
]
```

gives:

```json
["golang.org", "heedy.io", "github.com"]
```
