`append` is a transform made for strings - it appends all of the data it has seen so far into one large string:

Suppose your stream is:
```json
["hello ","world",24,true]
```

Then running the transform `append` on your data will give you:
```json
["hello ","hello world","hello world24","hello world24true"]
```
