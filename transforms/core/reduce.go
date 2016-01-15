package core

import (
	"errors"
	"fmt"

	"github.com/connectordb/pipescript"
)

// The maximum number of elements in a reduce
var SplitMax = 1000

type reduceTransform struct {
	script    *pipescript.Script // The uninitialized script to be used for splitting
	iter      *pipescript.SingleDatapointIterator
	scriptmap map[string]*pipescript.Script // Map of initialized scripts
	datamap   map[string]interface{}        // Map of data associated with scripts
}

func (t *reduceTransform) Copy() (pipescript.TransformInstance, error) {
	var err error
	scriptmap := make(map[string]*pipescript.Script)
	for i, val := range t.scriptmap {
		scriptmap[i], err = val.Copy()
		if err != nil {
			return nil, err
		}
	}
	datamap := make(map[string]interface{})
	for i, val := range t.datamap {
		datamap[i] = val // No need to worry abotu copying datapoints
	}
	return &reduceTransform{t.script, &pipescript.SingleDatapointIterator{}, scriptmap, datamap}, nil
}

func (t *reduceTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {

		// We need to take special care of finished sequences by clearing out the map, since the script may be reused
		t.scriptmap = make(map[string]*pipescript.Script)
		t.datamap = make(map[string]interface{})

		return te.Get()
	}

	// Convert the key value to string
	v, err := te.Args[0].String()
	if err != nil {
		return nil, err
	}

	//Check if the value exists
	s, ok := t.scriptmap[v]
	if !ok {
		if len(t.scriptmap) >= SplitMax {
			return nil, fmt.Errorf("Reached maximum reduce amount %d.", SplitMax)
		}

		// Initialize the new script, and add it to our map
		s, err = t.script.Copy()
		if err != nil {
			return nil, err
		}
		//Set the script input to be the internal iterator
		s.SetInput(t.iter)
		t.scriptmap[v] = s
	}

	//Send the current datapoint to the iterator
	t.iter.Set(te.Datapoint, nil)
	dp, err := s.Next()
	if err != nil {
		return nil, err
	}

	// Set the data in our map
	t.datamap[v] = dp.Data

	//Return the map
	return te.Set(t.datamap)
}

var reduce = pipescript.Transform{
	Name: "reduce",
	Description: `Splits the script by the first argument's value, creating new instances of the second argument's script.
Think of it as a switch statement where each choice has copies of the same code. It can also be thought of as a map reduce which returns intermediate results on the stream.
It is very useful for splitting by time. For example:
"split weekday {count} | if last" will return {"monday": ...,"tuesday":...} with the number of datapoints that happened in each day.`,
	OneToOne: true,
	Args: []pipescript.TransformArg{
		{
			Description: "The value to split on. This must be something that can be converted to string.",
		},
		{
			Description: "The script to instantiate for each different value of the first argument.",
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		if args[1].Peek {
			return nil, errors.New("Reduce cannot be used with transforms that peek.")
		}
		scriptmap := make(map[string]*pipescript.Script)
		datamap := make(map[string]interface{})
		return &pipescript.TransformInitializer{
			Args:      []*pipescript.Script{args[0]},
			Transform: &reduceTransform{args[1], &pipescript.SingleDatapointIterator{}, scriptmap, datamap},
		}, nil
	},
}
