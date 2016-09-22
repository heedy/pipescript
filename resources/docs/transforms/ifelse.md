`ifelse` is the same type of conditional statement that you might be used to in other programming languages.

Pipescript's `if` statement is a filter, because conditionals are much less common than filters.

```
ifelse($ > 5, $-5)
```

The above will take all datapoints with data greater than 5, and decrease their value by 5. There is also an optional `else`:

```
ifelse($ > 5, $-3,$+2)
```

The above will decrease datapoints > 5 by 3, and increase all others by 2.

This permits performing conditional computation on streams.
