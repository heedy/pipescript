package core

import (
	"errors"
	"reflect"

	"github.com/connectordb/pipescript"
)

var iflast *pipescript.Script

func init() {
	// Manually generate the if last script
	var err error
	last, err := Last.Script(nil)
	if err != nil {
		panic("Could not generate script for 'if last' for Reduce Transform " + err.Error())
	}
	iflast, err = If.Script([]*pipescript.Script{last})
	if err != nil {
		panic("Could not generate script for 'if last' for Reduce Transform " + err.Error())
	}
}

type reduceTransform struct {
	script *pipescript.Script
}

func (t *reduceTransform) Copy() (pipescript.TransformInstance, error) {
	return &reduceTransform{t.script}, nil
}

func (t *reduceTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// We only support maps for now in the data of our datapoint
	dmap := reflect.ValueOf(te.Datapoint.Data)
	if dmap.Kind() != reflect.Map {
		return nil, errors.New("Datapoint's data can't be reduced")
	}

	k := dmap.MapKeys()
	dpa := make([]pipescript.Datapoint, 0, len(k))
	for i := range k {
		dpa = append(dpa, pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: dmap.MapIndex(k[i]).Interface()})
	}

	dpi := pipescript.NewDatapointArrayIterator(dpa)

	// Now set up the script - and end it with an if last so only one result is given
	s, err := t.script.Copy()
	if err != nil {
		return nil, err
	}
	ifl, _ := iflast.Copy() // Shouldn't error
	s.Append(ifl)
	s.SetInput(dpi)

	// And now return the resutl!
	v, err := s.Next()
	if v == nil {
		return &pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: nil}, err
	}
	return v, err
}

var Reduce = pipescript.Transform{
	Name: "reduce",
	Description: `Takes a json object, and considers each field to be a separate datapoint's data.
It then hijacks its argument, and performs the given transform on all of the fields, returning the result.
It is mainly useful as a companion to the map transform, with which it is possible to create a simple map/reduce pipeline`,
	OneToOne: true,
	Args: []pipescript.TransformArg{
		{
			Description: "The script to instantiate to perform on all elements of input",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &reduceTransform{args[0]}}, nil
	},
}
