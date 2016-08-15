package math

import (
	"math"

	"github.com/connectordb/pipescript"
)

type maxTransform struct {
	maxdp  *pipescript.Datapoint
	curmax float64
}

func (t *maxTransform) Copy() (pipescript.TransformInstance, error) {
	return &maxTransform{t.maxdp, t.curmax}, nil
}

func (t *maxTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.curmax = math.Inf(-1)
		return te.Get()
	}
	v, err := te.Args[0].Float()
	if err != nil {
		return nil, err
	}
	if v > t.curmax {
		t.curmax = v
		t.maxdp = te.Datapoint
	}
	return t.maxdp, nil
}

var Max = pipescript.Transform{
	Name:        "max",
	Description: "Returns the maximum datapoint seen thus far",
	OneToOne:    true,
	Args: []pipescript.TransformArg{
		{
			Description: "Optional, if this is set, then the result of this argument is checked for max",
			Optional:    true,
			Default:     nil, // THIS HAS TO BE SET IN init.go to identity
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Args: args, Transform: &maxTransform{nil, math.Inf(-1)}}, nil
	},
}
