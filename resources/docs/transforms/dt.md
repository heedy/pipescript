Every datapoint in a stream has a duration. This value is hidden when performing operations in PipeScript, since all operations are performed on a datapoint's `d` (data) field. To permit processing based upon duration in PipeScript, the `dt` transform exposes the datapoint's duration.

Remember that raw datapoints are in the form:

```json
[
  {
    "t": 123456.23,
    "d": 4,
    "dt": 80.0
  }
]
```

When performing transforms, `$==4` will return `true`, since we operate on the "d" (data) field. But the duration is `80.0`, so the transform `dt` will result in the following datapoint:

```json
[
  {
    "t": 123456.23,
    "d": 80.0,
    "dt": 80.0
  }
]
```
