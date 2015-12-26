package transforms

import "github.com/connectordb/pipescript"

type lastTransform struct {
}

// Copy creates a copy of the last transform
func (lt *lastTransform) Copy() pipescript.TransformInstance {
	return &lastTransform{}
}

// Next returns the next element of the transform
func (lt *lastTransform) Next(dp pipescript.DatapointPeekIterator, args []*pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, err := dp.Next()
	if err != nil || dp == nil {
		return nil, err
	}
	// Peek at the next datapoint, to find out if it is nil (ie, the current datapoint is the last one)
	dp, err = dp.Peek(0)
	if err != nil {
		return nil, err
	}

	return &Datapoint{Timestamp: dp.Timestamp, Data: dp == nil}, nil
}

var last = pipescript.Transform{
	Name:         "last",
	Description:  "Returns true if last datapoint of a sequence, and false otherwise",
	OutputSchema: `{"type": "boolean"}`,
	OneToOne:     true,
	Generator: func(name string, args []*Datapoint) (TransformInstance, error) {
		return &lastTransform{}, nil
	},
}
