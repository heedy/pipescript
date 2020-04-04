package numeric

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Count = &pipescript.Transform{
	Name:          "count",
	Description:   "returns the total number of datapoints in the stream",
	Documentation: string(resources.MustAsset("docs/transforms/count.md")),
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		count := int64(0)
		out.Timestamp = 0
		out.Duration = 0
		dp, _, err := e.Next(nil)
		if err != nil {
			return dp, err
		}
		if dp != nil {
			out.Timestamp = dp.Timestamp
		}
		dp2 := dp

		for dp != nil {
			count++
			dp2 = dp
			dp, _, err = e.Next(nil)
			if err != nil {
				return dp, err
			}
		}
		if dp2 != nil {
			out.Duration = dp2.Timestamp + dp2.Duration - out.Timestamp
		}
		out.Data = count
		return out, nil
	}),
}
