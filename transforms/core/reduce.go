package core

import (
	"errors"
	"reflect"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

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
It then hijacks its argument, and performs the given transform on all of the fields, returning the result.`,
	Documentation: string(resources.MustAsset("docs/transforms/reduce.md")),
	OneToOne:      true,
	Args: []pipescript.TransformArg{
		{
			Description: "The script to instantiate to perform on all elements of input",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		ifl, _ := iflast.Copy() // Shouldn't error
		args[0].Append(ifl)
		return &pipescript.TransformInitializer{Transform: &reduceTransform{args[0]}}, nil
	},
}
