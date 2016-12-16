package math

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestTop(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "percent",
		Input: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 0, "b": -45, "c": 80}},
			{2, map[string]interface{}{"a": 5, "b": 6, "c": -7}},
			{3, map[string]interface{}{}},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": float64(0), "b": float64(0), "c": float64(1)}},
			{2, map[string]interface{}{"a": float64(5) / float64(11), "b": float64(6) / float64(11), "c": float64(0)}},
			{3, map[string]interface{}{}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{2, map[string]interface{}{"a": 5, "b": -6, "c": 7}},
			{3, map[string]interface{}{"a": -1.0}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": float64(5) / float64(18), "b": float64(6) / float64(18), "c": float64(7) / float64(18)}},
			{2, map[string]interface{}{"a": float64(5) / float64(12), "b": float64(0), "c": float64(7) / float64(12)}},
			{3, map[string]interface{}{"a": float64(0)}},
		},
	}.Run(t)
}
