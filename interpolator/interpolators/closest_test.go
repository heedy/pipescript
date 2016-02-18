package interpolators

import (
	"testing"

	"github.com/connectordb/pipescript/interpolator"
)

func TestClosest(t *testing.T) {
	Register()
	interpolator.TestCase{
		Interpolator: "closest",
		Input:        testDpa,

		Output: []interpolator.TestOutput{
			{0.5, &testDpa[0]},
			{2.1, &testDpa[1]}, //Make sure it gives closest value less
			{2.5, &testDpa[1]}, //Make sure that it gives smaller if equal dist
			{2.6, &testDpa[2]}, //Make sure that it gives greater if larger distance
			{5.0, &testDpa[4]}, //Make sure it can iterate through many
			{5.9, &testDpa[5]}, //Make sure it shows first of 2 when less
			{6.0, &testDpa[6]}, //Make sure it shows second of 2 when equal
			{8.0, &testDpa[8]}, //Make sure it ends by keeping oldest datapoint
			{20.0, &testDpa[8]},
			{30.0, &testDpa[8]},
		},
	}.Run(t)
}
