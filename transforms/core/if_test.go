package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestIf(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "if",
		ParseError: true,
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "if($)",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
			{4, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{2, 8},
		},

		OutputError: true,
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "if $",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
			{4, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{2, 8},
		},

		OutputError: true,
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "if $ < 5",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 8},
			{3, "false"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{3, "false"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, "hi"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, 4},
		},
		OutputError: true,
	}.Run(t)

}
