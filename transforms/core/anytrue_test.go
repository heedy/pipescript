package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestAnyTrue(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "anytrue",
		Input: []pipescript.Datapoint{
			{1, false},
			{2, false},
			{3, true},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{2, false},
			{3, true},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, true},
			{5, false},
			{6, false},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, true},
			{5, true},
			{6, true},
		},
	}.Run(t)
}
