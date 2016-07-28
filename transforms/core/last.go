package core

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/resources"
)

type lastTransform struct{}

// Copy creates a copy of the last transform
func (lt lastTransform) Copy() (pipescript.TransformInstance, error) {
	return &lastTransform{}, nil
}

// Next returns the next element of the transform
func (lt lastTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	// Peek at the next datapoint, to find out if it is nil (ie, the current datapoint is the last one)
	te2 := ti.Peek(0)

	return te.Set(te2.IsFinished())
}

var Last = pipescript.Transform{
	Name:          "last",
	Description:   "Returns true if last datapoint of a sequence, and false otherwise",
	Documentation: string(resources.MustAsset("docs/transforms/last.md")),
	OutputSchema:  `{"type": "boolean"}`,
	OneToOne:      true,
	Peek:          true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: lastTransform{}}, nil
	},
}
