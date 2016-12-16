The `percent` transform allows you to get the percentage contribution of each element in an object. Suppose you have the following:

```json
[{
    "a": 2,
    "b": 5,
    "c": 1,
    "d": 2
}]
```

The `percent` transform will return:

```json
[{
    "a": .2,
    "b": .5,
    "c": .1,
    "d": .2
}]
```

This allows you to quickly normalize a JSON object's values.