`contains` permits you to check if a datapoint with a string data value contains the given substring:

```json
["Hello World!","hello world","hi there"]
```

Running the transform `contains("World")` will give:
```json
[true,false,false]
```
