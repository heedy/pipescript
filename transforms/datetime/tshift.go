package datetime

import (
	"errors"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type tshiftIterator struct {
	shiftby float64
}

func (ti tshiftIterator) OneToOne() bool {
	return true
}

func (ti tshiftIterator) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	out.Timestamp = dp.Timestamp + ti.shiftby
	out.Duration = dp.Duration
	out.Data = dp.Data

	return out, nil
}

var Tshift = pipescript.Transform{
	Name:          "tshift",
	Description:   "Shift the datapoint timestamp by a constant number of seconds",
	Documentation: string(resources.MustAsset("docs/transforms/tshift.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "The number of seconds to shift the timestamp",
			Type:        pipescript.ConstArgType,
		},
	},
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		shiftby, ok := pipescript.Float(consts[0])
		if !ok {
			return nil, errors.New("Must shift by a number")
		}
		return tshiftIterator{shiftby}, nil
	},
}
