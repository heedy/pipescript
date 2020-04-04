The `new` transform allows you to check if the given datapoints were seen before in the stream.

The main use case here is `if new`, which returns all the datapoints with unique data.

Suppose your data is:
```json
["foo","foo","bar","foo","bar","baz"]
```

the transform `new` will return:
```json
[true,false,true,false,false,true]
```
