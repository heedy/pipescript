// +build !js

package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ConstantTestCase is a simplified test case for quick checking of properties that do not need the big machinery
// of a full TestCase. It must be used on Scripts that are constant.
type ConstantTestCase struct {
	Pipescript string
	Output     interface{}
}

// Run the ConstantTestCase
func (tc ConstantTestCase) Run(t *testing.T) {
	s, err := Parse(tc.Pipescript)
	require.NoError(t, err, "Failed to parse '%s'", tc.Pipescript)
	s2, err := s.Copy()
	require.NoError(t, err, "Failed to copy '%s'", tc.Pipescript)
	bv, err := s.GetConstant()
	require.NoError(t, err, "Failed to get constant from '%s'", tc.Pipescript)
	require.NotNil(t, bv, "Failed to get constant from '%s'", tc.Pipescript)
	require.Equal(t, tc.Output, bv.Data, tc.Pipescript)
	bv, err = s2.GetConstant()
	require.NoError(t, err, "Failed to get constant from copied '%s'", tc.Pipescript)
	require.NotNil(t, bv, "Failed to get constant from copied '%s'", tc.Pipescript)
	require.Equal(t, tc.Output, bv.Data, tc.Pipescript)
}

// ConstantTestCases is what is used to test - it allows quickly running multiple tests
type ConstantTestCases []ConstantTestCase

// Run the full array of ConstantTestCases
func (tc ConstantTestCases) Run(t *testing.T) {
	for i := range tc {
		tc[i].Run(t)
	}
}

// TestCase is a simplified test case generator for use in testing transforms.
// Use it to make sure your transforms are doing the right thing
type TestCase struct {
	Pipescript      string      // The script to run
	Input           []Datapoint // The input Datapoints
	ParseError      bool        // Whether there is to be a parsing error
	OutputError     bool        // Whether there is to be an error during output
	Output          []Datapoint // The output Datapoints
	SecondaryInput  []Datapoint // Input to attempt after first input passes through
	SecondaryOutput []Datapoint // The output expected from the SecondaryInput
}

// Run the test case
func (tc TestCase) Run(t *testing.T) {
	s, err := Parse(tc.Pipescript)
	if tc.ParseError {
		require.Error(t, err, "Parsing script '%s' didn't give expected parser error", tc.Pipescript)
		return
	}
	require.NoError(t, err, "Failed to parse '%s'", tc.Pipescript)
	s2, err := s.Copy()
	require.NoError(t, err, "Failed to copy script '%s'", tc.Pipescript)

	dpi := NewDatapointArrayIterator(tc.Input)
	dpi2 := NewDatapointArrayIterator(tc.Input)

	s.SetInput(dpi)
	s2.SetInput(dpi2)

	for i := range tc.Output {
		v, err := s.Next()
		require.NoError(t, err, "Script '%s' gave error on output %d", tc.Pipescript, i)
		v2, err := s2.Next()
		require.NoError(t, err, "Copied script '%s' gave error on output %d", tc.Pipescript, i)

		require.EqualValues(t, &tc.Output[i], v, "Script '%s'", tc.Pipescript)
		require.EqualValues(t, &tc.Output[i], v2, "Copied script '%s'", tc.Pipescript)
	}

	if tc.OutputError && tc.SecondaryInput == nil { // If there is supposed to be output error, and there is no secondary
		_, err := s.Next()
		require.Error(t, err, "Script '%s' did not give error", tc.Pipescript)
		_, err = s2.Next()
		require.Error(t, err, "Copied script '%s' did not give error", tc.Pipescript)
		return
	}

	// After finishing, we are supposed to get nil,nil
	v, err := s.Next()
	require.NoError(t, err, "Script '%s' gave error on nil", tc.Pipescript)
	require.Nil(t, v, "Script '%s' did not end with nil", tc.Pipescript)
	v2, err := s2.Next()
	require.NoError(t, err, "Copied script '%s' gave error on nil", tc.Pipescript)
	require.Nil(t, v2, "Copied script '%s' did not end with nil", tc.Pipescript)

	// Now run the transform on secondary output
	if tc.SecondaryInput != nil {
		dpi = NewDatapointArrayIterator(tc.SecondaryInput)
		dpi2 = NewDatapointArrayIterator(tc.SecondaryInput)

		s.SetInput(dpi)
		s2.SetInput(dpi2)

		for i := range tc.SecondaryOutput {
			v, err := s.Next()
			require.NoError(t, err, "Script '%s' gave error on output %d in secondary output %d", tc.Pipescript, i)
			v2, err := s2.Next()
			require.NoError(t, err, "Copied script '%s' gave error on output %d in secondary output %d", tc.Pipescript, i)

			require.EqualValues(t, &tc.SecondaryOutput[i], v, "Script '%s' in secondary output %d", tc.Pipescript, i)
			require.EqualValues(t, &tc.SecondaryOutput[i], v2, "Copied script '%s' in secondary output %d", tc.Pipescript, i)
		}

		if tc.OutputError {
			_, err := s.Next()
			require.Error(t, err, "Script '%s' did not give error on secondary output", tc.Pipescript)
			_, err = s2.Next()
			require.Error(t, err, "Copied script '%s' did not give error on secondary output", tc.Pipescript)
			return
		}

		// After finishing, we are supposed to get nil,nil
		v, err = s.Next()
		require.NoError(t, err, "Script '%s' gave error on nil in secondary output", tc.Pipescript)
		require.Nil(t, v, "Script '%s' did not end with nil in secondary output", tc.Pipescript)
		v2, err = s2.Next()
		require.NoError(t, err, "Copied script '%s' gave error on nil in secondary output", tc.Pipescript)
		require.Nil(t, v2, "Copied script '%s' did not end with nil in secondary output", tc.Pipescript)
	}
}
