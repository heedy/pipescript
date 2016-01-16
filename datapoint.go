package pipescript

import (
	"encoding/json"
	"fmt"

	"github.com/connectordb/duck"
)

// Datapoint is the core data type which is passed through the transform system. It encodes
// an element of a time series.
// Note that the data must be a simple structure, and if it is a map,
// it must be map[string] due to annoyances in reflection. While deeply nested data
// is supported, its use is not recommended. Data cannot be something that won't work
// simply with DeepCopy. If you are unsure what works, here is a couple things:
//
// - basic data types (int,string,float, etc)
// - map[string]interface{}
// - []int, []map[string]...
// - struct (fields recognized by duck (github.com/connectordb/duck))
//
// What doesn't work: pointers - make sure datapoints don't use pointers to data, as
// Transforms assume that they are able to Copy() a datapoint without affecting another
// Datapoint's value.
type Datapoint struct {
	Timestamp float64     `json:"t" msgpack:"t,omitempty" duck:"t"` // The time stamp in UNIX seconds for the current point
	Data      interface{} `json:"d" msgpack:"d,omitempty" duck:"d"` // The data associated with the datapoint.
}

// String returns the entire datapoint as a json string. Notice that this function conforms to
// the stringer interface, and as such is not the data string. If you want a function equivalent to Int() or Bool()
// you can use DataString()
func (d *Datapoint) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

// Copy creates a deep copy of the datapoint and its data. Modifying the data of the copied
// datapoint will not change data of the original datapoint. Wrapper for duck.Copy
func (d *Datapoint) Copy() *Datapoint {
	var d2 Datapoint
	d2.Timestamp = d.Timestamp

	d2.Data, _ = duck.Copy(d.Data)

	return &d2
}

// Int attempts to get the Data of a Datapoint as an integer. Wrapper for duck.Int
func (d *Datapoint) Int() (int64, error) {
	v, ok := duck.Int(d.Data)
	if !ok {
		return 0, fmt.Errorf("Could not convert '%v' to Int", d.Data)
	}
	return v, nil
}

// Float attempts to get the Data of a Datapoint as a floating point number. Wrapper for duck.Float
func (d *Datapoint) Float() (float64, error) {
	v, ok := duck.Float(d.Data)
	if !ok {
		return 0, fmt.Errorf("Could not convert '%v' to Float", d.Data)
	}
	return v, nil
}

// Bool attempts to get the Data of a Datapoint as a boolean. Wrapper for duck.Bool
func (d *Datapoint) Bool() (bool, error) {
	v, ok := duck.Bool(d.Data)
	if !ok {
		return false, fmt.Errorf("Could not convert '%v' to Bool", d.Data)
	}
	return v, nil
}

// DataString attempts to get the Data of a Datapoint as a string. Wrapper for duck.String
func (d *Datapoint) DataString() (string, error) {
	v, ok := duck.String(d.Data)
	if !ok {
		return "", fmt.Errorf("Could not convert '%v' to String", d.Data)
	}
	return v, nil
}

// Get attempts to get a sub-element from a Datapoint's data. The sub-element can be an array index, a
// map element, etc. Wrapper for duck.Get
func (d *Datapoint) Get(element interface{}) (interface{}, error) {
	v, ok := duck.Get(d.Data, element)
	if !ok {
		return nil, fmt.Errorf("Could not find '%v' in '%v'", element, d.Data)
	}
	return v, nil
}

// Set attempts to write to a sub-element for the Datapoint. If the datapoint does not have the given sub-element,
// then set fails. It can write to arrays (only elements that exist, cannot append), or to map[string]interface{},
// where it is capable of inserting new values. Wraps duck.Set
func (d *Datapoint) Set(element interface{}, newdata interface{}) error {
	if !duck.Set(d.Data, newdata, element) {
		return fmt.Errorf("Could not set '%v' in %v", element, d.Data)
	}
	return nil
}
