package config

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
	"github.com/connectordb/pipescript/interpolator/interpolators"
	"github.com/connectordb/pipescript/transforms/core"
	"github.com/stretchr/testify/require"
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

func TestConfig(t *testing.T) {
	d := Default()
	d.NextMax = 45

	d.Save("test_default.conf")
	require.NoError(t, d.Validate())

	d2, err := Load("test_default.conf")
	require.NoError(t, err)
	require.Equal(t, d.NextMax, d2.NextMax)

	require.NoError(t, d2.Set())

}

func TestDisable(t *testing.T) {
	core.Register()
	interpolators.Register()

	pipescript.TestCase{
		Pipescript: "sum",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 3},
			{3, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(1)},
			{2, float64(4)},
			{3, float64(8)},
		},
	}.Run(t)
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

	d := Default()

	d.DisableTransforms = []string{"sum"}
	d.DisableInterpolators = []string{"before"}

	d.Set()

	pipescript.TestCase{
		Pipescript: "sum",
		ParseError: true,
	}.Run(t)
	interpolator.TestCase{
		Interpolator: "before",
		ParseError:   true,
	}.Run(t)

	// Now disabling different ones reenables the old ones
	d.DisableTransforms = []string{"count"}
	d.DisableInterpolators = []string{"after"}
	d.Set()

	pipescript.TestCase{
		Pipescript: "sum",
		Input: []pipescript.Datapoint{
			{1, 1},
			{2, 3},
			{3, 4},
		},
		Output: []pipescript.Datapoint{
			{1, float64(1)},
			{2, float64(4)},
			{3, float64(8)},
		},
	}.Run(t)
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
