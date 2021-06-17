Oftentimes, you don't really want all of the raw data from your stream. You are usually interested in timed aggregations or other such things. This is the main use case of the `while` transform.

This transform performs a while loop on its second argument while its first argument is true. When the first argument becomes false, it returns the resulting datapoint, and begins the next loop.

This allows summing up datapoints over specified time periods. For example, `day==d[1]:day` is true if the current datapoint and next datapoint in the stream have timestamps from the same day.

This allows you to write a transform performing an aggregation per day:

```
while(day==d[1]:day, sum)
```

The above transform loops through datapoints while they come from the same day, and sums their values. Once the next datapoint is a different day than the current one, it ends the while loop, and returns the sum, giving the sum of all datapoints that day. It then starts a loop for the next day.

The transform can also be used for smoothing. Suppose you want to smooth your data every three datapoints:

```
while(i%3!=0, mean)
```

This returns the mean of every consecutive three datapoints, making it easy to do a basic smoothing procedure.

## Advanced Usage

The transform can also be used to implement error bars:

```
while(day==d[1]:day, {"max": max, "min": min, "mean": mean})
```

This transform returns the mean, max and min datapoint for the day all at once, allowing to plot with error bars.
