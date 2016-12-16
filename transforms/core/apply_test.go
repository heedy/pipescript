package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestApply(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "apply($+2)",
		Input: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 0, "b": -45, "c": 80}},
			{2, map[string]interface{}{"a": 5, "b": 6, "c": -7}},
			{3, map[string]interface{}{}},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": float64(2), "b": float64(-43), "c": float64(82)}},
			{2, map[string]interface{}{"a": float64(7), "b": float64(8), "c": float64(-5)}},
			{3, map[string]interface{}{}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{2, map[string]interface{}{"a": 5, "b": -6, "c": 7}},
			{3, map[string]interface{}{"a": -1.0}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": float64(7), "b": float64(8), "c": float64(9)}},
			{2, map[string]interface{}{"a": float64(7), "b": float64(-4), "c": float64(9)}},
			{3, map[string]interface{}{"a": float64(1.0)}},
		},
	}.Run(t)
}
