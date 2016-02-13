package core

import (
	"errors"

	"github.com/connectordb/duck"
	"github.com/connectordb/pipescript"
)

// A transform that sets values in objects
type setTransform struct {
	Obj string
}

func (t *setTransform) Copy() (pipescript.TransformInstance, error) {
	return &setTransform{t.Obj}, nil
}

func (t *setTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	// Since We can't use the original data map, we must create a copy!
	k, ok := duck.Keys(te.Datapoint.Data)
	if !ok {
		return nil, errors.New("Can't read object keys")
	}
	resultmap := make(map[string]interface{})
	for i := range k {
		v, err := te.Datapoint.Get(k[i])
		if err != nil {
			return nil, err
		}
		resultmap[k[i]] = v
	}

	//Next, set the value we want
	resultmap[t.Obj] = te.Args[0].Data

	return te.Set(resultmap)
}

var Set = pipescript.Transform{
	Name:        "set",
	Description: "Allows setting object values",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		{
			Description: "The name of field to set",
			Optional:    false,
			Default:     nil,
			Constant:    true,
		},
		{
			Description: "The value to set the field to",
			Optional:    false,
			Default:     nil,
			Constant:    false,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		// Get the value to set
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		s, err := dp.DataString()
		if err != nil {
			return nil, err
		}

		return &pipescript.TransformInitializer{Args: args[1:2], Transform: &setTransform{s}}, nil
	},
}
