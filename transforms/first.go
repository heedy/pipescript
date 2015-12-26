package transforms

import "github.com/connectordb/pipescript"

type firstTransform struct {
	isfirst bool
}

// Copy creates a copy of the first transform
func (ft *firstTransform) Copy() pipescript.TransformInstance {
	return &firstTransform{ft.isfirst}
}

// Next returns the next element of the transform
func (ft *firstTransform) Next(dp pipescript.DatapointPeekIterator, args []*pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, err := dp.Next()
	if err != nil {
		return nil, err
	}
	if dp == nil {
		// If there is a nil, it means that the sequence is over - any future datapoints will belong to another sequence
		ft.isfirst = true
	}

	dp := &Datapoint{Timestamp: dp.Timestamp, Data: ft.isfirst}
	ft.isfirst = false
	return dp, nil
}

var first = pipescript.Transform{
	Name:         "first",
	Description:  "Returns true if first datapoint of a sequence, and false otherwise",
	OutputSchema: `{"type": "boolean"}`,
	OneToOne:     true,
	Generator: func(name string, args []*Datapoint) (TransformInstance, error) {
		return &firstTransform{true}, nil
	},
}
