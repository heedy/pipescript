package datasets

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/transforms/numeric"
	"github.com/stretchr/testify/require"
)

func TestDataset(t *testing.T) {
	numeric.Count.Register()
	ds := NewTDataset(1, 5, 1)
	ipltr, err := GetInterpolator("count", nil, ds.Reference(), pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{Timestamp: 0.1, Data: 4},
		{Timestamp: 0.8, Data: 2},
		{Timestamp: 3.2, Data: 2},
		{Timestamp: 8, Data: 2},
	}))
	require.NoError(t, err)
	ds.Add("k1", ipltr)

	out := []pipescript.Datapoint{
		{Timestamp: 1, Data: map[string]interface{}{"k1": int64(2)}},
		{Timestamp: 2, Data: map[string]interface{}{"k1": int64(0)}},
		{Timestamp: 3, Data: map[string]interface{}{"k1": int64(0)}},
		{Timestamp: 4, Data: map[string]interface{}{"k1": int64(1)}},
	}

	for i := range out {
		dp, err := ds.Next(&pipescript.Datapoint{})
		require.NoError(t, err)
		require.EqualValues(t, &out[i], dp)
	}

	dp, err := ds.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Nil(t, dp)

}
