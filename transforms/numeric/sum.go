package numeric

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Sum = &pipescript.Transform{
	Name:          "sum",
	Description:   "Sums data",
	Documentation: string(resources.MustAsset("docs/transforms/sum.md")),
	InputSchema: map[string]interface{}{
		"type": "number",
	},
	OutputSchema: map[string]interface{}{
		"type": "number",
	},
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		dp, _, err := e.Next(nil)
		if err != nil || dp == nil {
			return nil, err
		}
		ldp := dp
		out.Timestamp = dp.Timestamp
		cursum := float64(0)
		for dp != nil {
			f, err := dp.Float()
			if err != nil {
				return nil, err
			}
			cursum += f
			ldp = dp
			dp, _, err = e.Next(nil)
			if err != nil {
				return nil, err
			}
		}

		out.Data = cursum
		out.Duration = ldp.Timestamp + ldp.Duration - out.Timestamp
		return out, nil
	}),
}
