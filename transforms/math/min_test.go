package math

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/transforms/core"
)

func TestMin(t *testing.T) {
	core.Length.Register()
	Register()
	pipescript.TestCase{
		Pipescript: "min",
		Input: []pipescript.Datapoint{
			{1, 3},
			{2, 1},
			{3, 2},
		},
		Output: []pipescript.Datapoint{
			{1, 3},
			{2, 1},
			{2, 1},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 3},
			{2, 1},
			{3, 2},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, 3},
			{2, 1},
			{2, 1},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "min(length)",
		Input: []pipescript.Datapoint{
			{1, "a b c"},
			{2, "a"},
			{3, "a b"},
		},
		Output: []pipescript.Datapoint{
			{1, "a b c"},
			{2, "a"},
			{2, "a"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "a b c"},
			{2, "a"},
			{3, "a b"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, "a b c"},
			{2, "a"},
			{2, "a"},
		},
	}.Run(t)
}
