The `next` transform returns the datapoint *next* in the stream.

Suppose your data is`1,2,3,3,4`. The transform `$ == next` will return `false,false,true,false,false`.


The following is a verbose expansion of what is going on:
```
1==2  false
2==3  false
3==3  true
3==4  false
4==null false
```

At the end, the next transform returns null data, since its stream has ended.
