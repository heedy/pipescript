package strings

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestWc(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "wc",
		Input: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, ""},
			{3, "foo, I like to eat"},
		},
		Output: []pipescript.Datapoint{
			{1, 2},
			{2, 0},
			{3, 5},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, "carbon fiber car"},
			{5, "I like traaains"},
			{6, "english mother*****, do you speak it?"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, 3},
			{5, 3},
			{6, 6},
		},
	}.Run(t)
}
