The regex transform checks if the data string matches the given regex.

For example, given a regex to check for valid usernames: `regex('^[a-z0-9_-]{3,16}$')`, we get:

```json
[
  "Hello World!",
  "valid_username",
  "1"
]
```

```json
[false,true,false]
```
