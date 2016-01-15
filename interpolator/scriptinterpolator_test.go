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
		InterpolationTime: []float64{0.1, 1.3, 2, 3, 4, 4, 6, 8},
		Output: []*pipescript.Datapoint{
			nil,
			{1, "1"},
			{2, "2"},
			nil,
			{4, "4"},
			nil,
			{5, "5"},
			nil,
		},
	}.Run(t)
}
