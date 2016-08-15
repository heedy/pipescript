package math

import (
	"math"

	"github.com/connectordb/pipescript"
)

type minTransform struct {
	mindp  *pipescript.Datapoint
	curmin float64
}

func (t *minTransform) Copy() (pipescript.TransformInstance, error) {
	return &minTransform{t.mindp, t.curmin}, nil
}

func (t *minTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		t.curmin = math.Inf(1)
		return te.Get()
	}
	v, err := te.Args[0].Float()
	if err != nil {
		return nil, err
	}
	if v < t.curmin {
		t.curmin = v
		t.mindp = te.Datapoint
	}
	return t.mindp, nil
}

var Min = pipescript.Transform{
	Name:        "min",
	Description: "Returns the minimum datapoint seen thus far",
	OneToOne:    true,
	Args: []pipescript.TransformArg{
		{
			Description: "Optional, if this is set, then the result of this argument is checked for min",
			Optional:    true,
			Default:     nil, // THIS HAS TO BE SET IN init.go to identity
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Args: args, Transform: &minTransform{nil, math.Inf(1)}}, nil
	},
}
