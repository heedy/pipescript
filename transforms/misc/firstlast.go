package misc

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type firstIter struct {
	isfirst bool
}

func (i *firstIter) OneToOne() bool {
	return true
}

func (i *firstIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	out.Data = i.isfirst
	i.isfirst = false

	return out, nil
}

var First = &pipescript.Transform{
	Name:          "first",
	Description:   "Returns true if first datapoint of a sequence, and false otherwise",
	Documentation: string(resources.MustAsset("docs/transforms/first.md")),
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return &firstIter{true}, nil
	},
}

type lastIter struct {
}

func (i lastIter) OneToOne() bool {
	return true
}

func (i lastIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	dnext, _, err := e.Peek(0, nil)
	if err != nil {
		return dp, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dnext == nil

	return out, nil
}

var Last = &pipescript.Transform{
	Name:          "last",
	Description:   "Returns true if last datapoint of a sequence, and false otherwise",
	Documentation: string(resources.MustAsset("docs/transforms/last.md")),
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return lastIter{}, nil
	},
}
