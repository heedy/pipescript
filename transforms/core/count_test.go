package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestCount(t *testing.T) {
	pipescript.TestCase{
		Pipescript: "count",
		Input: []pipescript.Datapoint{
			{1, int64(2)},
			{2, 3},
			{3, 4},
		},
		Output: []pipescript.Datapoint{
			{1, int64(1)},
			{2, int64(2)},
			{3, int64(3)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, int64(1)},
			{5, int64(2)},
			{6, int64(3)},
		},
	}.Run(t)
}
