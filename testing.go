package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Pipescript  string      // The script to run
	Input       []Datapoint // The input datapoints
	Parsed      string      // The output of the parsed pipe.String(). "" if there is to be an error
	OutputError bool        // Whether there is to be an error during output
	Output      []Datapoint // The output Datapoints
}

func (tc TestCase) Run(t *testing.T) {
	p, err := Parse(tc.Pipescript)
	if tc.Parsed == "error" {
		require.Error(t, err, "Parsing script '%s' didn't give expected parser error", tc.Pipescript)
		return
	}
	require.NoError(t, err, "Parsing script '%s' failed", tc.Pipescript)
	if tc.Parsed != "" {
		require.Equal(t, tc.Parsed, p.String())
	}

	p2 := p.Copy()

	dpi := NewDatapointArrayIterator(tc.Input)
	dpi2 := NewDatapointArrayIterator(tc.Input)

	p.InputIterator(dpi)
	p2.InputIterator(dpi2)

	for i := range tc.Output {
		v, err := p.Next(&Datapoint{})
		require.NoError(t, err, "Script '%s' (%s) gave error on output %d", tc.Pipescript, p.String(), i)
		v2, err := p2.Next(&Datapoint{})
		require.NoError(t, err, "Copied script '%s' (%s) gave error on output %d", tc.Pipescript, p.String(), i)

		require.EqualValues(t, &tc.Output[i], v, "Script '%s' (%s)", tc.Pipescript, p.String())
		require.EqualValues(t, &tc.Output[i], v2, "Copied script '%s' (%s)", tc.Pipescript, p.String())
	}

	if tc.OutputError { // If there is supposed to be output error, and there is no secondary
		_, err := p.Next(&Datapoint{})
		require.Error(t, err, "Script '%s' (%s) did not give error", tc.Pipescript, p.String())
		_, err = p2.Next(&Datapoint{})
		require.Error(t, err, "Copied script '%s' (%s) did not give error", tc.Pipescript, p.String())
		return
	}

	// After finishing, we are supposed to get nil,nil
	v, err := p.Next(&Datapoint{})
	require.NoError(t, err, "Script '%s' (%s) gave error on nil", tc.Pipescript, p.String())
	require.Nil(t, v, "Script '%s' (%s) did not end with nil", tc.Pipescript, p.String())
	v2, err := p2.Next(&Datapoint{})
	require.NoError(t, err, "Copied script '%s' (%s) gave error on nil", tc.Pipescript, p.String())
	require.Nil(t, v2, "Copied script '%s' (%s) did not end with nil", tc.Pipescript, p.String())
}
