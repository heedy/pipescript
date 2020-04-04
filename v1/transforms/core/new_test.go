package core

import (
	"testing"

	"github.com/heedy/pipescript"
)

func TestNew(t *testing.T) {
	Register()
	pipescript.TestCase{
		// This tests order of prescedence: ":" pipes are high prescedence, and will be executed first
		Pipescript: "new",
		Input: []pipescript.Datapoint{
			{1, "hi"},
			{2, "hi"},
			{3, "hoi"},
			{4, "1"},
			{5, 1},
			{6, 1.0},
		},
		Output: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, true},
			{4, true},
			{5, false},
			{6, false},
		},
		SecondaryInput: []pipescript.Datapoint{
			{1, "hi"},
			{2, "hi"},
			{3, "hoi"},
			{4, "1"},
			{5, 1},
			{6, 1.0},
		},
		SecondaryOutput: []pipescript.Datapoint{
			{1, true},
			{2, false},
			{3, true},
			{4, true},
			{5, false},
			{6, false},
		},
	}.Run(t)
}
