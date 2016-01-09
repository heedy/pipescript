package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestSum(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "sum",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 3},
			{3, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(1)},
			{2, float64(4)},
			{3, float64(8)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, float64(4)},
			{5, float64(9)},
			{6, float64(15)},
		},
	}.Run(t)
}
