package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestNext(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "next $",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "next 9999999",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "next (-1)",
		ParseError: true,
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "next",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 2},
			{3, 3},
		},
		Output: []pipescript.Datapoint{
			{2, 2},
			{3, 3},
			{3, nil},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, 4},
			{5, 5},
			{6, 6},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{5, 5},
			{6, 6},
			{6, nil},
		},
	}.Run(t)
}
