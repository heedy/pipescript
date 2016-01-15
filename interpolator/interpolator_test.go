package interpolator

import (
	"errors"
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/assert"
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
	assert.Error(t, Interpolator{}.Register())

	// Regsiter an interpolator
	assert.NoError(t, Interpolator{
		Name:        "test",
		Description: "I am testing!",

		Generator: func(name string, dpi pipescript.DatapointIterator) (InterpolatorInstance, error) {
			return nil, errors.New("Not a real interpolator")
		},
	}.Register())

	v, ok := InterpolatorRegistry["test"]
	assert.True(t, ok)
	assert.Equal(t, "I am testing!", v.Description)

	//Register an interpolator - but it already exists!
	assert.Error(t, Interpolator{
		Name:        "test",
		Description: "fail",

		Generator: func(name string, dpi pipescript.DatapointIterator) (InterpolatorInstance, error) {
			return nil, errors.New("Not a real interpolator")
		},
	}.Register())
}
