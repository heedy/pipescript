package pipescript

import "testing"

func TestIdentity(t *testing.T) {
	TestCase{
		Pipescript: "$",
		Input: []Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		Output: []Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		SecondaryInput: []Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []Datapoint{
			{4, 4},
			{5, 5},
		},
	}.Run(t)

	TestCase{
		Pipescript:  "$[1]",
		OutputError: true,
		Input: []Datapoint{
			{1, []string{"hi", "ho"}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
		},
		Output: []Datapoint{
			{1, "ho"},
			{2, "bar"},
		},
		SecondaryInput: []Datapoint{
			{4, map[string]string{"hello": "world", "1": "bar"}},
			{5, 5}, // This should give an error
		},
		SecondaryOutput: []Datapoint{
			{4, "bar"},
		},
	}.Run(t)

	TestCase{
		Pipescript: "$[$]", // The transform requries a constant
		ParseError: true,
	}.Run(t)

	TestCase{
		Pipescript:  "$ 1", // Bash-like usage
		OutputError: true,
		Input: []Datapoint{
			{1, []string{"hi", "ho"}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
		},
		Output: []Datapoint{
			{1, "ho"},
			{2, "bar"},
		},
		SecondaryInput: []Datapoint{
			{4, map[string]string{"hello": "world", "1": "bar"}},
			{5, 5}, // This should give an error
		},
		SecondaryOutput: []Datapoint{
			{4, "bar"},
		},
	}.Run(t)

}
