package core

import (
	"errors"

	"github.com/connectordb/duck"
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

// A transform that deletes values in objects
type delTransform struct {
	Obj string
}

func (t *delTransform) Copy() (pipescript.TransformInstance, error) {
	return &delTransform{t.Obj}, nil
}

func (t *delTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
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
		if k[i] != t.Obj {
			v, err := te.Datapoint.Get(k[i])
			if err != nil {
				return nil, err
			}
			resultmap[k[i]] = v
		}
	}

	return te.Set(resultmap)
}

var Del = pipescript.Transform{
	Name:          "del",
	Description:   "Allows deleting object values",
	Documentation: string(resources.MustAsset("docs/transforms/del.md")),
	OneToOne:      true,
	Stateless:     true,
	Args: []pipescript.TransformArg{
		{
			Description: "The name of field to delete",
			Optional:    false,
			Default:     nil,
			Constant:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		// Get the value to del
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		s, err := dp.DataString()
		if err != nil {
			return nil, err
		}

		return &pipescript.TransformInitializer{Transform: &delTransform{s}}, nil
	},
}
