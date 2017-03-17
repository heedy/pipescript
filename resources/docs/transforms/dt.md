The `dt` transform allows you to quickly find the time difference between the previous and this datapoint:

```json
[
    {"t": 4, "d": 4},
	{"t": 20, "d": 5},
	{"t": 50, "d": 6}
]
```

```
dt
```

```json
[
    {"t": 4, 0)},
	{"t": 20, "d": 16},
	{"t": 50, "d":30}
]
```