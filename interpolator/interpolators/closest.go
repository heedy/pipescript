package interpolators

import (
	"math"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// ClosestInterpolator interpolates an iterator by timestamp - getting the datapoint with the closest timestamp
type ClosestInterpolator struct {
	Iterator      pipescript.DatapointIterator
	prevDatapoint *pipescript.Datapoint
	curDatapoint  *pipescript.Datapoint
}

// Interpolate gets the datapoint corresponding to the interpolation timestamp
func (i *ClosestInterpolator) Interpolate(ts float64) (dp *pipescript.Datapoint, err error) {

	for i.curDatapoint != nil && i.curDatapoint.Timestamp <= ts {
		i.prevDatapoint = i.curDatapoint
		i.curDatapoint, err = i.Iterator.Next()
		if err != nil {
			return nil, err
		}
	}
	if i.prevDatapoint == nil {
		return i.curDatapoint, nil
	}
	if i.curDatapoint == nil {
		return i.prevDatapoint, nil
	}
	//Both prev and cur are not nil. Find which one is closer to ts
	if math.Abs(i.prevDatapoint.Timestamp-ts) <= math.Abs(i.curDatapoint.Timestamp-ts) {
		return i.prevDatapoint, nil
	}
	return i.curDatapoint, nil
}

// NewClosestInterpolator generates a new ClosestInterpolator given an Iterator.
func NewClosestInterpolator(dpi pipescript.DatapointIterator) (*ClosestInterpolator, error) {
	pd, err := dpi.Next()
	if err != nil {
		return nil, err
	}
	cd, err := dpi.Next()

	return &ClosestInterpolator{dpi, pd, cd}, err
}

var closest = interpolator.Interpolator{
	Name:        "closest",
	Description: "Uses the datapoint closest to the interpolation timestamp (both before and after)",
	Generator: func(name string, dpi pipescript.DatapointIterator) (i interpolator.InterpolatorInstance, err error) {
		return NewClosestInterpolator(dpi)
	},
}
