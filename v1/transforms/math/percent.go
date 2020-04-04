package math

import (
	"errors"
	"reflect"

	"github.com/heedy/duck"
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type percentTransform struct{}

func (t *percentTransform) Copy() (pipescript.TransformInstance, error) {
	return &percentTransform{}, nil
}

func (t *percentTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
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
	resultFloatMap := make(map[string]float64)
	sumvalue := float64(0.0)

	for i := range k {
		ks, ok := k[i].Interface().(string)
		if !ok {
			return nil, errors.New("Unable to transform non-string keys")
		}

		pointdata := dmap.MapIndex(k[i]).Interface()

		f, ok := duck.Float(pointdata)
		if !ok {
			return nil, errors.New("Could not convert object value to float")
		}
		if f < 0 {
			resultFloatMap[ks] = float64(0.0)
		} else {
			resultFloatMap[ks] = f
			sumvalue += f
		}
	}

	// Now normalize all values
	if sumvalue > 0 {
		for k := range resultFloatMap {
			resultFloatMap[k] = resultFloatMap[k] / sumvalue
		}
	}

	// And now convert to a map[string]interface{}
	resultMap := make(map[string]interface{})

	for k := range resultFloatMap {
		resultMap[k] = resultFloatMap[k]
	}

	return te.Set(resultMap)
}

var Percent = pipescript.Transform{
	Name:          "percent",
	Description:   `Normalizes the values of a JSON object`,
	Documentation: string(resources.MustAsset("docs/transforms/percent.md")),
	OneToOne:      true,
	Stateless:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &percentTransform{}}, nil
	},
}
