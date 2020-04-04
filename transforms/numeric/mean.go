package numeric

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

var Mean = &pipescript.Transform{
	Name:          "mean",
	Description:   "Finds the mean of the timeseries data",
	Documentation: string(resources.MustAsset("docs/transforms/mean.md")),
	Constructor: pipescript.NewAggregator(func(e *pipescript.TransformEnv, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		dp, _, err := e.Next(nil)
		if err != nil || dp == nil {
			return nil, err
		}
		ldp := dp
		out.Timestamp = dp.Timestamp
		cursum := float64(0)
		count := int64(0)
		for dp != nil {
			f, err := dp.Float()
			if err != nil {
				return nil, err
			}
			cursum += f
			count++
			ldp = dp
			dp, _, err = e.Next(nil)
			if err != nil {
				return nil, err
			}
		}

		out.Data = cursum / float64(count)
		out.Duration = ldp.Timestamp + ldp.Duration - out.Timestamp
		return out, nil
	}),
}
