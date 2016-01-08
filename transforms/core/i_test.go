package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestI(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "i",
		Input: []pipescript.Datapoint{
			{1, int64(1)},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{1, int64(0)},
			{2, int64(1)},
			{3, int64(2)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, int64(0)},
			{5, int64(1)},
			{6, int64(2)},
		},
	}.Run(t)
}
