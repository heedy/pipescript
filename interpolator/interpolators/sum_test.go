package interpolators

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

func TestSum(t *testing.T) {
	Register()
	interpolator.TestCase{
		Interpolator: "sum",
		Input:        testDpa,

		Output: []interpolator.TestOutput{
			{0.5, &pipescript.Datapoint{0.5, 0.0}},
			{2.5, &pipescript.Datapoint{2.5, float64(30)}},
			{5.0, &pipescript.Datapoint{5, float64(120)}},
			{6.0, &pipescript.Datapoint{6, float64(130)}},
			{8.0, &pipescript.Datapoint{8, float64(170)}},
			{20.0, &pipescript.Datapoint{20, 0.0}},
			{30.0, &pipescript.Datapoint{30, 0.0}},
		},
	}.Run(t)
}
