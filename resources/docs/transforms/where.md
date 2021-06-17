The where transform is used for filtering data.

Suppose the data portion of your dataset is as follows:

```
1,2,3,4,5,6
```

The transform:

```
where d >= 5
```

will leave you with:

```
5,6
```

Note that while convention is to use where without parentheses (bash style), `where` is a normal pipescript transform, and can be used as a function: `where(d >= 5)`.

## and/or

PipeScript supports python-like and/or statements to build up a boolean:

```
where d > 5 and d < 20
```

The above will only pass through datapoints between 5 and 20. Just like in other languages, you can use parentheses to force an order of operations.

PipeScript also has a built in negation:

```
where not d > 5
```

Combining and/or with not allows building up arbitrary conditions.
