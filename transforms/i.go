package transforms

import "github.com/connectordb/pipescript"

type iTransform struct {
	i int64
}

func (i *iTransform) Copy() pipescript.TransformInstance {
	return &iTransform{i.i}
}

func (i *iTransform) Next(dpi pipescript.DatapointPeekIterator, args []*pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, err := dpi.Next()
	if err != nil {
		return nil, err
	}
	if dp == nil {
		// If there is a nil, it means that the sequence is over - any future datapoints will belong to another sequence
		i.i = 0
		return nil, nil
	}

	dp = &pipescript.Datapoint{Timestamp: dp.Timestamp, Data: i.i}
	i.i++
	return dp, nil
}

var i = pipescript.Transform{
	Name:         "i",
	Description:  "Corresponds to the number of datapoints that have been seen. It is equivalent to the i in a loop over the sequence.",
	OutputSchema: `{"type": "integer","minimum": 0}`,
	OneToOne:     true,
	Generator: func(name string, args []*pipescript.Datapoint) (pipescript.TransformInstance, error) {
		return &iTransform{0}, nil
	},
}
