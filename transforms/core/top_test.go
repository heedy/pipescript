package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestTop(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "top(2)",
		Input: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 0, "b": -45, "c": 80}},
			{2, map[string]interface{}{"a": 5, "b": 6, "c": -7}},
			{3, map[string]interface{}{}},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 0, "c": 80}},
			{2, map[string]interface{}{"a": 5, "b": 6}},
			{3, map[string]interface{}{}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{2, map[string]interface{}{"a": 5, "b": -6, "c": 7}},
			{3, map[string]interface{}{"a": -1.0}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{ "b": 6,"c": 7}},
			{2, map[string]interface{}{"a": 5, "c": 7}},
			{3, map[string]interface{}{"a": -1.0}},
		},
	}.Run(t)
}
