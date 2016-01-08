package pipescript

import "testing"

func TestIf(t *testing.T) {
	TestCase{
		Pipescript: "if",
		ParseError: true,
	}.Run(t)

	TestCase{
		Pipescript: "if($)",
		Input: []Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
			{4, "hi"},
		},
		Output: []Datapoint{
			{1, 1},
			{2, 8},
		},

		OutputError: true,
	}.Run(t)

	TestCase{
		Pipescript: "if $",
		Input: []Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
			{4, "hi"},
		},
		Output: []Datapoint{
			{1, 1},
			{2, 8},
		},

		OutputError: true,
	}.Run(t)

	TestCase{
		Pipescript: "if $ < 5",
		Input: []Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
		},
		Output: []Datapoint{
			{1, 1},
			{3, "false"},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, "hi"},
		},
		SecondaryOutput: []Datapoint{
			{4, 4},
		},
		OutputError: true,
	}.Run(t)

}
