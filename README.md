[![ReportCard](https://goreportcard.com/badge/heedy/pipescript)](https://goreportcard.com/report/heedy/pipescript)
[![GoDoc](https://godoc.org/github.com/heedy/pipescript?status.svg)](http://godoc.org/github.com/heedy/pipescript)

# PipeScript

Pipescript can be used as a standalone executable (`pipes`) for quick data filtering and analysis. It can also be plugged into databases
as a query and filtering language (as is done in heedy's timeseries plugin).

## The Issue

SQL is the standard query language for databases. Unfortunately, SQL is not designed for usage with large time series, where the dataset might be enormous, and several disparate series need to be combined into datasets for use in ML applications.

First off, when operating on time series, we can make many assumptions, such as timestamps being ordered. This allows us to implement many operations on data as streaming transformations, rather than large aggregations. PipeScript can perform analysis on enormous files without much memory use, as all calculations are done (almost) locally.

Second, the natural form for time series is with each datapoint from separate sensors/data streams to have its own timestamp. Your phone reports usage data independently of your laptop. Nevertheless, the real power of our datasets comes when we can combine many different streams, all with differing time stamps into a coherent (and tabular) whole. This requires more manipulation of the data than a simple JOIN, and can be done very efficiently on streaming ordered time series.

PipeScript is a very basic transform language and interpolation/dataset generation machinery which works entirely on streaming (arbitrarily sized) data. It is the main query engine used for heedy timeseries.

## Tutorial

A tutorial will be available with heedy documentation once it is released.

<!--A general PipeScript tutorial [can be found here](https://heedy.io/pipescript/docs/basics.html).
Instructions for embedding and extending PipeScript with your own transforms and interpolators can be found [here](https://heedy.io/pipescript/docs/embedding.html).-->

## Standalone

### Compiling

```
go get github.com/heedy/pipescript
go test github.com/heedy/pipescript/...
go install github.com/heedy/pipescript/pipes
```

### Usage

The standalone can be used to run PipeScript queries on datapoint, json, and csv formatted data. The following will count the number of datapoints in a csv file:

```
pipes run -i myfile.csv -ifmt csv --notimestamp "count"
```

#### SQL

If you want to perform SQL queries on files, please look at [q](https://github.com/harelba/q).
Unfortunately, advanced SQL operations require keeping the entire dataset in memory - something PipeScript tries to avoid.
