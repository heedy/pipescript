package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestMean(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "mean",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 3},
			{3, 5},
		},
		Output: []pipescript.Datapoint{
			{1, float64(1)},
			{2, float64(2)},
			{3, float64(3)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, float64(4)},
			{5, float64(4.5)},
			{6, float64(5)},
		},
	}.Run(t)
}
