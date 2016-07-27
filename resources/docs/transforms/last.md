This is true on the last datapoint of a sequence.

It is *very* common in pipescript to end a transform with `if last` to only return the final datapoint, which contains the desired result of computation, without returning intermediate values.

```
sum | if last
```

will sum up all datapoints, and return a single datapoint with the full sum. If `if last` were not there, all intermediate values of the sum would be returned.
