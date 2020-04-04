package interpolators

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/interpolator"
)

// MeanInterpolator sums the data in each interpolator
type MeanInterpolator struct {
	Iterator pipescript.DatapointIterator
	dp       *pipescript.Datapoint
}

// Interpolate performs the interpolation
func (i *MeanInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	sum := float64(0)
	count := 0
	for i.dp != nil && i.dp.Timestamp <= ts {
		d, err := i.dp.Float()
		if err != nil {
			return nil, err
		}
		sum += d
		count += 1

		i.dp, err = i.Iterator.Next()
		if err != nil {
			return nil, err
		}
	}
	if count == 0 {
		// Mean is 0 if no data
		return &pipescript.Datapoint{ts, 0.0}, nil
	}

	return &pipescript.Datapoint{ts, sum / float64(count)}, nil
}

var Mean = interpolator.Interpolator{
	Name:        "mean",
	Description: "Returns the mean of the datapoints in its interpolation time",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		dp, err := dpi.Next()
		return &MeanInterpolator{dpi, dp}, err
	},
}
