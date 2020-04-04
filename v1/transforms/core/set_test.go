package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestSet(t *testing.T) {
	pipescript.TestCase{
		Pipescript: `set("hi","ho2")`,
		Input: []pipescript.Datapoint{
			{1, map[string]string{}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
			{3, map[string]string{"hi": "world"}},
		},
		Output: []pipescript.Datapoint{
			{1, map[string]interface{}{"hi": "ho2"}},
			{2, map[string]interface{}{"hello": "world", "1": "bar", "hi": "ho2"}},
			{3, map[string]interface{}{"hi": "ho2"}},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, map[string]string{}},
			{2, map[string]string{"hello": "world", "1": "bar"}},
			{3, map[string]string{"hi": "world"}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, map[string]interface{}{"hi": "ho2"}},
			{2, map[string]interface{}{"hello": "world", "1": "bar", "hi": "ho2"}},
			{3, map[string]interface{}{"hi": "ho2"}},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "set($,$)", // The transform requries a constant
		ParseError: true,
	}.Run(t)

}
