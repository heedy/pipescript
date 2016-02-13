package pipescript_test

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestObject(t *testing.T) {
	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: `{"hi": 1, "ho": 4}`,
		Input: []pipescript.Datapoint{
			{1, 4},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{"hi": float64(1), "ho": float64(4)}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 4},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{"hi": float64(1), "ho": float64(4)}},
		},
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: `{"hi": 1, "hi": 4}`,
		ParseError: true,
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: `{"sum": sum:$, "count": count} | if last`,
		Input: []pipescript.Datapoint{
			{1, 4},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{3, map[string]interface{}{"sum": float64(9), "count": int64(3)}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 4},
			{2, 2},
			{3, 5},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{3, map[string]interface{}{"sum": float64(11), "count": int64(3)}},
		},
	}.Run(t)
}
