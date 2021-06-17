package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type whereIter struct {
	args []*pipescript.Datapoint
}

func (f whereIter) OneToOne() bool {
	return false
}

func (f whereIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
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

var Where = &pipescript.Transform{
	Name:          "where",
	Description:   "Filters all datapoints that do not pass the given conditional",
	Documentation: string(resources.MustAsset("docs/transforms/where.md")),
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
		return whereIter{make([]*pipescript.Datapoint, 1)}, nil
	},
}
