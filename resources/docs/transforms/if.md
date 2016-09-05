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
