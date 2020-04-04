package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestReduce(t *testing.T) {
	Reduce.Register()
	pipescript.TestCase{
		Pipescript: "reduce",
		Parsed:     "error",
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "reduce(i+1)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: map[string]interface{}{"a": 5, "b": 6, "c": 7}},
			{Timestamp: 2, Data: map[string]interface{}{"a": 5, "b": -6}},
			{Timestamp: 3, Data: map[string]interface{}{}},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: float64(3)},
			{Timestamp: 2, Data: float64(2)},
			{Timestamp: 3, Data: nil},
		},
	}.Run(t)
}
