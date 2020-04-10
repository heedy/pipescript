package misc

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Alltrue = &pipescript.Transform{
	Name:          "alltrue",
	Description:   "Returns true if all datapoints seen have been true, otherwise returns false",
	Documentation: string(resources.MustAsset("docs/transforms/alltrue.md")),
	InputSchema: map[string]interface{}{
		"type": "boolean",
	},
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		dp, _, err := e.Next(nil)
		if err != nil || dp == nil {
			return nil, err
		}
		ldp := dp
		out.Timestamp = dp.Timestamp
		for dp != nil {
			b, err := dp.Bool()
			if err != nil {
				return nil, err
			}
			if !b {
				out.Timestamp = dp.Timestamp
				out.Duration = dp.Duration
				out.Data = false
				return out, nil
			}

			ldp = dp
			dp, _, err = e.Next(nil)
			if err != nil {
				return nil, err
			}
		}

		out.Data = true
		out.Duration = ldp.Timestamp + ldp.Duration - out.Timestamp
		return out, nil
	}),
}
var Anytrue = pipescript.Transform{
	Name:          "anytrue",
	Description:   "Returns true if at least one of the datapoints seen was true",
	Documentation: string(resources.MustAsset("docs/transforms/anytrue.md")),
	InputSchema: map[string]interface{}{
		"type": "boolean",
	},
	OutputSchema: map[string]interface{}{
		"type": "boolean",
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		dp, _, err := e.Next(nil)
		if err != nil || dp == nil {
			return nil, err
		}
		ldp := dp
		out.Timestamp = dp.Timestamp
		for dp != nil {
			b, err := dp.Bool()
			if err != nil {
				return nil, err
			}

			ldp = dp
			if b {
				out.Timestamp = dp.Timestamp
				out.Duration = dp.Duration
				out.Data = true
				return out, nil
			}
			dp, _, err = e.Next(nil)
			if err != nil {
				return nil, err
			}
		}

		out.Data = false
		out.Duration = ldp.Timestamp + ldp.Duration - out.Timestamp
		return out, nil
	}),
}
