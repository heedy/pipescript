package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type filterIter struct {
	args []*pipescript.Datapoint
}

func (f filterIter) OneToOne() bool {
	return false
}

func (f filterIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	for {
		dp, args, err := e.Next(f.args)
		if dp == nil || err != nil {
			return dp, nil
		}
		argval, err := args[0].Bool()
		if err != nil {
			return nil, err
		}
		if argval {
			out.Timestamp = dp.Timestamp
			out.Duration = dp.Duration
			out.Data = dp.Data
			return out, nil
		}
	}
}

var Filter = &pipescript.Transform{
	Name:          "filter",
	Description:   "Filters all datapoints that do not pass the given conditional",
	Documentation: string(resources.MustAsset("docs/transforms/filter.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "Statement to check for truth value",
			Type:        pipescript.TransformArgType,
			Schema: map[string]interface{}{
				"type": "boolean",
			},
		},
	},
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return filterIter{make([]*pipescript.Datapoint, 1)}, nil
	},
}
