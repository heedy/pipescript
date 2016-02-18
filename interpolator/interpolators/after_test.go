package interpolators

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// The standard datapoint array to use when testing
var testDpa = []pipescript.Datapoint{
	{1., 10},
	{2., 20},
	{3., 30},
	{4., 40},
	{5., 50},
	{6., 60},
	{6., 70},
	{7., 80},
	{8., 90},
}

func TestAfter(t *testing.T) {
	Register()
	interpolator.TestCase{
		Interpolator: "after",
		Input:        testDpa,
		Output: []interpolator.TestOutput{
			{0.5, &testDpa[0]},
			{0.7, &testDpa[0]},
			{2.0, &testDpa[2]},
			{5.5, &testDpa[5]},
			{6.0, &testDpa[7]},
			{8.0, nil},
		},
	}.Run(t)
}
