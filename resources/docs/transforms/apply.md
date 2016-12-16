The `apply` transform allows you to apply a script element-wise to a JSON object. For example:

```json 
[{
    "a": 5,
    "b": -2
}]
```

```
apply($+2)
```

will return:

```json 
[{
    "a": 7,
    "b": 0
}]
```