This is true when the datapoint is first in a sequence. It is useful mainly for filtering:

```
filter first or last
```

will return the first and last datapoint in your stream:

```
1,2,3,4,5
```

```
1,5
```

### Usage

This transform could be finding your wakeup time, based upon the first time you turn on your phone screen in the morning:

```
filter dayhour > 4 | while(day == next:day, first) | t
```

The above transform removes the datapoints taken from midnight to 4am (to filter out long nights), and then returns the first datapoint of each day, finaally returning only the timestamp.
