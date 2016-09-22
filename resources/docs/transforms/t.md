Every datapoint in a stream has a timestamp. This timestamp is hidden when performing operations in PipeScript, since all operations are performed on a datapoint's `data` field. To permit processing based upon timestamps in PipeScript, the `t` transform exposes the datapoint's timestamp.

Remember that raw datapoints are in the form:

```json
[{
  "t": 123456.23,
  "d": 4
}]
```

When performing transforms, `$==4` will return `true`, since we operate on the "d" (data) field. But `t==123456.23`, which corresponds
to the following raw datapoint:
```json
[{
  "t": 123456.23,
  "d": 123456.23
}]
```
