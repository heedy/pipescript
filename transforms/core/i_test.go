package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestI(t *testing.T) {
	I.Register()
	pipescript.TestCase{
		Pipescript: "i",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: int64(2)},
			{Timestamp: 2, Data: 3},
			{Timestamp: 3, Data: 4},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: int64(0)},
			{Timestamp: 2, Data: int64(1)},
			{Timestamp: 3, Data: int64(2)},
		},
	}.Run(t)
}
