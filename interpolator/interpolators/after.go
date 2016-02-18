package interpolators

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// AfterInterpolator returns the closest datapoint /after/ the given timestamp
type AfterInterpolator struct {
	Iterator      pipescript.DatapointIterator // The DatapointIterator to use
	prevDatapoint *pipescript.Datapoint
}

// Interpolate performs interpolation
func (i *AfterInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	if i.prevDatapoint != nil && i.prevDatapoint.Timestamp > ts {
		// If the cached datapoint has greater timestamp, return that
		return i.prevDatapoint, nil
	}

	//We no longer care about prevDatapoint - get a datapoint that satisfies the constraint...
	//or return nil
	var err error
	i.prevDatapoint, err = i.Iterator.Next()
	for i.prevDatapoint != nil && err == nil && i.prevDatapoint.Timestamp <= ts {
		i.prevDatapoint, err = i.Iterator.Next()
	}
	return i.prevDatapoint, err
}

// NewAfterInterpolator returns the After interpolator for the given iterator
func NewAfterInterpolator(dpi pipescript.DatapointIterator) (*AfterInterpolator, error) {
	return &AfterInterpolator{Iterator: dpi}, nil
}

var after = interpolator.Interpolator{
	Name:        "after",
	Description: "Uses the closest datapoint after the interpolation timestamp",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		return NewAfterInterpolator(dpi)
	},
}
