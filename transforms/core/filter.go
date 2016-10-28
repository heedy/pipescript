package core

import (
	"errors"
	"reflect"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type filterTransform struct {
	script *pipescript.Script
}

func (t *filterTransform) Copy() (pipescript.TransformInstance, error) {
	return &filterTransform{t.script}, nil
}

func (t *filterTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// We only support maps for now in the data of our datapoint
	dmap := reflect.ValueOf(te.Datapoint.Data)
	if dmap.Kind() != reflect.Map {
		return nil, errors.New("Datapoint's data can't be filterd")
	}

	k := dmap.MapKeys()
	resultMap := make(map[string]interface{})
	for i := range k {
		ks, ok := k[i].Interface().(string)
		if !ok {
			return nil, errors.New("Unable to transform non-string keys")
		}
		// Now set up the script - and end it with an if last so only one result is given
		s, err := t.script.Copy()
		if err != nil {
			return nil, err
		}

		pointdata := dmap.MapIndex(k[i]).Interface()

		dpi := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
			pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: pointdata},
		})
		s.SetInput(dpi)

		// And now return the result!
		v, err := s.Next()
		if v == nil || err != nil {
			return &pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: nil}, err
		}
		b, err := v.Bool()
		if err != nil {
			return nil, err
		}
		if !b {
			resultMap[ks] = pointdata
		}
	}

	return te.Set(resultMap)
}

var Filter = pipescript.Transform{
	Name: "filter",
	Description: `Takes a json object, and considers each field to be a separate datapoint's data.
It removes all elements for which its first argument returns true (filters).`,
	Documentation: string(resources.MustAsset("docs/transforms/filter.md")),
	OneToOne:      true,
	Args: []pipescript.TransformArg{
		{
			Description: "The script to instantiate to perform on all elements of input, one at a time.",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		ifl, _ := iflast.Copy() // Shouldn't error
		args[0].Append(ifl)
		return &pipescript.TransformInitializer{Transform: &filterTransform{args[0]}}, nil
	},
}
