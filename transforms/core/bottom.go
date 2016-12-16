package core

import (
	"errors"
	"reflect"
	"sort"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

// BottomTransform is same as Top transform with reverse functionality
type BottomTransform struct {
	number int
	script *pipescript.Script
}

func (t *BottomTransform) Copy() (pipescript.TransformInstance, error) {
	return &BottomTransform{t.number, t.script}, nil
}

func (t *BottomTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
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

	// see top.go for keyPair declarations
	kpl := make(KeyPairList, dmap.Len())

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

		v, err := s.Next()
		if err != nil {
			return &pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp, Data: nil}, err
		}
		f, err := v.Float()
		if err != nil {
			return nil, err
		}

		kpl[i] = keyPairValue{
			key:    ks,
			value:  pointdata,
			weight: f,
		}

	}

	// We have set up the kpl array. Now sort it
	sort.Sort(kpl)

	resultMap := make(map[string]interface{})

	for i := range kpl {
		if i >= t.number {
			break
		}
		curval := kpl[i]
		resultMap[curval.key] = curval.value
	}

	return te.Set(resultMap)
}

var Bottom = pipescript.Transform{
	Name:          "bottom",
	Description:   `Takes a json object, and returns the bottom n elements`,
	Documentation: string(resources.MustAsset("docs/transforms/bottom.md")),
	OneToOne:      true,
	Args: []pipescript.TransformArg{
		{
			Description: "The number of elements to retain of the object",
		},
		{
			Description: "The script to run ($ default) to generate weights over which to sort",
			Hijacked:    true,
			Optional:    true,
			Default:     nil, // THIS HAS TO BE SET IN init.go to identity

		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		num, err := dp.Int()
		if err != nil {
			return nil, err
		}
		ifl, _ := iflast.Copy() // Shouldn't error
		args[1].Append(ifl)
		return &pipescript.TransformInitializer{Transform: &BottomTransform{int(num), args[1]}}, nil
	},
}
