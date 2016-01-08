package pipescript

import "testing"

func TestConstantAlgebra(t *testing.T) {
	ConstantTestCases{
		{"5+5", float64(10)},
		{"'hello'+'world'", "helloworld"},
		{"23 - 3", float64(20)},
		{"23-3", float64(20)}, // This caused error since lexer checked negative numbers
		{"5*3", float64(15)},
		{"15/3", float64(5)},
		{"4%2", float64(0)},
		{"5%2", float64(1)},
		{"3^2", float64(9)},
	}.Run(t)
}

func TestAlgebra(t *testing.T) {
	TestCase{
		Pipescript: "5+3",
		Input: []Datapoint{
			{1, 1},
		},
		Output: []Datapoint{
			{1, float64(8)},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, float64(8)},
			{5, float64(8)},
		},
	}.Run(t)

	TestCase{
		Pipescript: "14-6",
		Input: []Datapoint{
			{1, 1},
		},
		Output: []Datapoint{
			{1, float64(8)},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, float64(8)},
			{5, float64(8)},
		},
	}.Run(t)

	TestCase{
		Pipescript: "40/5",
		Input: []Datapoint{
			{1, 1},
		},
		Output: []Datapoint{
			{1, float64(8)},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, float64(8)},
			{5, float64(8)},
		},
	}.Run(t)

	TestCase{
		Pipescript: "28%20",
		Input: []Datapoint{
			{1, 1},
		},
		Output: []Datapoint{
			{1, float64(8)},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, float64(8)},
			{5, float64(8)},
		},
	}.Run(t)

	TestCase{
		Pipescript: "2^3",
		Input: []Datapoint{
			{1, 1},
		},
		Output: []Datapoint{
			{1, float64(8)},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, float64(8)},
			{5, float64(8)},
		},
	}.Run(t)
}
