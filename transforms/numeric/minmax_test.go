package numeric

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestMin(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "min",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 2, Data: 1},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "min(-$)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
		},
	}.Run(t)

}

func TestMax(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "max(-$)",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 2, Data: 1},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "max",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
			{Timestamp: 2, Data: 1},
			{Timestamp: 3, Data: 2},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Data: 3},
		},
	}.Run(t)
}
