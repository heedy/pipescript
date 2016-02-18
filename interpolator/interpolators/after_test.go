package interpolators

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
)

// The standard datapoint array to use when testing
var testDpa = []pipescript.Datapoint{
	{1., "test0"},
	{2., "test1"},
	{3., "test2"},
	{4., "test3"},
	{5., "test4"},
	{6., "test5"},
	{6., "test6"},
	{7., "test7"},
	{8., "test8"},
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
