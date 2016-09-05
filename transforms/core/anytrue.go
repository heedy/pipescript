package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type anytrueTransform struct {
	istrue bool
}

func (t *anytrueTransform) Copy() (pipescript.TransformInstance, error) {
	return &anytrueTransform{t.istrue}, nil
}

func (t *anytrueTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.istrue = true
		return te.Get()
	}
	v, err := te.Datapoint.Bool()
	if err != nil {
		return nil, err
	}
	if v {
		t.istrue = true
	}
	return te.Set(t.istrue)
}

var AnyTrue = pipescript.Transform{
	Name:          "anytrue",
	Description:   "Returns true if at least one of the datapoints seen was true",
	Documentation: string(resources.MustAsset("docs/transforms/anytrue.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &anytrueTransform{false}}, nil
	},
}
