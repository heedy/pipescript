package interpolator

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	dpa1 := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{1, 1},
		{2, 1},
		{3, 1},
	})

	dpa2 := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{1.1, 2},
		{2, 2},
		{2.9, 2},
		{3.5, 2},
	})
	m, err := Merge([]pipescript.DatapointIterator{dpa1, dpa2})
	require.NoError(t, err)

	dpa3 := []pipescript.Datapoint{
		{1, 1},
		{1.1, 2},
		{2, 1},
		{2, 2},
		{2.9, 2},
		{3, 1},
		{3.5, 2},
	}

	for i := range dpa3 {
		v, err := m.Next()
		require.NoError(t, err)
		require.EqualValues(t, &dpa3[i], v)
	}

	v, err := m.Next()
	require.NoError(t, err)
	require.Nil(t, v)

}
