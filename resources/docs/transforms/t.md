Every datapoint in a stream has a timestamp. This timestamp is hidden when performing operations in PipeScript, since all operations are performed on a datapoint's `d` (data) field. To permit processing based upon timestamps in PipeScript, the `t` transform exposes the datapoint's timestamp.

Remember that raw datapoints are in the form:

```json
[
  {
    "t": 123456.23,
    "d": 4,
    "dt": 0
  }
]
```

When performing transforms, `$==4` will return `true`, since we operate on the "d" (data) field. But the timestamp is `123456.23`, so the transform `t` will result in the following datapoint:

```json
[
  {
    "t": 123456.23,
    "d": 123456.23,
    "dt": 0
  }
]
```
