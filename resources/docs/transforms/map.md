The map transform is an example of a transform which hijacks its second argument. Please note that it is only lightly related to the standard map function in most programming languages. It splits datapoints along its first argument, and runs independent scripts on each value:

```json
[
{
    "t": 1,
    "d": {
          "steps": 14,
          "activity": "walking"
        }
},
{
    "t": 2,
    "d": {
          "steps": 10,
          "activity": "running"
        }
},
{
    "t": 3,
    "d": {
          "steps": 12,
          "activity": "walking"
        }
},{
    "t": 4,
    "d": {
          "steps": 5,
          "activity": "running"
        }
}]
```


Mapping by activity:
```
map($("activity"),$("steps"):sum)
```

```json
[{
    "t": 4,
    "d": {
          "walking": 26,
          "running": 15
        }
}]
```

The map transform split the dataset by its first argument (`$("activity")`), and performed the second transform (`$("steps"):sum`) on each subset independently. This allows directly returning the sum of datapoints where you were walking and running.

The map transform is frequently used in conjunction with the `reduce` transform for quick analysis.

Very common use for the map transform is splitting by time periods. For example, to get the total number of steps taken per day (suppose steps is a number stream, where the data value is number of steps taken)

```
map(day,sum)
```

A histogram of weekdays can also be generated:

```
map(weekday,sum)
```

A more advanced way would be to average weekday values:

```
map(weekday, imap(day,sum):reduce(mean))
```

Note the imap transform within the map. It functions in exactly the same way as map, but returns all intermediate values (which allows it to be embedded in map). The above transform first splits by weekday, then further splits by day. It sums up the steps taken per day and finds their mean. That leaves with daily averages per weekday.
