`remember` allows you to save a chosen datapoint. Whenever the argument of `remember` is `true`, it saves the current datapoint, and keeps repeating it while the argument is `false`.

An example will explain this very nicely. If your data is:
```
20,40,-50,20,-10,3,-9,40,50
```

then the transform `remember($ < 0)` will return:
```
20,20,-50,-50,-10,-10,-9,-9,-9
```

The reason the `20` is repeated at the beginning, despite it being positive, is because `remember` is always initialized with your first datapoint.
