[![Build Status](https://travis-ci.org/connectordb/pipescript.svg)](https://travis-ci.org/connectordb/pipescript)[![Coverage Status](https://coveralls.io/repos/connectordb/pipescript/badge.svg?branch=master&service=github)](https://coveralls.io/github/connectordb/pipescript?branch=master)[![GoDoc](https://godoc.org/github.com/connectordb/pipescript?status.svg)](http://godoc.org/github.com/connectordb/pipescript)

**Warning: This is a work in progress - we are performing upgrades to the query engine as we move it over to this repository**

# PipeScript

When writing ConnectorDB, we noticed that our transform and query engine was actually quite powerful, and would be extremely useful as a
plugin and as a standalone data analysis tool. We therefore decided to extract our transform language, called 'pipescript', and our
dataset and query generation capabilities into this tool.

Pipescript can be used as a standalone executable for use in quick data filtering and analysis. It can also be plugged into databases
as a query and filtering language.

## Compiling

```
go get github.com/connectordb/pipescript
cd "$(GOPATH)
go test ./...
cd main
go build
```

## Using as Query Engine

This is the transform and query engine used in ConnectorDB. PipeScript includes methods to perform transforms on iterators of data (streams),
and also includes a structured query format for generation of datasets. More to come on this topic.

## Standalone

Pipescript can be compiled into a standalone executable `pipes`. This enables you to perform filtering and data analysis on json and csv data,
as well as any dataset type delimited by newline.

For example, summing the values in the `steps` column of a CSV file:
```
pipes -i data.csv "sum($.steps)"
```

#### SQL
If you want to perform SQL queries on files, please look at [q](https://github.com/harelba/q).
Unfortunately, advanced SQL operations require keeping the entire dataset in memory.

PipeScript is a streaming language, which makes `pipes` able to handle gigabytes of data, but unfortunately makes us unable to implement SQL. While a small subset of SQL might be implemented at some point, it is not a priority. Our priority is adding useful transforms into pipescript.

### PipeScript commands

If you link the `pipes` executable to name of a command, it will perform the command. So `ln -s sum pipes` allows:

```
cat data.csv | sum "$.steps"
```
will perform the same as above.



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

To view currently implemented transform documentation, please visit the [Transform List](https://connectordb.com/www/docs/transforms.html)
