# Pipescript



This is the transform and query engine used in ConnectorDB. Pipescript includes methods to perform transforms on iterators of data (streams),
and also includes a structured query format for generation of datasets.

**Warning: This is a work in progress - we are performing upgrades to the query engine as we move it over to this repository**

## The Issue

While SQL is the standard query language of databases (and pipescript might be extended to include sql-like SELECT syntax in the future),
when performing analysis on multiple independent streams of data with the intention of use towards machine learning or plotting,
it helps to have an extremely simple pipeline of data preparation, and special queries for time series interpolation.

## Examples

### Transforms
Suppose you have a time series of the number of steps taken. To sum over all steps:
```
sum | if last
```

The sum transform sums all the datapoints - it would still return each datapoint with the intermediate sums.
We therefore use Pipescript's if statement to filter all datapoints except the last one.

Ultimately, each "transform" is separated by a pipe (`|` or `:`), and the result of the entire transform pipeline is returned.

### Queries

Combining multiple time series into tabular datasets is critically important for ML applications. The Pipescript query engine allows
you to do just that!
