// Code generated by go-bindata.
// sources:
// docs/transforms/$.md
// docs/transforms/alltrue.md
// docs/transforms/anytrue.md
// docs/transforms/bucket.md
// docs/transforms/changed.md
// docs/transforms/contains.md
// docs/transforms/count.md
// docs/transforms/distance.md
// docs/transforms/dt.md
// docs/transforms/filter.md
// docs/transforms/first.md
// docs/transforms/i.md
// docs/transforms/last.md
// docs/transforms/map.md
// docs/transforms/mean.md
// docs/transforms/reduce.md
// docs/transforms/regex.md
// docs/transforms/sum.md
// docs/transforms/t.md
// docs/transforms/tshift.md
// docs/transforms/wc.md
// docs/transforms/while.md
// DO NOT EDIT!

package resources

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _docsTransformsMd = []byte(`The identity transform is a placeholder for the "current datapoint". It returns whatever is passed from the timeseries.

Suppose your timeseries has the following data:

`+"`"+``+"`"+``+"`"+`json
(44, 18, 20, -35, 20.23)
`+"`"+``+"`"+``+"`"+`

The \$ transform will simply return your data unchanged:

`+"`"+``+"`"+``+"`"+`json
(44, 18, 20, -35, 20.23)
`+"`"+``+"`"+``+"`"+`

### Usage

The identity transform is perhaps the most used transform in all of pipescript.
It is frequently used in comparisons and filters

For example, the transform `+"`"+`$ > 20`+"`"+` is a comparison - it checks whether the current datapoint, represented by the identity is greater than 20. The result of this transform would be:

`+"`"+``+"`"+``+"`"+`json
(true, false, false, false, true)
`+"`"+``+"`"+``+"`"+`

This is frequently used in filters: `+"`"+`filter $ > 20`+"`"+` would return

`+"`"+``+"`"+``+"`"+`json
(44, 20.23)
`+"`"+``+"`"+``+"`"+`

### Objects

The `+"`"+`$`+"`"+` transform accepts an optional argument. Sometimes, a datapoint isn't just your data - it can be an object:

`+"`"+``+"`"+``+"`"+`json
({ "hi": "hello", "foo": "bar" }, { "hi": "world", "foo": "baz" })
`+"`"+``+"`"+``+"`"+`

Running this transform:

`+"`"+``+"`"+``+"`"+`javascript
$("hi");
`+"`"+``+"`"+``+"`"+`

gives us:

`+"`"+``+"`"+``+"`"+`json
("hello", "world")
`+"`"+``+"`"+``+"`"+`

### Peeking

If given an integer index, the `+"`"+`$`+"`"+` transform performs a _peek_ operation - instead of returning the current datapoint, it returns the one at a relative index. For example, you can find the differences between the data of each successive datapoint with the following: `+"`"+`$[1] - $`+"`"+`, which would give

`+"`"+``+"`"+``+"`"+`
-26,2,-55,55.23
`+"`"+``+"`"+``+"`"+`

Currently, you can only peek forward in the timeseries (i.e. look at future datapoints).
`)

func docsTransformsMdBytes() ([]byte, error) {
	return _docsTransformsMd, nil
}

func docsTransformsMd() (*asset, error) {
	bytes, err := docsTransformsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/$.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsAlltrueMd = []byte(`The `+"`"+`alltrue`+"`"+` transform returns `+"`"+`true`+"`"+` if and only if all of the datapoints in the stream are true:

Given the following data:

`+"`"+``+"`"+``+"`"+`
true,true,true,false,true,true
`+"`"+``+"`"+``+"`"+`

`+"`"+`alltrue`+"`"+` will return:

`+"`"+``+"`"+``+"`"+`
false
`+"`"+``+"`"+``+"`"+`

### Why It's Useful

Oftentimes you might want to check something in a `+"`"+`while`+"`"+`, or in a `+"`"+`map`+"`"+`. For example,
the following transform will return true for each day where the entire 24 hours was spent at home:

`+"`"+``+"`"+``+"`"+`
distance(<latitude>,<longitude>) < 40 | while(day==next:day,alltrue)
`+"`"+``+"`"+``+"`"+`

The above transform will run a while loop while the current datapoint is part of the same day as the next datapoint, and check whether all location datapoints that day were within 40 meters of your chosen latitude and longitude.
`)

func docsTransformsAlltrueMdBytes() ([]byte, error) {
	return _docsTransformsAlltrueMd, nil
}

func docsTransformsAlltrueMd() (*asset, error) {
	bytes, err := docsTransformsAlltrueMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/alltrue.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsAnytrueMd = []byte(`The `+"`"+`anytrue`+"`"+` transform returns `+"`"+`true`+"`"+` if any of the datapoints in the stream were true.

Given the following data:

`+"`"+``+"`"+``+"`"+`
false,false,false,true,false,false
`+"`"+``+"`"+``+"`"+`

`+"`"+`anytrue`+"`"+` will return:

`+"`"+``+"`"+``+"`"+`
true
`+"`"+``+"`"+``+"`"+`

### Why It's Useful

Oftentimes you might want to check something in a `+"`"+`while`+"`"+`, or in a `+"`"+`map`+"`"+`. A great example
would be to check which days you went to the gym:

`+"`"+``+"`"+``+"`"+`
distance(<latitude>,<longitude>) < 40 | while(day==next:day,anytrue)
`+"`"+``+"`"+``+"`"+`

The above transform will run a while loop while the current datapoint is part of the same day as the next datapoint, and check whether any of the GPS coordinates were within 40 meters of your gym.
`)

func docsTransformsAnytrueMdBytes() ([]byte, error) {
	return _docsTransformsAnytrueMd, nil
}

func docsTransformsAnytrueMd() (*asset, error) {
	bytes, err := docsTransformsAnytrueMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/anytrue.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsBucketMd = []byte(`The bucket transform allows you to put numbers into buckets of custom size.

For example, given this data:

`+"`"+``+"`"+``+"`"+`
2,16,84,-5,1
`+"`"+``+"`"+``+"`"+`
We can choose buckets of size `+"`"+`10`+"`"+` (starting from `+"`"+`0`+"`"+` by default)
`+"`"+``+"`"+``+"`"+`
bucket(10)
`+"`"+``+"`"+``+"`"+`
to get:
`+"`"+``+"`"+``+"`"+`json
"[0,10)", "[10,20)", "[80,90)","[-10,0)","[0,10)"
`+"`"+``+"`"+``+"`"+`

The bucket transform uses [interval notation][1]. To process the range in code, you can manually change the last character from `+"`"+`)`+"`"+` to `+"`"+`]`+"`"+`, and parse the result as a json array.

[1]: https://en.wikipedia.org/wiki/Interval_(mathematics)

### Histograms

Using the `+"`"+`bucket`+"`"+` transform in conjunction with the `+"`"+`map`+"`"+` transform, it is easy to generate histograms:

`+"`"+``+"`"+``+"`"+`
map(bucket(10),count)
`+"`"+``+"`"+``+"`"+`
Running this transform on the above data will give:
`+"`"+``+"`"+``+"`"+`json
[{
  "[-10,0)": 1,
  "[0,10)": 2,
  "[10,20)": 1,
  "[80,90)": 1
}]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsBucketMdBytes() ([]byte, error) {
	return _docsTransformsBucketMd, nil
}

func docsTransformsBucketMd() (*asset, error) {
	bytes, err := docsTransformsBucketMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/bucket.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsChangedMd = []byte(`The changed transform returns true if the current datapoint's data is different from the previous datapoint.

Given the following data:
`+"`"+``+"`"+``+"`"+`
1,2,3,3,4
`+"`"+``+"`"+``+"`"+`
running the `+"`"+`changed`+"`"+` transform on it will result in:
`+"`"+``+"`"+``+"`"+`
true,true,true,false,true
`+"`"+``+"`"+``+"`"+`

### Usage

The `+"`"+`changed`+"`"+` transform is useful whenever you don't care about the specific datapoints,
but when they change with respect to a certain metric.

Suppose you have a stream of activities gathered by a phone:
`+"`"+``+"`"+``+"`"+`
walking,walking,still,still,still,walking,walking
`+"`"+``+"`"+``+"`"+`

You usually care about the transitions between activities:
`+"`"+``+"`"+``+"`"+`
if changed
`+"`"+``+"`"+``+"`"+`
gives:
`+"`"+``+"`"+``+"`"+`
walking,still,running
`+"`"+``+"`"+``+"`"+`

Remember that each datapoint comes with a timestamp, so the length of each activity can be extracted by looking at the timestamp differences.
`)

func docsTransformsChangedMdBytes() ([]byte, error) {
	return _docsTransformsChangedMd, nil
}

func docsTransformsChangedMd() (*asset, error) {
	bytes, err := docsTransformsChangedMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/changed.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsContainsMd = []byte(``+"`"+`contains`+"`"+` permits you to check if a datapoint with a string data value contains the given substring:

`+"`"+``+"`"+``+"`"+`json
["Hello World!","hello world","hi there"]
`+"`"+``+"`"+``+"`"+`

Running the transform `+"`"+`contains("World")`+"`"+` will give:
`+"`"+``+"`"+``+"`"+`json
[true,false,false]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsContainsMdBytes() ([]byte, error) {
	return _docsTransformsContainsMd, nil
}

func docsTransformsContainsMd() (*asset, error) {
	bytes, err := docsTransformsContainsMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/contains.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsCountMd = []byte(`Count represents the total number of datapoints passed through the transform. It is equivalent to an `+"`"+`i`+"`"+` used in a loop over an array, with the difference that count starts from 1, rather than 0.

No matter what the datapoints, the sequence of data that count returns is:

`+"`"+``+"`"+``+"`"+`
1,2,3,4,5...
`+"`"+``+"`"+``+"`"+`

### Usage

#### Counting Mood

Suppose you want to count the number of times you were in a great mood:

`+"`"+``+"`"+``+"`"+`
filter $ >= 8 | count
`+"`"+``+"`"+``+"`"+`

The above transform will return only the datapoints where your mood rating was 8 or above, it will count them, and only return the last datapoint (which contains the full count).

#### Counting Visits

Suppose you want to find how many times you visited a friend:

`+"`"+``+"`"+``+"`"+`
distance(<latitude>,<longitude>) < 50 | filter changed | filter $ | count
`+"`"+``+"`"+``+"`"+`

The above transform finds when you were within 50 meters of the given coordinates (your friend's home), filters these datapoints so only changes remain (so each time you visit, you get one `+"`"+`true`+"`"+`, followed by a `+"`"+`false`+"`"+` when you leave), filter the false values, and count the number of times you visited. Note that `+"`"+`filter $ | count`+"`"+` can be replaced with `+"`"+`sum`+"`"+` in this case.

#### Counting Weekdays

Now you want to see which weekdays you use your computer the most. You can simply count the datapoints in your laptop's stream to see what days have most data:

`+"`"+``+"`"+``+"`"+`
map(weekday,count)
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsCountMdBytes() ([]byte, error) {
	return _docsTransformsCountMd, nil
}

func docsTransformsCountMd() (*asset, error) {
	bytes, err := docsTransformsCountMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/count.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsDistanceMd = []byte(`The `+"`"+`distance`+"`"+` transform computes the distance in meters from the current datapoint to its argument coordinates.

The datapoint is assumed to have `+"`"+`latitude`+"`"+` and `+"`"+`longitude`+"`"+` fields in decimal coordinates. It returns the distance in meters computed using the [Haversine formula](https://en.wikipedia.org/wiki/Haversine_formula).

`+"`"+``+"`"+``+"`"+`json
[{
  "latitude": 40.4277304,
  "longitude": -86.9170587
}]
`+"`"+``+"`"+``+"`"+`
Given the above stream, we find its distance to the chosen coordinate: `+"`"+`distance(40.426841,-86.9165106)`+"`"+`:
`+"`"+``+"`"+``+"`"+`json
[109.238]
`+"`"+``+"`"+``+"`"+`

The two coordinates above are about 109 meters apart.
`)

func docsTransformsDistanceMdBytes() ([]byte, error) {
	return _docsTransformsDistanceMd, nil
}

func docsTransformsDistanceMd() (*asset, error) {
	bytes, err := docsTransformsDistanceMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/distance.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsDtMd = []byte(`Every datapoint in a stream has a duration. This value is hidden when performing operations in PipeScript, since all operations are performed on a datapoint's `+"`"+`d`+"`"+` (data) field. To permit processing based upon duration in PipeScript, the `+"`"+`dt`+"`"+` transform exposes the datapoint's duration.

Remember that raw datapoints are in the form:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "t": 123456.23,
    "d": 4,
    "dt": 80.0
  }
]
`+"`"+``+"`"+``+"`"+`

When performing transforms, `+"`"+`$==4`+"`"+` will return `+"`"+`true`+"`"+`, since we operate on the "d" (data) field. But the duration is `+"`"+`80.0`+"`"+`, so the transform `+"`"+`dt`+"`"+` will result in the following datapoint:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "t": 123456.23,
    "d": 80.0,
    "dt": 80.0
  }
]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsDtMdBytes() ([]byte, error) {
	return _docsTransformsDtMd, nil
}

func docsTransformsDtMd() (*asset, error) {
	bytes, err := docsTransformsDtMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/dt.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsFilterMd = []byte(`The filter transform is used for filtering data.

Suppose the data portion of your dataset is as follows:

`+"`"+``+"`"+``+"`"+`
1,2,3,4,5,6
`+"`"+``+"`"+``+"`"+`

The transform:

`+"`"+``+"`"+``+"`"+`
filter $ >= 5
`+"`"+``+"`"+``+"`"+`

will leave you with:

`+"`"+``+"`"+``+"`"+`
5,6
`+"`"+``+"`"+``+"`"+`

Note that while convention is to use filter without parentheses (bash style), `+"`"+`filter`+"`"+` is a normal pipescript transform, and can be used as a function: `+"`"+`filter($ >= 5)`+"`"+`.

## and/or

PipeScript supports python-like and/or statements to build up a boolean:

`+"`"+``+"`"+``+"`"+`
filter $ > 5 and $ < 20
`+"`"+``+"`"+``+"`"+`

The above will only pass through datapoints between 5 and 20. Just like in other languages, you can use parentheses to force an order of operations.

PipeScript also has a built in negation:

`+"`"+``+"`"+``+"`"+`
filter not $ > 5
`+"`"+``+"`"+``+"`"+`

Combining and/or with not allows building up arbitrary conditions.
`)

func docsTransformsFilterMdBytes() ([]byte, error) {
	return _docsTransformsFilterMd, nil
}

func docsTransformsFilterMd() (*asset, error) {
	bytes, err := docsTransformsFilterMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/filter.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsFirstMd = []byte(`This is true when the datapoint is first in a sequence. It is useful mainly for filtering:

`+"`"+``+"`"+``+"`"+`
filter first or last
`+"`"+``+"`"+``+"`"+`

will return the first and last datapoint in your stream:

`+"`"+``+"`"+``+"`"+`
1,2,3,4,5
`+"`"+``+"`"+``+"`"+`

`+"`"+``+"`"+``+"`"+`
1,5
`+"`"+``+"`"+``+"`"+`

### Usage

This transform could be finding your wakeup time, based upon the first time you turn on your phone screen in the morning:

`+"`"+``+"`"+``+"`"+`
filter dayhour > 4 | while(day == next:day, first) | t
`+"`"+``+"`"+``+"`"+`

The above transform removes the datapoints taken from midnight to 4am (to filter out long nights), and then returns the first datapoint of each day, finaally returning only the timestamp.
`)

func docsTransformsFirstMdBytes() ([]byte, error) {
	return _docsTransformsFirstMd, nil
}

func docsTransformsFirstMd() (*asset, error) {
	bytes, err := docsTransformsFirstMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/first.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsIMd = []byte(`The `+"`"+`i`+"`"+` transform gives the index in the timeseries array, starting with 0

`+"`"+``+"`"+``+"`"+`
"data","data","data","data",...
`+"`"+``+"`"+``+"`"+`

gives:

`+"`"+``+"`"+``+"`"+`
0,1,2,3
`+"`"+``+"`"+``+"`"+`

When used within a pipe, it gives the index of datapoint _within the pipe_. The following transform filters half the data, and returns the indices at the second pipe location

`+"`"+``+"`"+``+"`"+`
filter i%2==0 | i
`+"`"+``+"`"+``+"`"+`

`+"`"+``+"`"+``+"`"+`
0,1
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsIMdBytes() ([]byte, error) {
	return _docsTransformsIMd, nil
}

func docsTransformsIMd() (*asset, error) {
	bytes, err := docsTransformsIMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/i.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsLastMd = []byte(``+"`"+`last`+"`"+` is true on the last datapoint of a sequence.

`+"`"+``+"`"+``+"`"+`
filter first or last
`+"`"+``+"`"+``+"`"+`

will return the first and last datapoint in your stream:

`+"`"+``+"`"+``+"`"+`
1,2,3,4,5
`+"`"+``+"`"+``+"`"+`

`+"`"+``+"`"+``+"`"+`
1,5
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsLastMdBytes() ([]byte, error) {
	return _docsTransformsLastMd, nil
}

func docsTransformsLastMd() (*asset, error) {
	bytes, err := docsTransformsLastMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/last.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsMapMd = []byte(`The map transform is an example of a transform which hijacks its second argument. Please note that it is only lightly related to the standard map function in most programming languages. It splits datapoints along its first argument, and runs independent scripts on each value:


Suppose your data stream is the following:
`+"`"+``+"`"+``+"`"+`json
[{
  "steps": 14,
  "activity": "walking"
},{
  "steps": 10,
  "activity": "running"
},{
  "steps": 12,
  "activity": "walking"
},{
  "steps": 5,
  "activity": "running"
}]
`+"`"+``+"`"+``+"`"+`


Mapping by activity:
`+"`"+``+"`"+``+"`"+`
map($("activity"), $("steps"):sum )
`+"`"+``+"`"+``+"`"+`

gives you
`+"`"+``+"`"+``+"`"+`json
[{
  "walking": 26,
  "running": 15
}]
`+"`"+``+"`"+``+"`"+`

The map transform split the dataset by its first argument (`+"`"+`$("activity")`+"`"+`), and performed the second transform (`+"`"+`$("steps"):sum`+"`"+`) on each subset independently. This allows directly returning the sum of datapoints where you were walking and running.

The map transform is frequently used in conjunction with the `+"`"+`reduce`+"`"+` transform for quick analysis.

Very common use for the map transform is splitting by time periods. For example, to get the total number of steps taken per weekday (suppose steps is a number stream, where the data value is number of steps taken)

`+"`"+``+"`"+``+"`"+`
map(weekday,sum)
`+"`"+``+"`"+``+"`"+`

One issue with the above transform is that it gives a total sum of all datapoints. We
might want to compute the *average* for a weekday

`+"`"+``+"`"+``+"`"+`
map(weekday, imap(day,sum):reduce(mean))
`+"`"+``+"`"+``+"`"+`

Note the imap transform within the map. It functions in exactly the same way as map, but returns all intermediate values (which allows it to be embedded in map). The above transform first splits by weekday, then further splits by day. It sums up the steps taken per day and finds their mean. That leaves with daily averages per weekday.

Another way to perform the same calculation is:

`+"`"+``+"`"+``+"`"+`
while(day==next:day,sum) | map(weekday,mean)
`+"`"+``+"`"+``+"`"+`

This sums up datapoints for each day, and only maps the summed datapoints. This is the recommended method for such computations.
`)

func docsTransformsMapMdBytes() ([]byte, error) {
	return _docsTransformsMapMd, nil
}

func docsTransformsMapMd() (*asset, error) {
	bytes, err := docsTransformsMapMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/map.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsMeanMd = []byte(`The `+"`"+`mean`+"`"+` transform returns the average of all datapoints:

### Input

`+"`"+``+"`"+``+"`"+`
1,2,6
`+"`"+``+"`"+``+"`"+`

### Output

`+"`"+``+"`"+``+"`"+`
3
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsMeanMdBytes() ([]byte, error) {
	return _docsTransformsMeanMd, nil
}

func docsTransformsMeanMd() (*asset, error) {
	bytes, err := docsTransformsMeanMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/mean.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsReduceMd = []byte(``+"`"+`reduce`+"`"+` performs a given transform on all the elements of a multi-element datapoint.

Suppose you have the following data:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "a": 3,
    "b": 5
  },
  {
    "q": 10,
    "z": 23
  }
]
`+"`"+``+"`"+``+"`"+`

Then the transform `+"`"+`reduce(sum)`+"`"+` will give you:

`+"`"+``+"`"+``+"`"+`json
[8, 33]
`+"`"+``+"`"+``+"`"+`

Reduce took each element in the given datapoint, and applied the transform `+"`"+`sum`+"`"+` to it.

## Usage

The reduce transform is particularly useful in conjunction with the `+"`"+`map`+"`"+` transform.

Suppose you want to find the average number of steps taken every weekday.

Running the following transform will give you a map of weekday to average step count:

`+"`"+``+"`"+``+"`"+`
while(day==$[1]:day,sum) | map(weekday, mean)
`+"`"+``+"`"+``+"`"+`

For example, a possible result of the above transform could be:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "Monday": 12243,
    "Tuesday": 13452,
    "Wednesday": 14523,
    "Thursday": 9543,
    "Friday": 20487,
    "Saturday": 3000,
    "Sunday": 4000
  }
]
`+"`"+``+"`"+``+"`"+`

You can now find the average per weekday by running `+"`"+`reduce(mean)`+"`"+`, giving a final transform:

`+"`"+``+"`"+``+"`"+`
while(day==$[1]:day,sum) | map(weekday, mean) | reduce(mean)
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsReduceMdBytes() ([]byte, error) {
	return _docsTransformsReduceMd, nil
}

func docsTransformsReduceMd() (*asset, error) {
	bytes, err := docsTransformsReduceMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/reduce.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsRegexMd = []byte(`The regex transform checks if the data string matches the given regex.

For example, given a regex to check for valid usernames: `+"`"+`regex('^[a-z0-9_-]{3,16}$')`+"`"+`, we get:

`+"`"+``+"`"+``+"`"+`json
[
  "Hello World!",
  "valid_username",
  "1"
]
`+"`"+``+"`"+``+"`"+`

`+"`"+``+"`"+``+"`"+`json
[false,true,false]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsRegexMdBytes() ([]byte, error) {
	return _docsTransformsRegexMd, nil
}

func docsTransformsRegexMd() (*asset, error) {
	bytes, err := docsTransformsRegexMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/regex.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsSumMd = []byte(`The `+"`"+`sum`+"`"+` transform sums up numeric values. Given the data:

`+"`"+``+"`"+``+"`"+`json
[2, 5, 1, 6]
`+"`"+``+"`"+``+"`"+`

running the transform `+"`"+`sum`+"`"+` will give:

`+"`"+``+"`"+``+"`"+`json
[14]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsSumMdBytes() ([]byte, error) {
	return _docsTransformsSumMd, nil
}

func docsTransformsSumMd() (*asset, error) {
	bytes, err := docsTransformsSumMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/sum.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsTMd = []byte(`Every datapoint in a stream has a timestamp. This timestamp is hidden when performing operations in PipeScript, since all operations are performed on a datapoint's `+"`"+`d`+"`"+` (data) field. To permit processing based upon timestamps in PipeScript, the `+"`"+`t`+"`"+` transform exposes the datapoint's timestamp.

Remember that raw datapoints are in the form:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "t": 123456.23,
    "d": 4,
    "dt": 0
  }
]
`+"`"+``+"`"+``+"`"+`

When performing transforms, `+"`"+`$==4`+"`"+` will return `+"`"+`true`+"`"+`, since we operate on the "d" (data) field. But the timestamp is `+"`"+`123456.23`+"`"+`, so the transform `+"`"+`t`+"`"+` will result in the following datapoint:

`+"`"+``+"`"+``+"`"+`json
[
  {
    "t": 123456.23,
    "d": 123456.23,
    "dt": 0
  }
]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsTMdBytes() ([]byte, error) {
	return _docsTransformsTMd, nil
}

func docsTransformsTMd() (*asset, error) {
	bytes, err := docsTransformsTMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/t.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsTshiftMd = []byte(`This transform is not particularly useful for PipeScript by itself, but becomes very frequently used in dataset and merge queries.

Every datapoint has a data portion, as well as a timestamp, which is hidden from computations in PipeScript by default. `+"`"+`tshift`+"`"+` shifts the timestamps of a stream by the given amount in seconds. This allows making it seem like the data of a stream came before/after its actual timestamps. This is useful in datasets, since a tshift can allow interpolating between different time ranges - it allows asking questions such as "does exercise today impact my mood a week later?". The datapoints corresponding to mood can be tshifted back by a week to correspond directly to the original datapoints where your exercise data is shown.
`)

func docsTransformsTshiftMdBytes() ([]byte, error) {
	return _docsTransformsTshiftMd, nil
}

func docsTransformsTshiftMd() (*asset, error) {
	bytes, err := docsTransformsTshiftMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/tshift.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsWcMd = []byte(`This transform counts the words in a string:

`+"`"+``+"`"+``+"`"+`json
["Hello World!"]
`+"`"+``+"`"+``+"`"+`

Running `+"`"+`wc`+"`"+` on the above returns:
`+"`"+``+"`"+``+"`"+`json
[2]
`+"`"+``+"`"+``+"`"+`
`)

func docsTransformsWcMdBytes() ([]byte, error) {
	return _docsTransformsWcMd, nil
}

func docsTransformsWcMd() (*asset, error) {
	bytes, err := docsTransformsWcMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/wc.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _docsTransformsWhileMd = []byte(`Oftentimes, you don't really want all of the raw data from your stream. You are usually interested in timed aggregations or other such things. This is the main use case of the `+"`"+`while`+"`"+` transform.

This transform performs a while loop on its second argument while its first argument is true. When the first argument becomes false, it returns the resulting datapoint, and begins the next loop.

This allows summing up datapoints over specified time periods. For example, `+"`"+`day==$[1]:day`+"`"+` is true if the current datapoint and next datapoint in the stream have timestamps from the same day.

This allows you to write a transform performing an aggregation per day:

`+"`"+``+"`"+``+"`"+`
while(day==$[1]:day, sum)
`+"`"+``+"`"+``+"`"+`

The above transform loops through datapoints while they come from the same day, and sums their values. Once the next datapoint is a different day than the current one, it ends the while loop, and returns the sum, giving the sum of all datapoints that day. It then starts a loop for the next day.

The transform can also be used for smoothing. Suppose you want to smooth your data every three datapoints:

`+"`"+``+"`"+``+"`"+`
while(i%3!=0, mean)
`+"`"+``+"`"+``+"`"+`

This returns the mean of every consecutive three datapoints, making it easy to do a basic smoothing procedure.

## Advanced Usage

The transform can also be used to implement error bars:

`+"`"+``+"`"+``+"`"+`
while(day==$[1]:day, {"max": max, "min": min, "mean": mean})
`+"`"+``+"`"+``+"`"+`

This transform returns the mean, max and min datapoint for the day all at once, allowing to plot with error bars.
`)

func docsTransformsWhileMdBytes() ([]byte, error) {
	return _docsTransformsWhileMd, nil
}

func docsTransformsWhileMd() (*asset, error) {
	bytes, err := docsTransformsWhileMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "docs/transforms/while.md", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"docs/transforms/$.md": docsTransformsMd,
	"docs/transforms/alltrue.md": docsTransformsAlltrueMd,
	"docs/transforms/anytrue.md": docsTransformsAnytrueMd,
	"docs/transforms/bucket.md": docsTransformsBucketMd,
	"docs/transforms/changed.md": docsTransformsChangedMd,
	"docs/transforms/contains.md": docsTransformsContainsMd,
	"docs/transforms/count.md": docsTransformsCountMd,
	"docs/transforms/distance.md": docsTransformsDistanceMd,
	"docs/transforms/dt.md": docsTransformsDtMd,
	"docs/transforms/filter.md": docsTransformsFilterMd,
	"docs/transforms/first.md": docsTransformsFirstMd,
	"docs/transforms/i.md": docsTransformsIMd,
	"docs/transforms/last.md": docsTransformsLastMd,
	"docs/transforms/map.md": docsTransformsMapMd,
	"docs/transforms/mean.md": docsTransformsMeanMd,
	"docs/transforms/reduce.md": docsTransformsReduceMd,
	"docs/transforms/regex.md": docsTransformsRegexMd,
	"docs/transforms/sum.md": docsTransformsSumMd,
	"docs/transforms/t.md": docsTransformsTMd,
	"docs/transforms/tshift.md": docsTransformsTshiftMd,
	"docs/transforms/wc.md": docsTransformsWcMd,
	"docs/transforms/while.md": docsTransformsWhileMd,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"docs": &bintree{nil, map[string]*bintree{
		"transforms": &bintree{nil, map[string]*bintree{
			"$.md": &bintree{docsTransformsMd, map[string]*bintree{}},
			"alltrue.md": &bintree{docsTransformsAlltrueMd, map[string]*bintree{}},
			"anytrue.md": &bintree{docsTransformsAnytrueMd, map[string]*bintree{}},
			"bucket.md": &bintree{docsTransformsBucketMd, map[string]*bintree{}},
			"changed.md": &bintree{docsTransformsChangedMd, map[string]*bintree{}},
			"contains.md": &bintree{docsTransformsContainsMd, map[string]*bintree{}},
			"count.md": &bintree{docsTransformsCountMd, map[string]*bintree{}},
			"distance.md": &bintree{docsTransformsDistanceMd, map[string]*bintree{}},
			"dt.md": &bintree{docsTransformsDtMd, map[string]*bintree{}},
			"filter.md": &bintree{docsTransformsFilterMd, map[string]*bintree{}},
			"first.md": &bintree{docsTransformsFirstMd, map[string]*bintree{}},
			"i.md": &bintree{docsTransformsIMd, map[string]*bintree{}},
			"last.md": &bintree{docsTransformsLastMd, map[string]*bintree{}},
			"map.md": &bintree{docsTransformsMapMd, map[string]*bintree{}},
			"mean.md": &bintree{docsTransformsMeanMd, map[string]*bintree{}},
			"reduce.md": &bintree{docsTransformsReduceMd, map[string]*bintree{}},
			"regex.md": &bintree{docsTransformsRegexMd, map[string]*bintree{}},
			"sum.md": &bintree{docsTransformsSumMd, map[string]*bintree{}},
			"t.md": &bintree{docsTransformsTMd, map[string]*bintree{}},
			"tshift.md": &bintree{docsTransformsTshiftMd, map[string]*bintree{}},
			"wc.md": &bintree{docsTransformsWcMd, map[string]*bintree{}},
			"while.md": &bintree{docsTransformsWhileMd, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

