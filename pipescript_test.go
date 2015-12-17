package pipescript

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {
	assert.Equal(t, 1, 1, "Oh no!")
}

func TestHello(t *testing.T) {
	PrintHello()
}
