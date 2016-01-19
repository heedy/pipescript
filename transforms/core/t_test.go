package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestT(t *testing.T) {
	Register()
	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "t",
		Input: []pipescript.Datapoint{
			{1, "hi"},
			{2, "hi"},
			{3, "hi"},
		},
		Output: []pipescript.Datapoint{
			{1, float64(1)},
			{2, float64(2)},
			{3, float64(3)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, "hi"},
			{5, "hi"},
			{6, "hi"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, float64(4)},
			{5, float64(5)},
			{6, float64(6)},
		},
	}.Run(t)
}
