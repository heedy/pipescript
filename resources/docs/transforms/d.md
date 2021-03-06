The identity transform is a placeholder for the "current datapoint". It returns whatever is passed from the timeseries.

Suppose your timeseries has the following data:

```json
[44, 18, 20, -35, 20.23]
```

The d transform will simply return your data unchanged:

```json
[44, 18, 20, -35, 20.23]
```

### Usage

The identity transform is perhaps the most used transform in all of pipescript.
It is frequently used in comparisons and filters

For example, the transform `d > 20` is a comparison - it checks whether the current datapoint, represented by the identity is greater than 20. The result of this transform would be:

```json
[true, false, false, false, true]
```

This is frequently used in filters: `where d > 20` would return

```json
[44, 20.23]
```

### Objects

The d transform accepts an optional argument. Sometimes, a datapoint isn't just your data - it can be an object:

```json
[
  { "hi": "hello", "foo": "bar" },
  { "hi": "world", "foo": "baz" }
]
```

Running this transform:

```javascript
d("hi");
```

gives us:

```json
["hello", "world"]
```

### Peeking

If given an integer index, the \d transform performs a _peek_ operation - instead of returning the current datapoint, it returns the one at a relative index. For example, you can find the differences between the data of each successive datapoint with the following:

```
d[1] - d
```

```json
[-26, 2, -55, 55.23]
```

Currently, you can only peek forward in the timeseries (i.e. look at future datapoints).
