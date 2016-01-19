package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestFirst(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "first",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, false},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, true},
			{5, false},
			{6, false},
		},
	}.Run(t)
}
