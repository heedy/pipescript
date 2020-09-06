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
		cursum := float64(0)
		out.Timestamp = 0
		out.Duration = 0
		dp, _, err := e.Next(nil)
		if err != nil {
			return nil, err
		}

		if dp != nil {
			out.Timestamp = dp.Timestamp
		}
		ldp := dp

		for dp != nil {
			f, err := dp.Float()
			if err != nil {
				return nil, err
			}
			cursum += f

			dp, _, err = e.Next(nil)
			if err != nil {
				return nil, err
			}
			ldp = dp
		}

		out.Data = cursum
		if ldp != nil {
			out.Duration = ldp.Timestamp + ldp.Duration - out.Timestamp
		}

		return out, nil
	}),
}
