package math

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/transforms/core"
)

func TestMax(t *testing.T) {
	core.Length.Register()
	Register()
	pipescript.TestCase{
		Pipescript: "max",
		Input: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{3, 4},
		},
		Output: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{2, 5},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{3, 4},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{2, 5},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "max(length)",
		Input: []pipescript.Datapoint{
			{1, "a"},
			{2, "a b c"},
			{3, "a b"},
		},
		Output: []pipescript.Datapoint{
			{1, "a"},
			{2, "a b c"},
			{2, "a b c"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "a"},
			{2, "a b c"},
			{3, "a b"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, "a"},
			{2, "a b c"},
			{2, "a b c"},
		},
	}.Run(t)
}
