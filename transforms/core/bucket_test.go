package core

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestBucket(t *testing.T) {
	Register()
	pipescript.TestCase{
		Pipescript: "bucket",
		Input: []pipescript.Datapoint{
			{1, 2},
			{2, 16},
			{3, 84},
			{4, -5},
			{5, 1},
		},
		Output: []pipescript.Datapoint{
			{1, "[0,10)"},
			{2, "[10,20)"},
			{3, "[80,90)"},
			{4, "[-10,0)"},
			{5, "[0,10)"},
		},
		SecondaryInput: []pipescript.Datapoint{
			{4, "0"},
			{5, -0.1},
			{6, 3.14},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{4, "[0,10)"},
			{5, "[-10,0)"},
			{6, "[0,10)"},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "bucket(0.5,0.1)",
		Input: []pipescript.Datapoint{
			{1, 2},
			{2, 2.1},
			{3, -0.4},
		},
		Output: []pipescript.Datapoint{
			{1, "[1.6,2.1)"},
			{2, "[2.1,2.6)"},
			{3, "[-0.4,0.1)"},
		},
	}.Run(t)
}
