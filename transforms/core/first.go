package core

import "github.com/connectordb/pipescript"

type firstTransform struct {
	isfirst bool
}

// Copy creates a copy of the first transform
func (ft *firstTransform) Copy() pipescript.TransformInstance {
	return &firstTransform{ft.isfirst}
}

// Next returns the next element of the transform
func (ft *firstTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.Datapoint == nil {
		// If there is a nil, it means that the sequence is over - any future datapoints will belong to another sequence
		ft.isfirst = true
	}
	if te.IsFinished() {
		return te.Get()
	}

	dp, err := te.Set(ft.isfirst)
	ft.isfirst = false
	return dp, err
}

var first = pipescript.Transform{
	Name:         "first",
	Description:  "Returns true if first datapoint of a sequence, and false otherwise",
	OutputSchema: `{"type": "boolean"}`,
	OneToOne:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &firstTransform{true}}, nil
	},
}
