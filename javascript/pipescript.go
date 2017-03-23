package main

import (
	"bytes"
	"io"
	"math"
	"strings"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/bytestreams"
	"github.com/connectordb/pipescript/interpolator"
	"github.com/connectordb/pipescript/interpolator/interpolators"
	"github.com/connectordb/pipescript/transforms"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// Register all transforms & Interpolators
	transforms.Register()
	interpolators.Register()

	ps := map[string]interface{}{
		"Script":        New,
		"transforms":    pipescript.TransformRegistry,
		"interpolators": interpolator.InterpolatorRegistry,
	}

	// Make it usable in script tags
	js.Global.Set("pipescript", ps)

	// Make it usable in node. Note that the above makes it register in global
	// context also, which can't really be avoided easily
	if js.Module != js.Undefined {
		js.Module.Set("exports", ps)
	}

}

type Script struct {
	*js.Object
	scriptstring string
	script       *pipescript.Script
}

// Run runs PipeScript on the input, using the given input and output types
// The types supported by PipeScript for both input and output are:
// "datapoint" - datapoint representation (same as used internally)
// "json" - general json document
// "csv" - general csv document
// The same script can ONLY be run twice if it did not have an error.
func (s *Script) Run(input string, inputType string, outputType string, timestamphint string, notimestamp bool) string {
	if inputType == "undefined" {
		inputType = "dp"
	}
	if timestamphint == "undefined" {
		timestamphint = ""
	}

	// Read in the input as the correct type
	// TODO: currently only datapoint is supported
	r := strings.NewReader(input)

	var dpr pipescript.DatapointIterator
	var err error
	switch inputType {
	case "dp":
		dpr, err = bytestreams.NewDatapointReader(r)
		if err != nil {
			panic(err)
		}
	case "json":
		dpr, err = bytestreams.NewJSONDatapointReader(r, timestamphint, notimestamp)
		if err != nil {
			panic(err)
		}
	case "csv":
		dpr, err = bytestreams.NewCSVDatapointReader(r, timestamphint, notimestamp)
		if err != nil {
			panic(err)
		}
	default:
		panic("Unrecognized input format " + inputType)
	}
	s.script.SetInput(dpr)

	var jr io.Reader
	jr, err = bytestreams.NewJsonReader(s.script, "[\n", ",\n", "\n]", "", "  ")
	s.errorPanic(err)

	// Now create the output buffer
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(jr)
	s.errorPanic(err)
	return buf.String()
}

// The special datapoint type used for interacting with javascript
type Datapoint struct {
	*js.Object
	Data      interface{} `js:"d"`
	Timestamp float64     `js:"t"`
}

// DatapointArrayIterator is a DatapointIterator which iterates through the given array one datapoint
// at a time. Copied verbatim from the original array iterator
type DatapointArrayIterator struct {
	Datapoints []*Datapoint

	i int // i is the current location in the array
}

func NewDatapointArrayIterator(dp []*Datapoint) *DatapointArrayIterator {
	return &DatapointArrayIterator{dp, 0}
}

// Next returns the next datapoint in the array
func (d *DatapointArrayIterator) Next() (*pipescript.Datapoint, error) {
	if d.i < len(d.Datapoints) {
		dp := d.Datapoints[d.i]
		d.i++
		if math.IsNaN(dp.Timestamp) {
			panic("Datapoint didn't include timestamp field")
		}

		return &pipescript.Datapoint{Timestamp: dp.Timestamp, Data: dp.Data}, nil
	}
	return nil, nil
}

// Transform performs the given transform on a Datapoint array.
// It returns the result array. Remember to check the error response
func (s *Script) Transform(dpa []*Datapoint) []map[string]interface{} {
	s.script.SetInput(NewDatapointArrayIterator(dpa))

	// Now get the result into an array
	res := make([]map[string]interface{}, 0, len(dpa))
	dp, err := s.script.Next()
	for err == nil && dp != nil {
		res = append(res, map[string]interface{}{"t": dp.Timestamp, "d": dp.Data})
		dp, err = s.script.Next()
	}
	s.errorPanic(err)
	return res
}

// an errorPanic performs a cleanup before panicing. This allows running
// a pipeline after it has an error.
func (s *Script) errorPanic(err error) {
	if err != nil {
		s.script, _ = pipescript.Parse(s.scriptstring)
		panic(err)
	}
}

// New creates a new pipescript javascript object
func New(scriptstring string) *js.Object {
	s, err := pipescript.Parse(scriptstring)
	if err != nil {
		panic(err)
	}
	return js.MakeWrapper(&Script{scriptstring: scriptstring, script: s})
}
