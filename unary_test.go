package pipescript

import "testing"

func TestUnaryConstant(t *testing.T) {
	ConstantTestCases{
		{"not true", false},
		{"not false", true},
		{"-true", false},
		{"-false", true},
		{"- 50", float64(-50)},
		{"- '0.54'", float64(-0.54)},
	}.Run(t)
}

// These tests make sure it behaves correctly in normal usage
func TestUnary(t *testing.T) {
	TestCase{
		Pipescript: "not true",
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
		Pipescript: "-true",
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
