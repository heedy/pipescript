package core

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

func (nt *nextTransform) Copy() (pipescript.TransformInstance, error) {
	return &nextTransform{nt.peekindex}, nil
}

func (nt *nextTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	te2 := ti.Peek(nt.peekindex)
	if te2.IsFinished() {
		return te.Set(nil)
	}
	return te2.Datapoint, nil
}

var Next = pipescript.Transform{
	Name:        "next",
	Description: "Returns the datapoint that will be next in the sequence. If given an argument, can return the nth datapoint forward.",
	OneToOne:    true,
	Peek:        true,
	Args: []pipescript.TransformArg{
		{
			Description: "The number of datapoints forward to look. Starts at 1.",
			Optional:    true,
			Default:     1,
			Constant:    true,
		},
	},

	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		// The args array is guaranteed to be ordered according to the args. The Constant args
		// are guaranteed to have values already. So we are free to set things up directly
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		i, err := dp.Int()
		if err != nil {
			return nil, err
		}
		if i < 1 {
			return nil, errors.New("next must look at least one datapoint forward")
		}
		if i > NextMax {
			return nil, fmt.Errorf("next cannot look more than %d datapoints forward", NextMax)
		}

		// Looks like everything is valid - remove the constant arg from consideration
		return &pipescript.TransformInitializer{Transform: &nextTransform{int(i - 1)}}, nil
	},
}
