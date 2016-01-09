package core

import "github.com/connectordb/pipescript"

type countTransform struct {
	i int64
}

func (t *countTransform) Copy() (pipescript.TransformInstance, error) {
	return &countTransform{t.i}, nil
}

func (t *countTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.Datapoint == nil {
		// If there is a nil, it means that the sequence is over - any future datapoints will belong to another sequence
		t.i = 1
	}
	if te.IsFinished() {
		return te.Get()
	}

	dp, err := te.Set(t.i)
	t.i++
	return dp, err
}

var count = pipescript.Transform{
	Name:         "count",
	Description:  "Counts the number of datapoints that have been seen. It is equivalent to the i in a loop over the sequence.",
	OutputSchema: `{"type": "integer","minimum": 0}`,
	OneToOne:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &countTransform{1}}, nil
	},
}
