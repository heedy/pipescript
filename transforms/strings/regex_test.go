package strings

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestRegex(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "regex('^[a-z0-9_-]{3,16}$')",
		Input: []pipescript.Datapoint{
			{Timestamp:1,Data: "Hello World!"},
			{Timestamp:2,Data: "high1"},
			{Timestamp:3,Data: 34},
		},
		Output: []pipescript.Datapoint{
			{Timestamp:1,Data: false},
			{Timestamp:2,Data: true},
			{Timestamp:3,Data: false},
		},
	}.Run(t)
}
