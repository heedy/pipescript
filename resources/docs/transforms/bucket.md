The bucket transform allows you to put numbers into buckets of custom size.

For example, given this data:

```
2,16,84,-5,1
```
We can choose buckets of size `10` (starting from `0` by default)
```
bucket(10)
```
to get:
```json
"[0,10)", "[10,20)", "[80,90)","[-10,0)","[0,10)"
```

The bucket transform uses [interval notation][1]. To process the range in code, you can manually change the last character from `)` to `]`, and parse the result as a json array.

[1]: https://en.wikipedia.org/wiki/Interval_(mathematics)

### Histograms

Using the `bucket` transform in conjunction with the `map` transform, it is easy to generate histograms:

```
map(bucket(10),count)
```
Running this transform on the above data will give:
```json
[{
  "[-10,0)": 1,
  "[0,10)": 2,
  "[10,20)": 1,
  "[80,90)": 1
}]
```
