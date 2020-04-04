The `i` transform gives the index in the timeseries array, starting with 0

```
"data","data","data","data",...
```

gives:

```
0,1,2,3
```

When used within a pipe, it gives the index of datapoint _within the pipe_. The following transform filters half the data, and returns the indices at the second pipe location

```
filter i%2==0 | i
```

```
0,1
```
