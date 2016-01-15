package interpolator

import (
	"errors"

	"github.com/connectordb/pipescript"
)

// iflast is the script corresponding to "if last"
var iflast *pipescript.Script

// timestampIterator is a DatapointIterator which manually ends at the given Timestamp.
// this allows pretending the iterator finished for the script.
type timestampIterator struct {
	Timestamp     float64
	dpi           pipescript.DatapointIterator
	nextDatapoint *pipescript.Datapoint
}

func (t *timestampIterator) Next() (dp *pipescript.Datapoint, err error) {
	if t.nextDatapoint == nil {
		return nil, nil
	}
	// If this timestamp is within range, return the associated datapoint
	if t.nextDatapoint.Timestamp <= t.Timestamp {
		dp = t.nextDatapoint
		t.nextDatapoint, err = t.dpi.Next()
		return dp, err
	}

	// Otherwise, pretend we're done
	return nil, nil
}

// newTimestampIterator sets up the timestampIterator
func newTimestampIterator(dpi pipescript.DatapointIterator) (*timestampIterator, error) {
	dp, err := dpi.Next()
	if err != nil {
		return nil, err
	}
	return &timestampIterator{
		dpi:           dpi,
		nextDatapoint: dp,
	}, nil
}

// ScriptInterpolator can turn any pipescript.Script into an interpolator. For each interpolation step it effectively
// performs a `{script} | if last` and returns the resulting datapoint. This allows you to sum/count/whatever in the
// interpolation steps. The script is effectively wiped at each step
type ScriptInterpolator struct {
	s      *pipescript.Script
	tsi    *timestampIterator
	nilans *pipescript.Datapoint
}

// NewScriptInterpolator generates an interpolator based upon a given script and the given DatapointIterator.
// Sometimes the script might return nil as an answer for a certain time period. When that happens, the ScriptIterator will
// return your nilanswer datapoint. Leave it nil if you don't want any special handling
func NewScriptInterpolator(s *pipescript.Script, dpi pipescript.DatapointIterator, nilanswer *pipescript.Datapoint) (*ScriptInterpolator, error) {

	// Add the "if last" to the end of the script
	newiflast, err := iflast.Copy()
	if err != nil {
		return nil, err
	}

	// Combine the script with our if last
	err = s.Append(newiflast)
	if err != nil {
		return nil, err
	}

	tsi, err := newTimestampIterator(dpi)
	s.SetInput(tsi)
	return &ScriptInterpolator{s, tsi, nilanswer}, err
}

// Interpolate allows you to get the interpolated datapoint
func (s *ScriptInterpolator) Interpolate(timestamp float64) (*pipescript.Datapoint, error) {
	// Clear out the script (it might have some remaining values in pipeline from previous run)
	dp, err := s.s.Next()
	if err != nil {
		return nil, err
	}
	if dp != nil {
		return nil, errors.New("Failed to clear PipeScript used in interpolator")
	}

	// Set the timestamp
	s.tsi.Timestamp = timestamp

	// Get the datapoint
	dp, err = s.s.Next()
	if dp == nil && s.nilans != nil {
		s.nilans.Timestamp = timestamp
		return s.nilans, err
	}
	return dp, err
}
