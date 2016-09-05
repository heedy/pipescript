package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type alltrueTransform struct {
	istrue bool
}

func (t *alltrueTransform) Copy() (pipescript.TransformInstance, error) {
	return &alltrueTransform{t.istrue}, nil
}

func (t *alltrueTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.istrue = true
		return te.Get()
	}
	v, err := te.Datapoint.Bool()
	if err != nil {
		return nil, err
	}
	if !v {
		t.istrue = false
	}
	return te.Set(t.istrue)
}

var AllTrue = pipescript.Transform{
	Name:          "alltrue",
	Description:   "Returns true if all datapoints seen have been true",
	Documentation: string(resources.MustAsset("docs/transforms/alltrue.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &alltrueTransform{true}}, nil
	},
}
