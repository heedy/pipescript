The changed transform returns true if the current datapoint's data is different from the previous datapoint.

Given the following data:

```json
[1, 2, 3, 3, 4]
```

running the `changed` transform on it will result in:

```json
(true, true, true, false, true)
```

### Usage

The `changed` transform is useful whenever you don't care about the specific datapoints,
but when they change with respect to a certain metric.

Suppose you have a stream of activities gathered by a phone:

```json
["walking", "walking", "still", "still", "still", "walking", "walking"]
```

You usually care about the transitions between activities:

```
if changed
```

gives:

```json
["walking", "still", "running"]
```

Remember that each datapoint comes with a timestamp, so the length of each activity can be extracted by looking at the timestamp differences.
