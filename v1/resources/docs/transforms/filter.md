`filter` allows you to remove values from an object based upon a condition. For example:

```json
[{
  "a": 45,
  "b": 23,
  "c": -3
}]
```

Running the transform `filter($<0)` on the above datapoint, will give an output:

```json
[{
  "a": 45,
  "b": 23
}]
```

The transform "filtered" out all values that are less than 0.

## Examples

The main use case of the `filter` transform is in combination with the `map` transform, when `map` returns too many values.
For example, if looking at browsing history, you might have visited thousands of websites, but only be interested in the ones you visited more than 30 times:

```
map(domain,count) | filter($<=30)
```
