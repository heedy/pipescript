package strings

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestRegex(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "regex('^[a-z0-9_-]{3,16}$')",
		Input: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, "high1"},
			{3, 34},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{2, true},
			{3, false},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, "high1"},
			{3, 34},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, false},
			{2, true},
			{3, false},
		},
	}.Run(t)
}
