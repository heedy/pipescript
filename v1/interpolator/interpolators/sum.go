package interpolators

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/interpolator"
)

// SumInterpolator sums the data in each interpolator
type SumInterpolator struct {
	Iterator pipescript.DatapointIterator
	dp       *pipescript.Datapoint
}

// Interpolate performs the interpolation
func (i *SumInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	sum := float64(0)
	for i.dp != nil && i.dp.Timestamp <= ts {
		d, err := i.dp.Float()
		if err != nil {
			return nil, err
		}
		sum += d

		i.dp, err = i.Iterator.Next()
		if err != nil {
			return nil, err
		}
	}

	return &pipescript.Datapoint{ts, sum}, nil
}

var Sum = interpolator.Interpolator{
	Name:        "sum",
	Description: "Returns the sum of datapoints in its interpolation time.",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		dp, err := dpi.Next()
		return &SumInterpolator{dpi, dp}, err
	},
}
