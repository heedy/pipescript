package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestDT(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "dt",
		Input: []pipescript.Datapoint{
			{1, int64(2)},
			{2, 3},
			{8, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(0)},
			{2, float64(1)},
			{8, float64(6)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{20, 5},
			{50, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, float64(0)},
			{20, float64(16)},
			{50, float64(30)},
		},
	}.Run(t)
}
