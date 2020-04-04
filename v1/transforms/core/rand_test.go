package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/heedy/pipescript"
)

func TestRand(t *testing.T) {
	s, err := pipescript.Parse("rand")
	require.NoError(t, err)

	iter := &pipescript.SingleDatapointIterator{}

	s.SetInput(iter)
	iter.Set(&pipescript.Datapoint{Timestamp: 1, Data: 1}, nil)

	v1, err := s.Next()
	require.NoError(t, err)
	v2, err := s.Next()
	require.NoError(t, err)

	v1f, err := v1.Float()
	require.NoError(t, err)
	v2f, err := v2.Float()
	require.NoError(t, err)

	require.True(t, v1f != v2f)
}
