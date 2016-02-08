package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestIfelse(t *testing.T) {
	// This needs much more testing
	pipescript.TestCase{
		Pipescript: "ifelse($>1,0)",
		Input: []pipescript.Datapoint{
			{1, 5},
			{2, 1},
		},
		Output: []pipescript.Datapoint{
			{1, float64(0)},
			{2, 1},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 5},
			{2, 1},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, float64(0)},
			{2, 1},
		},
	}.Run(t)

}
