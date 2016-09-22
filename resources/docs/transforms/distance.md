The `distance` transform computes the distance in meters from the current datapoint to its argument coordinates.

The datapoint is assumed to have `latitude` and `longitude` fields in decimal coordinates. It returns the distance in meters computed using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula).

```json
[{
  "latitude": 40.4277304,
  "longitude": -86.9170587
}]
```
Given the above stream, we find its distance to the chosen coordinate: `distance(40.426841,-86.9165106)`:
```json
[109.238]
```

The two coordinates above are about 109 meters apart.
