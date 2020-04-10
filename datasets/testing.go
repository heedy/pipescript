package datasets

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

// TestOutput is an output of the interpolator
type TestOutput struct {
	Timestamp float64 // The timestamp to use while interpolating
	Duration  float64 // Duration to use

	Output *pipescript.Datapoint // The expected output datapoint. nil if error expected (and last)
}

// TestCase is a simplified test case generator for use in testing Interpolators.
// Use this to make sure interpolators are doing the right thing.
type TestCase struct {
	Interpolator string // The interpolator to run
	Options      map[string]interface{}
	Reference    []pipescript.Datapoint
	Stream       []pipescript.Datapoint // The input datapoints
	GetError     bool                   // Whether there is to be a parsing error
	OutputError  bool                   // Whether there is to be an error during output
	Output       []pipescript.Datapoint // The output associated with the interpolator
}

// Run the test case
func (tc TestCase) Run(t *testing.T) {
	ref := pipescript.NewBuffer(pipescript.NewDatapointArrayIterator(tc.Reference))
	stream := pipescript.NewDatapointArrayIterator(tc.Stream)
	iter, err := GetInterpolator(tc.Interpolator, tc.Options, ref.Iterator(), stream)
	if tc.GetError {
		require.Error(t, err, "Parsing interpolator '%s' didn't give expected parser error", tc.Interpolator)
		return
	}
	require.NoError(t, err, "Failed to parse '%s'", tc.Interpolator)
	for i := range tc.Output {
		dp, err := iter.Next(&pipescript.Datapoint{})
		require.NoError(t, err, "Interpolator '%s' gave error on output %d", tc.Interpolator, i)
		require.EqualValues(t, &tc.Output[i], dp, "Interpolator: %s element %d (timestamp: %f)", tc.Interpolator, i, tc.Output[i].Timestamp)
	}
	// After finishing, we are supposed to get nil,nil
	v, err := iter.Next(&pipescript.Datapoint{})
	require.NoError(t, err, "Interpolator '%s' gave error on nil", tc.Interpolator)
	require.Nil(t, v, "Interpolator '%s' did not end with nil", tc.Interpolator)
}
