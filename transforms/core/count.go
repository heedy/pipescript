package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type countTransform struct {
	i int64
}

func (t *countTransform) Copy() (pipescript.TransformInstance, error) {
	return &countTransform{t.i}, nil
}

func (t *countTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.i = 0
		return te.Get()
	}
	t.i++
	return te.Set(t.i)
}

var Count = pipescript.Transform{
	Name:          "count",
	Description:   "Counts the number of datapoints that have been seen. It is equivalent to the i in a loop over the sequence.",
	Documentation: string(resources.MustAsset("docs/transforms/count.md")),
	OutputSchema:  `{"type": "integer","minimum": 0}`,
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &countTransform{0}}, nil
	},
}
