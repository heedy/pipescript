package interpolator

import (
	"errors"
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/assert"
)

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
