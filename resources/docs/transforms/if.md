The if transform is used for filtering data. If you are looking for an if statement more akin to other languages (not a filter), you can use the `ifelse` transform.

Suppose the data portion of your dataset is as follows:
```
1,2,3,4,5,6
```
The transform:
```
if $ >= 5
```
will leave you with:
```
5,6
```

Note that while convention is to use if without parentheses (bash style), `if` is a normal pipescript transform, and can be used as a function: `if($ >= 5)`.

## and/or

PipeScript supports python-like and/or statements to build up a boolean:

```
if $ > 5 and $ < 20
```

The above will only pass through datapoints between 5 and 20. Just like in other languages, you can use parentheses to force an order of operations.

PipeScript also has a built in negation:

```
if not $ > 5
```

Combining and/or with not allows building up arbitrary conditions.
