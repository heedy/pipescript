package main

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/bytestreams"
	"github.com/connectordb/pipescript/transforms"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// Register all transforms
	transforms.Register()

	// Make it usable in script tags
	js.Global.Set("pipescript", map[string]interface{}{
		"Script":     New,
		"Transforms": Transforms,
	})

	// Make it usable in node. Note that the above makes it register in global
	// context also, which can't really be avoided easily
	js.Module.Get("exports").Set("Script", New)
	js.Module.Get("exports").Set("Transforms", New)

}

type Script struct {
	script       *pipescript.Script
	errorMessage string
}

// IsValid returns whether the script parsed correctly
func (s *Script) IsValid() bool {
	return s.script != nil
}

// Returns the error message
func (s *Script) Error() string {
	return s.errorMessage
}

// Run runs PipeScript on the input, using the given input and output types
// The types supported by PipeScript for both input and output are:
// "datapoint" - datapoint representation (same as used internally)
// "json" - general json document
// "csv" - general csv document
// The same script can ONLY be run twice if it did not have an error.
func (s *Script) Run(input string, inputType string, outputType string) string {
	if s.script == nil || s.errorMessage != "" {
		return ""
	}

	// Read in the input as the correct type
	// TODO: currently only datapoint is supported
	r := strings.NewReader(input)

	dpr, err := bytestreams.NewDatapointReader(r)
	if err != nil {
		s.errorMessage = err.Error()
		return ""
	}
	s.script.SetInput(dpr)

	var jr io.Reader
	jr, err = bytestreams.NewJsonReader(s.script, "[\n", ",\n", "\n]", "", "  ")
	if err != nil {
		s.errorMessage = err.Error()
		return ""
	}

	// Now create the output buffer
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(jr)
	if err != nil {
		s.errorMessage = err.Error()
		return ""
	}
	return buf.String()
}

// New creates a new pipescript javascrit object
func New(scriptstring string) *js.Object {
	errorstring := ""
	s, err := pipescript.Parse(scriptstring)
	if err != nil {
		errorstring = err.Error()
	}
	return js.MakeWrapper(&Script{s, errorstring})
}

// Returns a string json object of the documentation
func Transforms() string {
	b, _ := json.Marshal(pipescript.TransformRegistry)

	return string(b)
}
