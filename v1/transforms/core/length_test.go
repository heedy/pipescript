package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestLength(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "length",
		Input: []pipescript.Datapoint{
			{1, "hi"},
			{2, []int{1, 2, 3, 4, 5}},
			{3, map[string]interface{}{"hi": nil, "hr": nil, "ree": nil}},
		},
		Output: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{3, 3},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "hi"},
			{2, []int{1, 2, 3, 4, 5}},
			{3, map[string]interface{}{"hi": nil, "hr": nil, "ree": nil}},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, 2},
			{2, 5},
			{3, 3},
		},
	}.Run(t)
}
