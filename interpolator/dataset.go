package interpolator

import (
	"fmt"

	"github.com/connectordb/pipescript"
)

// Dataset allows to perform interpolation on multiple interpolators at once
type Dataset map[string]InterpolatorInstance

func (d Dataset) Interpolate(timestamp float64) (dp *pipescript.Datapoint, err error) {
	res := make(map[string]*pipescript.Datapoint)
	for key := range d {
		res[key], err = d[key].Interpolate(timestamp)
		if err != nil {
			return nil, err
		}
	}
	return &pipescript.Datapoint{Timestamp: timestamp, Data: res}, nil
}

// This is a special interpolator + TimeRange at the same time made explicitly for use with GetXDataset
type timeRangeInterpolator struct {
	dpi pipescript.DatapointIterator

	dp *pipescript.Datapoint
}

// Timestamp performs the TimeRange portion of the interpolator - and sets the datapoint to return
// when called later as an interpolator
func (t *timeRangeInterpolator) Timestamp() (float64, error) {
	var err error
	t.dp, err = t.dpi.Next()
	if err != nil {
		return 0, err
	}
	if t.dp == nil {
		return 0, ErrEOF
	}
	return t.dp.Timestamp, nil
}

// Returns the datapoint whose timestamp was just returned in the TimeRange
func (t *timeRangeInterpolator) Interpolate(ts float64) (*pipescript.Datapoint, error) {
	return t.dp, nil
}

// GetXDataset generates an XDataset, which is a DatapointIterator generated using the given iterator's timestamps, the given Dataset,
// but also includes the iterator's datapoints under xname
func GetXDataset(xiterator pipescript.DatapointIterator, xname string, dataset Dataset) (pipescript.DatapointIterator, error) {
	_, ok := dataset[xname]
	if ok {
		return nil, fmt.Errorf("Dataset already contains field of name '%s'", xname)
	}

	// Set up the xiterator
	tri := &timeRangeInterpolator{dpi: xiterator}
	dataset[xname] = tri

	// Use the timestamps from the xiterator, and the datapoints also
	return NewInterpolationIterator(dataset, tri), nil
}

// GetTDataset returns a DatapointIterator representing the dataset with datapoints in the given range
func GetTDataset(t1, t2, dt float64, dataset Dataset) (pipescript.DatapointIterator, error) {
	tr, err := NewUniformRange(t1, t2, dt)
	if err != nil {
		return nil, err
	}
	return NewInterpolationIterator(dataset, tr), nil
}
