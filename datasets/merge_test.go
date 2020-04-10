package datasets

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

func TestMerge(t *testing.T) {
	dpa1 := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{Timestamp: 1, Data: 1},
		{Timestamp: 2, Data: 1},
		{Timestamp: 3, Data: 1},
	})

	dpa2 := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{Timestamp: 1.1, Data: 2},
		{Timestamp: 2, Data: 2},
		{Timestamp: 2.9, Data: 2},
		{Timestamp: 3.5, Data: 2},
	})
	m, err := Merge([]pipescript.Iterator{dpa1, dpa2})
	require.NoError(t, err)

	dpa3 := []pipescript.Datapoint{
		{Timestamp: 1, Data: 1},
		{Timestamp: 1.1, Data: 2},
		{Timestamp: 2, Data: 1},
		{Timestamp: 2, Data: 2},
		{Timestamp: 2.9, Data: 2},
		{Timestamp: 3, Data: 1},
		{Timestamp: 3.5, Data: 2},
	}

	for i := range dpa3 {
		v, err := m.Next(&pipescript.Datapoint{})
		require.NoError(t, err)
		require.EqualValues(t, &dpa3[i], v)
	}

	v, err := m.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Nil(t, v)

}
