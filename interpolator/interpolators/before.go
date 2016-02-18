package interpolators

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// BeforeInterpolator returns the closest datapoint /before (or equal) to the given timestamp
type BeforeInterpolator struct {
	Iterator pipescript.DatapointIterator

	prevDatapoint *pipescript.Datapoint
	curDatapoint  *pipescript.Datapoint
}

// Interpolate performs interpolation
func (i *BeforeInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	var err error
	for i.curDatapoint != nil && i.curDatapoint.Timestamp <= ts {
		i.prevDatapoint = i.curDatapoint
		i.curDatapoint, err = i.Iterator.Next()
		if err != nil {
			return nil, err
		}
	}
	if i.prevDatapoint != nil && i.prevDatapoint.Timestamp > ts {
		return nil, nil
	}
	return i.prevDatapoint, nil
}

// NewBeforeInterpolator creates a new BeforeInterpolator
func NewBeforeInterpolator(dpi pipescript.DatapointIterator) (*BeforeInterpolator, error) {
	pd, err := dpi.Next()
	if err != nil {
		return nil, err
	}
	cd, err := dpi.Next()
	return &BeforeInterpolator{dpi, pd, cd}, err
}

var Before = interpolator.Interpolator{
	Name:        "before",
	Description: "Uses the closest datapoint before the interpolation timestamp",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		return NewBeforeInterpolator(dpi)
	},
}
