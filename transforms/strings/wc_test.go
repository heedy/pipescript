package strings

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestWc(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "wc",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "Hello World!"},
			{Timestamp: 2, Data: ""},
			{Timestamp: 3, Data: "foo, I like to eat"},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: 2},
			{Timestamp: 2, Data: 0},
			{Timestamp: 3, Data: 5},
		},
	}.Run(t)
}
