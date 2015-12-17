/**
Copyright (c) 2015 The PipeScript Contributors (see AUTHORS)
Licensed under the MIT license.
**/
package pipescript

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	//Make sure BS transforms are rejected
	assert.Error(t, Transform{}.Register())

	//Register a transform
	assert.NoError(t, Transform{
		Name:        "test",
		Description: "I am testing!",
	}.Register())

	v, ok := TransformRegistry["test"]
	assert.True(t, ok)
	assert.Equal(t, "I am testing!", v.Description)
}