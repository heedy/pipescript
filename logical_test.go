package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAndTransform(t *testing.T) {
	// Now get error
	s, err := andScript(ConstantScript("pff"), ConstantScript("false"))
	require.NoError(t, err)
	_, err = s.GetConstant()
	require.Error(t, err)

	s, err = andScript(ConstantScript("false"), ConstantScript("pff"))
	require.NoError(t, err)
	_, err = s.GetConstant()
	require.Error(t, err)

	one := ConstantScript("true")
	one.OneToOne = false
	s, err = andScript(one, ConstantScript("false"))
	require.Error(t, err)
}

func TestOrTransform(t *testing.T) {
	// Now get error
	s, err := orScript(ConstantScript("pff"), ConstantScript("false"))
	require.NoError(t, err)
	_, err = s.GetConstant()
	require.Error(t, err)

	s, err = orScript(ConstantScript("false"), ConstantScript("pff"))
	require.NoError(t, err)
	_, err = s.GetConstant()
	require.Error(t, err)

	one := ConstantScript("true")
	one.OneToOne = false
	s, err = orScript(one, ConstantScript("false"))
	require.Error(t, err)
}

func TestLogicalConstant(t *testing.T) {
	ConstantTestCases{
		{"true and true", true},
		{"true and false", false},
		{"1.0 and true", true},
		{"1.0 and false", false},
		{"false and false", false},

		{"true or true", true},
		{"false or true", true},
		{"false or false", false},
		{"0 or true", true},
		{"0 or false", false},
	}.Run(t)

}

// These make sure that iterators pass through correctly
func TestLogical(t *testing.T) {
	TestCase{
		Pipescript: "true and false",
		Input: []Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []Datapoint{
			{1, false},
			{2, false},
			{3, false},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, false},
			{5, false},
		},
	}.Run(t)

	TestCase{
		Pipescript: "false or false",
		Input: []Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []Datapoint{
			{1, false},
			{2, false},
			{3, false},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, false},
			{5, false},
		},
	}.Run(t)
}
