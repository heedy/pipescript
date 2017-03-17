package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestDel(t *testing.T) {
	pipescript.TestCase{
		Pipescript: `del("hi")`,
		Input: []pipescript.Datapoint{
			{1, map[string]string{}},
			{2, map[string]string{"hi": "world", "1": "bar"}},
			{3, map[string]string{"hi": "world"}},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{}},
			{2, map[string]interface{}{"1": "bar"}},
			{3, map[string]interface{}{}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]string{}},
			{2, map[string]string{"hi": "world", "1": "bar"}},
			{3, map[string]string{"hi": "world"}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{}},
			{2, map[string]interface{}{"1": "bar"}},
			{3, map[string]interface{}{}},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "del($)", // The transform requries a constant
		ParseError: true,
	}.Run(t)

}
