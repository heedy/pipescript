package interpolators

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/interpolator"
)

func TestCount(t *testing.T) {
	Register()
	interpolator.TestCase{
		Interpolator: "count",
		Input:        testDpa,

		Output: []interpolator.TestOutput{
			{0.5, &pipescript.Datapoint{0.5, 0}},
			{2.5, &pipescript.Datapoint{2.5, 2}},
			{5.0, &pipescript.Datapoint{5, 3}},
			{6.0, &pipescript.Datapoint{6, 2}},
			{8.0, &pipescript.Datapoint{8, 2}},
			{20.0, &pipescript.Datapoint{20, 0}},
			{30.0, &pipescript.Datapoint{30, 0}},
		},
	}.Run(t)
}
