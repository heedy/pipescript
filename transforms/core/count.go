package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type countTransform struct {
	i     int64
	reset int64 // The value to use when resetting
}

func (t *countTransform) Copy() (pipescript.TransformInstance, error) {
	return &countTransform{t.i, t.reset}, nil
}

func (t *countTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.i = t.reset
		return te.Get()
	}
	t.i++
	return te.Set(t.i)
}

var Count = pipescript.Transform{
	Name:          "count",
	Description:   "Counts the number of datapoints that have been seen.",
	Documentation: string(resources.MustAsset("docs/transforms/count.md")),
	OutputSchema:  `{"type": "integer","minimum": 0}`,
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &countTransform{0, 0}}, nil
	},
}

var I = pipescript.Transform{
	Name:          "i",
	Description:   "Equivalent to the i in a loop over the sequence, starting from 0.",
	Documentation: "This transform returns 1 less than the `count` transform (`i` starts from 0, `count` from 1). Refer to the documentation for `count` for details.",
	OutputSchema:  Count.OutputSchema,
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &countTransform{-1, -1}}, nil
	},
}
