package interpolator

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

// TestCase is a simplified test case generator for use in testing Interpolators.
// Use this to make sure interpolators are doing the right thing.
type TestCase struct {
	Interpolator      string                  // The interpolator to run
	Input             []pipescript.Datapoint  // The input datapoints
	ParseError        bool                    // Whether there is to be a parsing error
	OutputError       bool                    // Whether there is to be an error during output
	InterpolationTime []float64               // The timestamps to use for interpolation
	Output            []*pipescript.Datapoint // The output associated with the input timestamps
}

// Run the test case
func (tc TestCase) Run(t *testing.T) {
	if len(tc.InterpolationTime) != len(tc.Output) && !tc.OutputError {
		t.Error("The interpolation time array must be same size as Output array for test case to run!")
		return
	}
	if len(tc.InterpolationTime) != len(tc.Output)+1 && tc.OutputError {
		t.Error("The interpolation time array must be one larger than the Output array when OutputError is set to true in test case")
	}
	dpi := pipescript.NewDatapointArrayIterator(tc.Input)
	ipltr, err := Parse(tc.Interpolator, dpi)
	if tc.ParseError {
		require.Error(t, err, "Parsing interpolator '%s' didn't give expected parser error", tc.Interpolator)
		return
	}
	require.NoError(t, err, "Failed to parse '%s'", tc.Interpolator)
	for i := range tc.Output {
		dp, err := ipltr.Next(tc.InterpolationTime[i])
		require.NoError(t, err, "Interpolator: %s", tc.Interpolator)
		require.EqualValues(t, tc.Output[i], dp, "Interpolator: %s element %d (timestamp: %f)", tc.Interpolator, i, tc.InterpolationTime[i])
	}

	if tc.OutputError {
		_, err := ipltr.Next(tc.InterpolationTime[len(tc.InterpolationTime)-1])
		require.Error(t, err, "Interpolator: %s OutputError not given", tc.Interpolator)
	}
}
