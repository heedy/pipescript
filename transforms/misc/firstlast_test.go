package misc

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestLast(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "last",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 2},
			{Timestamp: 3, Data: 3},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: false},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: true},
		},
	}.Run(t)

}

func TestFirst(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "first",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 1},
			{Timestamp: 2, Data: 2},
			{Timestamp: 3, Data: 3},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: true},
			{Timestamp: 2, Data: false},
			{Timestamp: 3, Data: false},
		},
	}.Run(t)
}
