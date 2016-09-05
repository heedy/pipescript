The `alltrue` transform returns `true` if and only if all of the datapoints seen thus far in the stream have been true:

Given the following data:
```
true,true,true,false,true,true
```

`alltrue` will return:
```
true,true,true,false,false,false
```

### Why It's Useful

Oftentimes you might want to check something in a `while`, or in a `map`. For example,
the following transform will return true for each day where the entire 24 hours was spent at home:

```
distance(<latitude>,<longitude>) < 40 | while(day==next:day,alltrue)
```

The above transform will run a while loop while the current datapoint is part of the same day as the next datapoint, and check whether all location datapoints that day were within 40 meters of your chosen latitude and longitude.
