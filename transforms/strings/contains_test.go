package strings

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestContains(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "contains('hi')",
		Input: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, " high"},
			{3, 34},
		},
		Output: []pipescript.Datapoint{
			{1, false},
			{2, true},
			{3, false},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, " high"},
			{3, 34},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, false},
			{2, true},
			{3, false},
		},
	}.Run(t)
}
