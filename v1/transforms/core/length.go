package core

import (
	"github.com/heedy/duck"
	"github.com/heedy/pipescript"
)

type lengthTransform struct{}

func (t *lengthTransform) Copy() (pipescript.TransformInstance, error) {
	return &lengthTransform{}, nil
}

func (t *lengthTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	i, ok := duck.Length(te.Datapoint.Data)
	if !ok {
		return te.Set(0)
	}

	return te.Set(i)
}

var Length = pipescript.Transform{
	Name:         "length",
	Description:  "Returns the length of the given data (returns 0 on non-array-like)",
	OutputSchema: `{"type": "integer"}`,
	OneToOne:     true,
	Stateless:    true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &lengthTransform{}}, nil
	},
}
