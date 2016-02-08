package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestReduce(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "reduce",
		ParseError: true,
	}.Run(t)

	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "reduce(sum)",
		Input: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{2, map[string]interface{}{"a": 5, "b": -6, "c": 7}},
			{3, map[string]interface{}{}},
		},
		Output: []pipescript.Datapoint{
			{1, float64(18)},
			{2, float64(6)},
			{3, nil},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{2, map[string]interface{}{"a": 5, "b": -6, "c": 7}},
			{3, map[string]interface{}{}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, float64(18)},
			{2, float64(6)},
			{3, nil},
		},
	}.Run(t)
}
