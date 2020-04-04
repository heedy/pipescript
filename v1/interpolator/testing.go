// +build !js

package interpolator

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

// TestOutput is an output of the interpolator
type TestOutput struct {
	InterpolationTimestamp float64               // The timestamp to use while interpolating
	Output                 *pipescript.Datapoint // The expected output datapoint. nil if error expected (and last)
}

// TestCase is a simplified test case generator for use in testing Interpolators.
// Use this to make sure interpolators are doing the right thing.
type TestCase struct {
	Interpolator string                 // The interpolator to run
	Input        []pipescript.Datapoint // The input datapoints
	ParseError   bool                   // Whether there is to be a parsing error
	OutputError  bool                   // Whether there is to be an error during output
	Output       []TestOutput           // The output associated with the interpolator
}

// Run the test case
func (tc TestCase) Run(t *testing.T) {
	dpi := pipescript.NewDatapointArrayIterator(tc.Input)
	ipltr, err := Parse(tc.Interpolator, dpi)
	if tc.ParseError {
		require.Error(t, err, "Parsing interpolator '%s' didn't give expected parser error", tc.Interpolator)
		return
	}
	require.NoError(t, err, "Failed to parse '%s'", tc.Interpolator)
	for i := range tc.Output {
		dp, err := ipltr.Interpolate(tc.Output[i].InterpolationTimestamp)

		if len(tc.Output)-1 == i && tc.OutputError {
			require.Error(t, err, "Interpolator: %s OutputError not given", tc.Interpolator)
			return
		}
		require.NoError(t, err, "Interpolator: %s", tc.Interpolator)
		require.EqualValues(t, tc.Output[i].Output, dp, "Interpolator: %s element %d (timestamp: %f)", tc.Interpolator, i, tc.Output[i].InterpolationTimestamp)
	}
}
