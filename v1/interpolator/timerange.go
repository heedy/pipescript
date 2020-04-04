package interpolator

import (
	"errors"

	"github.com/heedy/pipescript"
)

// The error returned by a timestamper when its stream is finished
var ErrEOF = errors.New("End of stream")

// TimeRange is an iterator which is used to get interpolation timestamps. Unlike pipescript's
// DatapointIterator, it returns ErrEOF when finished.
type TimeRange interface {
	Timestamp() (float64, error)
}

// uniformTimestamper returns tiemstamps by dt until it reaches end
type uniformRange struct {
	current float64 // The current timestamp
	dt      float64 // The amount to change timestamp with each step
	end     float64 // The end (exclusive)
}

// Timestamp gets the necessary timestamps
func (t *uniformRange) Timestamp() (float64, error) {
	if t.end <= t.current {
		return 0, ErrEOF
	}
	out := t.current
	t.current += t.dt
	return out, nil
}

// NewUniformRange returns an evenly spaced range of timestamps from t1 to t2 with spacing dt
func NewUniformRange(t1, t2, dt float64) (TimeRange, error) {
	if t1 >= t2 || dt <= 0 {
		return nil, errors.New("Invalid time range given")
	}

	return &uniformRange{t1, dt, t2}, nil
}

type iteratorRange struct {
	iter pipescript.DatapointIterator
}

// Timestamp returns the iterator's range
func (r *iteratorRange) Timestamp() (float64, error) {
	dp, err := r.iter.Next()

	if err != nil {
		return 0, err
	}
	if dp == nil {
		return 0, ErrEOF
	}

	return dp.Timestamp, nil
}

// NewIteratorRange creates a TimeRange based on a DatapointIterator. That is, it returns the timestamps
// of each consecutive datapoint.
func NewIteratorRange(dpi pipescript.DatapointIterator) (TimeRange, error) {
	return &iteratorRange{dpi}, nil
}
