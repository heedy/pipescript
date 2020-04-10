package core

import (
	"math/rand"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type randTransform struct{}

func (t randTransform) Copy() (pipescript.TransformInstance, error) {
	return randTransform{}, nil
}

func (t randTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	return te.Set(rand.Float64())
}

var Rand = pipescript.Transform{
	Name:          "rand",
	Description:   "Returns a random float in [0.0,1.0)",
	Documentation: string(resources.MustAsset("docs/transforms/rand.md")),
	OutputSchema:  `{"type": "number","minimum": 0, "exclusiveMaximum": 1}`,
	OneToOne:      true,
	// while it does not depend on datapoints, it is not stateless, since it does not return same result

	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &randTransform{}}, nil
	},
}
