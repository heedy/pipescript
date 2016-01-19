package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestTshift(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "tshift $",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "tshift 2.34",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{3.34, 1},
			{4.34, 2},
			{5.34, 3},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{6.34, 4},
			{7.34, 5},
			{8.34, 6},
		},
	}.Run(t)
}
