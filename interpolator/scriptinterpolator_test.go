package interpolator

import (
	"testing"

	"github.com/connectordb/pipescript"
)

func TestScriptInterpolator(t *testing.T) {
	TestCase{
		Interpolator: "$", // Use a PipeScript
		Input: []pipescript.Datapoint{
			{1, "1"},
			{2, "2"},
			{4, "3"},
			{4, "4"},
			{5, "5"},
		},
		Output: []TestOutput{
			{0.1, nil},
			{1.3, &pipescript.Datapoint{1, "1"}},
			{2, &pipescript.Datapoint{2, "2"}},
			{3, nil},
			{4, &pipescript.Datapoint{4, "4"}},
			{4, nil},
			{6, &pipescript.Datapoint{5, "5"}},
			{8, nil},
		},
	}.Run(t)
}
