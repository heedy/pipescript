package misc

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestChanged(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "changed",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: true},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: true},
		},
	}.Run(t)
}
