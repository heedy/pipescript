package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestAllTrue(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "alltrue",
		Input: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, true},
		},
		Output: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, false},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, true},
			{5, true},
			{6, true},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, true},
			{5, true},
			{6, true},
		},
	}.Run(t)
}
