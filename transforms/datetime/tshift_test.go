package datetime

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestTshift(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "tshift $",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "tshift 2.34",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 2},
			{Timestamp: 3, Data: 3},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 3.34, Data: 1},
			{Timestamp: 4.34, Data: 2},
			{Timestamp: 5.34, Data: 3},
		},
	}.Run(t)
}
