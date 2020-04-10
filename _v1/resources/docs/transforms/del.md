The `del` transform allows to delete values in an object-formatted datapoint.

```json
[
  {},
  {"foo": "baz"},
  {"a": 1,"b":2,"foo":7}
]
```
With the above data, and the transform `del("foo")`, we get:

```json
[
  {},
  {},
  {"a": 1,"b":2}
]
```

Remember that PipeScript has native support for json-like data. You can directly set objects with transforms like the following:
```json
{"foo": "bar", "a": $("a")}
```
