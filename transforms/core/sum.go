package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type sumTransform struct {
	cursum float64
}

func (t *sumTransform) Copy() (pipescript.TransformInstance, error) {
	return &sumTransform{t.cursum}, nil
}

func (t *sumTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.cursum = 0
		return te.Get()
	}
	v, err := te.Datapoint.Float()
	if err != nil {
		return nil, err
	}
	t.cursum += v
	return te.Set(t.cursum)
}

var Sum = pipescript.Transform{
	Name:          "sum",
	Description:   "Adds all of the values of the datapoints that pass through it",
	Documentation: string(resources.MustAsset("docs/transforms/sum.md")),
	OneToOne:      true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &sumTransform{0}}, nil
	},
}
