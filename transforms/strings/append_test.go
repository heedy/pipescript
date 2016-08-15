package strings

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestCount(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "append",
		Input: []pipescript.Datapoint{
			{1, 2},
			{2, " "},
			{3, "Hello!"},
		},
		Output: []pipescript.Datapoint{
			{1, "2"},
			{2, "2 "},
			{3, "2 Hello!"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 2},
			{2, " "},
			{3, "Hello!"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, "2"},
			{2, "2 "},
			{3, "2 Hello!"},
		},
	}.Run(t)
}
