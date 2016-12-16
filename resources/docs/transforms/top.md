The `top` transform allows you to filter smaller elements of an object:

```json
[{
    "a": 1,
    "b": 2,
    "c": 3,
    "d": 4
}]
```

Running `top(2)` on the above will return:

```json
[{
    "c": 3,
    "d": 4
}]
```

This transform is extremely useful for limiting results - for example, you don't usually
care about every website you ever visted, but rather the top 20 websites. The following
transform will return the top 20 websites visited by url:

```
map(domain,count) | top(20)
```

You might also want to get the top 20 elements of a non-object (for example, the top
20 scores you ever got in a game). This can
be done by converting it into an object first:

```
map(count,$) | top(20)
```

Note that the `map` transform can only hold a limited number of datpoints. it is therefore
useful to pre-filter datapoints that you know can't be in the top 20:

```
if $ > 100 | map(count,$) | top(20)
```