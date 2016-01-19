package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestLast(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "if last",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 2},
		},
		Output: []pipescript.Datapoint{
			{2, 2},
		},
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "last",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{2, false},
			{3, true},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, false},
			{5, false},
			{6, true},
		},
	}.Run(t)

}
