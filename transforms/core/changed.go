package core

import (
	"github.com/connectordb/duck"
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type changedTransform struct {
	Old interface{}
}

func (t *changedTransform) Copy() (pipescript.TransformInstance, error) {
	return &changedTransform{t.Old}, nil
}

func (t *changedTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.Old = nil
		return te.Get()
	}
	eq, ok := duck.Equal(t.Old, te.Datapoint.Data)
	t.Old = te.Datapoint.Data
	return te.Set(!(eq && ok))
}

var Changed = pipescript.Transform{
	Name:          "changed",
	Description:   "Returns true if the datapoint has a different value from the previous one",
	Documentation: string(resources.MustAsset("docs/transforms/changed.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &changedTransform{nil}}, nil
	},
}
