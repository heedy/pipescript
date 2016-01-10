package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestChanged(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "changed",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 1},
			{3, 2},
		},
		Output: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, true},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 2},
			{5, 5},
			{6, 5},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, true},
			{5, true},
			{6, false},
		},
	}.Run(t)
}
