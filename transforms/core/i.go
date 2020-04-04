package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type iIter struct {
	i int64
}

func (i *iIter) OneToOne() bool {
	return true
}

func (i *iIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = i.i
	i.i++
	return out, nil
}

var I = &pipescript.Transform{
	Name:          "i",
	Description:   "Gives array index of the timeseries",
	Documentation: string(resources.MustAsset("docs/transforms/i.md")),
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return &iIter{}, nil
	},
}
