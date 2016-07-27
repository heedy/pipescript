package text

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestSentiment(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "sentiment",
		Input: []pipescript.Datapoint{
			{1, "Hello World!"},
			{2, "I hate this fucking place"},
			{3, "foo, I like to eat"},
		},
		Output: []pipescript.Datapoint{
			{1, float32(0)},
			{2, float32(-0.28)},
			{3, float32(0.08)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, "carbon fiber car"},
			{5, "I like traaains"},
			{6, "I am dismayed"},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, float32(0)},
			{5, float32(0.13333334)},
			{6, float32(-0.13333334)},
		},
	}.Run(t)
}
