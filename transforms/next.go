package transforms

import (
	"errors"
	"fmt"

	"github.com/connectordb/pipescript"
)

// NextMax is the maximum number of datapoints forward to permit looking
var NextMax = int64(100)

type nextTransform struct {
	peekindex int
}

func (nt *nextTransform) Copy() pipescript.TransformInstance {
	return &nextTransform{nt.peekindex}
}

func (nt *nextTransform) Next(dpi pipescript.DatapointPeekIterator, args []*pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, err := dpi.Next()
	if err != nil || dp == nil {
		return nil, err
	}
	dp, err = dpi.Peek(nt.peekindex)
	if err != nil {
		return nil, err
	}
	if dp == nil {
		return &pipescript.Datapoint{Timestamp: dp.Timestamp, Data: nil}, nil
	}
	return dp.Copy(), nil
}

var next = pipescript.Transform{
	Name:        "next",
	Description: "Returns the datapoint that will be next in the sequence. If given an argument, can return the nth datapoint forward.",
	OneToOne:    true,
	Args: []pipescript.TransformArg{
		{
			Description: "The number of datapoints forward to look. Starts at 1.",
			Optional:    true,
			Default:     1,
			Constant:    true,
		},
	},

	Generator: func(name string, args []*pipescript.Datapoint) (pipescript.TransformInstance, error) {
		// The args array is guaranteed to be ordered according to the args. The Constant args
		// are guaranteed to have values already. So we are free to set things up directly
		i, err := args[0].Int()
		if err != nil {
			return nil, err
		}
		if i < 1 {
			return nil, errors.New("next must look at least one datapoint forward")
		}
		if i > NextMax {
			return nil, fmt.Errorf("next cannot look more than %d datapoints forward", NextMax)
		}

		// Looks like everything is valid
		return &nextTransform{int(i - 1)}, nil
	},
}
