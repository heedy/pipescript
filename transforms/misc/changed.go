package misc

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type changedIter struct{}

func (ci changedIter) OneToOne() bool {
	return true
}

func (ci changedIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, _, err := e.Next(nil)
	if err != nil || dp == nil {
		return dp, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	prev, _, err := e.Peek(-2, nil)
	if err != nil {
		if err == pipescript.ErrBeforeStart {
			// First datapoint is considered changed
			out.Data = true
			return out, nil
		}
	}

	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = !pipescript.Equal(dp.Data, prev.Data)

	return out, nil
}

var Changed = pipescript.Transform{
	Name:          "changed",
	Description:   "Returns true if the datapoint has a different value from the previous one",
	Documentation: string(resources.MustAsset("docs/transforms/changed.md")),
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return changedIter{}, nil
	},
}
