package interpolator_test

import (
	"errors"
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/interpolator"
	"github.com/heedy/pipescript/interpolator/interpolators"
	"github.com/stretchr/testify/assert"
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

func TestRegister(t *testing.T) {

	// Make sure empty interpolator is rejected
	assert.Error(t, interpolator.Interpolator{}.Register())

	// Regsiter an interpolator
	assert.NoError(t, interpolator.Interpolator{
		Name:        "test",
		Description: "I am testing!",

		Generator: func(name string, dpi pipescript.DatapointIterator) (interpolator.InterpolatorInstance, error) {
			return nil, errors.New("Not a real interpolator")
		},
	}.Register())

	v, ok := interpolator.InterpolatorRegistry["test"]
	assert.True(t, ok)
	assert.Equal(t, "I am testing!", v.Description)

	//Register an interpolator - but it already exists!
	assert.Error(t, interpolator.Interpolator{
		Name:        "test",
		Description: "fail",

		Generator: func(name string, dpi pipescript.DatapointIterator) (interpolator.InterpolatorInstance, error) {
			return nil, errors.New("Not a real interpolator")
		},
	}.Register())
}

func TestInterpolationIterator(t *testing.T) {
	dpi := pipescript.NewDatapointArrayIterator(testDpa)
	c, err := interpolators.NewClosestInterpolator(dpi)
	require.NoError(t, err)

	tr, err := interpolator.NewUniformRange(0, 11, 2)
	require.NoError(t, err)

	ii := interpolator.NewInterpolationIterator(c, tr)

	dp, err := ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[0])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[1])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[3])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[6])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[8])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &testDpa[8])
	dp, err = ii.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}
