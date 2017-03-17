package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type dtTransform struct {
	prev float64
}

func (t *dtTransform) Copy() (pipescript.TransformInstance, error) {
	return &dtTransform{t.prev}, nil
}

func (t *dtTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.prev = 0
		return te.Get()
	}
	prev := t.prev
	t.prev = te.Datapoint.Timestamp
	if prev == 0 {
		return te.Set(float64(0))
	}
	return te.Set(te.Datapoint.Timestamp - prev)
}

var Dt = pipescript.Transform{
	Name:          "dt",
	Description:   "Returns time difference between this and previous datapoint",
	Documentation: string(resources.MustAsset("docs/transforms/dt.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &dtTransform{0}}, nil
	},
}
