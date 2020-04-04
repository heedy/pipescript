The `anytrue` transform returns `true` if any of the datapoints in the stream were true.

Given the following data:

```
false,false,false,true,false,false
```

`anytrue` will return:

```
true
```

### Why It's Useful

Oftentimes you might want to check something in a `while`, or in a `map`. A great example
would be to check which days you went to the gym:

```
distance(<latitude>,<longitude>) < 40 | while(day==next:day,anytrue)
```

The above transform will run a while loop while the current datapoint is part of the same day as the next datapoint, and check whether any of the GPS coordinates were within 40 meters of your gym.
