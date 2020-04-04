The `set` transform allows to add/modify values in an object-formatted datapoint.

```json
[
  {},
  {"foo": "baz"},
  {"a": 1,"b":2}
]
```
With the above data, and the transform `set("foo","bar")`, we get:

```json
[
  {"foo":"bar"},
  {"foo": "bar"},
  {"a": 1,"b":2,"foo":"bar"}
]
```

Remember that PipeScript has native support for json-formatted data. You can directly set objects with transforms like the following:
```
{"foo": "bar", "a": $("a")}
```
