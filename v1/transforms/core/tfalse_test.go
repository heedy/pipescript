package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestTfalse(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "tfalse",
		Input: []pipescript.Datapoint{
			{1, false},
			{2, false},
			{4, false},
			{5, true},
			{6, true},
			{7, false},
			{8, true},
			{9, false},
		},
		Output: []pipescript.Datapoint{
			{5, float64(4)},
			{8, float64(1)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{23.4, false},
			{23.4, true},
			{23.5, true},
			{25, false},
			{27, true},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{23.4, float64(0)},
			{27, float64(2)},
		},
	}.Run(t)

	// On a single point, it doesn't return anything, since there is no time to check
	pipescript.TestCase{
		Pipescript: "tfalse",
		Input: []pipescript.Datapoint{
			{1, false},
		},
		Output: []pipescript.Datapoint{},
		SecondaryInput: []pipescript.Datapoint{
			{1, true},
		},
		SecondaryOutput: []pipescript.Datapoint{},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "tfalse",
		Input: []pipescript.Datapoint{
			{5, true},
			{6, true},
			{7, false},
			{8, false},
			{9, false},
		},
		Output: []pipescript.Datapoint{
			{9, float64(2)},
		},
	}.Run(t)
}
