package core

import (
	"errors"
	"fmt"

	"github.com/connectordb/pipescript"
)

// The maximum number of elements in a map
var SplitMax = 1000

type mapTransform struct {
	script *pipescript.Script // The uninitialized script to be used for splitting
}

func (t *mapTransform) Copy() (pipescript.TransformInstance, error) {
	return &mapTransform{t.script}, nil
}

func (t *mapTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	scriptmap := make(map[string]*pipescript.Script)
	datamap := make(map[string]interface{})
	iter := &pipescript.SingleDatapointIterator{}
	var lasttimestamp float64

	for !te.IsFinished() {
		lasttimestamp = te.Datapoint.Timestamp

		// Convert the key value to string
		v, err := te.Args[0].DataString()
		if err != nil {
			return nil, err
		}

		//Check if the value exists
		s, ok := scriptmap[v]
		if !ok {
			if len(scriptmap) >= SplitMax {
				return nil, fmt.Errorf("Reached maximum map amount %d.", SplitMax)
			}

			// Initialize the new script, and add it to our map
			s, err = t.script.Copy()
			if err != nil {
				return nil, err
			}
			//Set the script input to be the internal iterator
			s.SetInput(iter)
			scriptmap[v] = s
		}

		//Send the current datapoint to the iterator
		iter.Set(te.Datapoint, nil)
		dp, err := s.Next()
		if err != nil {
			return nil, err
		}

		// Set the data in our map
		datamap[v] = dp.Data

		te = ti.Next()

	}

	return &pipescript.Datapoint{Timestamp: lasttimestamp, Data: datamap}, te.Error

}

// Map splits the dtaapoints by its first argument
var Map = pipescript.Transform{
	Name: "map",
	Description: `Splits the script by the first argument's value, creating new instances of the second argument's script.
Think of it as a switch statement where each choice has copies of the same code.
It is very useful for splitting by time. For example:
"split(weekday,count)" will return {"Monday": ...,"Tuesday":...} with the number of datapoints that happened in each day.`,
	Args: []pipescript.TransformArg{
		{
			Description: "The value to split on. This must be something that can be converted to string.",
		},
		{
			Description: "The script to instantiate for each different value of the first argument.",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		if args[1].Peek {
			return nil, errors.New("Map cannot be used with transforms that peek.")
		}
		return &pipescript.TransformInitializer{
			Args:      []*pipescript.Script{args[0]},
			Transform: &mapTransform{args[1]},
		}, nil
	},
}
