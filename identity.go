package pipescript

import (
	"errors"

	"github.com/heedy/pipescript/resources"
)

type peekIterator struct {
	Peek int
}

func (i peekIterator) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	if i.Peek != 0 {
		dp, _, err = e.Peek(i.Peek-1, nil)
		if err != nil || dp == nil {
			return dp, err
		}
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Data

	return out, nil
}

func (i peekIterator) OneToOne() bool {
	return true
}

func basicGetElement(dp *Datapoint, args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error) {
	el, err := dp.MapElement(consts[0].(string))
	if err != nil {
		return nil, err
	}
	out.Data = el
	return out, nil

}

var Identity = &Transform{
	Name:          "$",
	Description:   `Represents the current datapoint. It is the identity transform.`,
	Documentation: string(resources.MustAsset("docs/transforms/$.md")),
	Args: []TransformArg{
		TransformArg{
			Description: "If it is given an integer, peeks to the given place in the array. If string, tries to get the element of an object",
			Type:        ConstArgType,
			Optional:    true,
			Default:     MustPipe(NewConstTransform(0), nil),
		},
	},
	Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		idx, ok := Int(consts[0])
		if ok {
			return peekIterator{int(idx)}, nil
		}

		// Otherwise, it is a string
		_, ok = consts[0].(string)
		if !ok {
			return nil, errors.New("Argument of invalid type")
		}
		return &Basic{
			ConstArgs: consts,
			PipeArgs:  pipes,
			Args:      nil,
			f:         basicGetElement,
		}, nil
	},
}

var IdentityPipe = MustPipe(Identity, []*Pipe{MustPipe(NewConstTransform(0), nil)})

type dtIter struct{}

func (t dtIter) OneToOne() bool {
	return true
}

func (t dtIter) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}

	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Duration

	return out, nil
}

var Dt = &Transform{
	Name:          "dt",
	Description:   "Gives access to the datapoint's duration",
	Documentation: string(resources.MustAsset("docs/transforms/dt.md")),
	Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		return dtIter{}, nil
	},
}

type tIter struct{}

func (t tIter) OneToOne() bool {
	return true
}

func (t tIter) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}

	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Timestamp

	return out, nil
}

var T = &Transform{
	Name:          "t",
	Description:   "Gives access to the datapoint's timestamp",
	Documentation: string(resources.MustAsset("docs/transforms/t.md")),
	Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		return tIter{}, nil
	},
}

func init() {
	Identity.Register()
	T.Register()
	Dt.Register()
}
