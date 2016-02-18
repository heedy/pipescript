package interpolators

import (
	"testing"

	"github.com/connectordb/pipescript/interpolator"
)

func TestBefore(t *testing.T) {
	Register()
	interpolator.TestCase{
		Interpolator: "before",
		Input:        testDpa,

		Output: []interpolator.TestOutput{
			{0.5, nil},
			{2.5, &testDpa[1]},
			{5.0, &testDpa[4]},
			{6.0, &testDpa[6]},
			{8.0, &testDpa[8]},
			{20.0, &testDpa[8]},
			{30.0, &testDpa[8]},
		},
	}.Run(t)
}
