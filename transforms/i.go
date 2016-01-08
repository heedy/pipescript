package transforms

import "github.com/connectordb/pipescript"

type iTransform struct {
	i int64
}

func (i *iTransform) Copy() pipescript.TransformInstance {
	return &iTransform{i.i}
}

func (i *iTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.Datapoint == nil {
		// If there is a nil, it means that the sequence is over - any future datapoints will belong to another sequence
		i.i = 0
	}
	if te.IsFinished() {
		return te.Get()
	}

	dp, err := te.Set(i.i)
	i.i++
	return dp, err
}

var i = pipescript.Transform{
	Name:         "i",
	Description:  "Corresponds to the number of datapoints that have been seen. It is equivalent to the i in a loop over the sequence.",
	OutputSchema: `{"type": "integer","minimum": 0}`,
	OneToOne:     true,
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		return &pipescript.TransformInitializer{Transform: &iTransform{0}}, nil
	},
}
