package pipescript

import "testing"

func TestConstantComparison(t *testing.T) {
	ConstantTestCases{
		{"true==true", true},
		{"false==false", true},
		{"false!=false", false},
		{"false<true", true},
		{"'hello'=='hello'", true},
		{"'hello'!='hello'", false},
		{"'hello'=='Hello'", false},
		{"'hello'!='Hello'", true},
		{"5==5", true},
		{"6 ==5", false},
		{"6!=5", true},
		{" 5 != 5 ", false},
		{"5 < 5", false},
		{"5 < 6", true},
		{"20 < 135", true},
		{"5 > 5", false},
		{"5 > 6", false},
		{"200 > 135", true},
		{"5.3 <= 5.3", true},
		{"5.3 >= 5.3", true},
		{"5.2 <= 5.3", true},
		{"5.2 >= 5.3", false},
		{"5.3 <= 5", false},
		{"5.3 >= 5", true},
	}.Run(t)
}

func TestComparison(t *testing.T) {
	TestCase{
		Pipescript: "true==false",
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
		Pipescript: "6 < 5",
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
		Pipescript: "true!=true",
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
		Pipescript: "5 > 6",
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
		Pipescript: "true!=true",
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
		Pipescript: "5 > 6",
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
		Pipescript: "5 >= 6",
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
		Pipescript: "6 <= 5",
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
