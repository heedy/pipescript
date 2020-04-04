package core

import (
	"errors"
	"reflect"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type applyTransform struct {
	script *pipescript.Script
}

func (t *applyTransform) Copy() (pipescript.TransformInstance, error) {
	return &applyTransform{t.script}, nil
}

func (t *applyTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// We only support maps for now in the data of our datapoint
	dmap := reflect.ValueOf(te.Datapoint.Data)
	if dmap.Kind() != reflect.Map {
		return nil, errors.New("Must have object-valued datapoint")
	}

	k := dmap.MapKeys()
	resultMap := make(map[string]interface{})

	for i := range k {
		ks, ok := k[i].Interface().(string)
		if !ok {
			return nil, errors.New("Unable to transform non-string keys")
		}
		// Now set up the script
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
		if err != nil {
			return nil, err
		}
		if v != nil {
			resultMap[ks] = v.Data
		}
	}

	return te.Set(resultMap)
}

var Apply = pipescript.Transform{
	Name:          "apply",
	Description:   `Applies the given transform to each value of a JSON object`,
	Documentation: string(resources.MustAsset("docs/transforms/apply.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The script to run on each element of the object",
			Hijacked:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &applyTransform{args[0]}}, nil
	},
}
