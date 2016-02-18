package interpolators

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// CountInterpolator counts the nubmer of datapoints that pass through in each period
type CountInterpolator struct {
	Iterator pipescript.DatapointIterator
	dp       *pipescript.Datapoint
}

// Interpolate performs the interpolation
func (i *CountInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	count := 0
	var err error
	for i.dp != nil && i.dp.Timestamp <= ts {
		count += 1

		i.dp, err = i.Iterator.Next()
		if err != nil {
			return nil, err
		}
	}

	return &pipescript.Datapoint{ts, count}, nil
}

var Count = interpolator.Interpolator{
	Name:        "count",
	Description: "Returns the number of datapoints in its interpolation time.",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		dp, err := dpi.Next()
		return &CountInterpolator{dpi, dp}, err
	},
}
