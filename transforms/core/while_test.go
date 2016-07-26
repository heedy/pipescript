package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestReset(t *testing.T) {
	// This needs much more testing
	pipescript.TestCase{
		Pipescript: "while $!=1 sum",
		Input: []pipescript.Datapoint{
			{1, 5},
			{2, 2},
			{3, 3},
			{4, 1},
			{5, 2},
		},
		Output: []pipescript.Datapoint{
			{1, float64(5)},
			{2, float64(7)},
			{3, float64(10)},
			{4, float64(1)},
			{5, float64(3)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 5},
			{2, 2},
			{3, 3},
			{4, 1},
			{5, 2},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, float64(5)},
			{2, float64(7)},
			{3, float64(10)},
			{4, float64(1)},
			{5, float64(3)},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "while(count%3!=1, sum)",
		Input: []pipescript.Datapoint{
			{1, 5},
			{2, 2},
			{3, 3},
			{4, 1},
			{5, 2},
		},
		Output: []pipescript.Datapoint{
			{1, float64(5)},
			{2, float64(7)},
			{3, float64(10)},
			{4, float64(1)},
			{5, float64(3)},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, 5},
			{2, 2},
			{3, 3},
			{4, 1},
			{5, 2},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, float64(5)},
			{2, float64(7)},
			{3, float64(10)},
			{4, float64(1)},
			{5, float64(3)},
		},
	}.Run(t)

}
