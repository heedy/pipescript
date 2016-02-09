package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestRemember(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "remember($==true)",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{3, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, 1},
			{2, true},
			{2, true},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, 4},
			{4, 4},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "remember($==true,sum)",
		Input: []pipescript.Datapoint{
			{1, 2},
			{2, true},
			{3, 2},
		},
		Output: []pipescript.Datapoint{
			{1, float64(2)},
			{2, float64(3)},
			{2, float64(3)},
		},
	}.Run(t)

}
