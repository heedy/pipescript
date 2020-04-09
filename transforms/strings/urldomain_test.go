package strings

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestUrldomain(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "urldomain",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: "http://google.com/hi"},
			{Timestamp: 2, Data: "random"},
			{Timestamp: 3, Data: 34},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: "google.com"},
			{Timestamp: 2, Data: ""},
			{Timestamp: 3, Data: ""},
		},
	}.Run(t)
}
