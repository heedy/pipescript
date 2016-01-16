package interpolator

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

func TestDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa)
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i1, err := NewClosestInterpolator(dpi1)
	require.NoError(t, err)
	i2, err := NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := Dataset(map[string]InterpolatorInstance{"foo": i1, "bar": i2})

	dp, err := ds.Interpolate(1.1)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1.1, Data: map[string]*pipescript.Datapoint{
		"foo": {1, "test0"},
		"bar": {2, "test1"},
	}})
	dp, err = ds.Interpolate(5)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 5, Data: map[string]*pipescript.Datapoint{
		"foo": {5, "test4"},
		"bar": {6, "test5"},
	}})

	dp, err = ds.Interpolate(10)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 10, Data: map[string]*pipescript.Datapoint{
		"foo": {8, "test8"},
		"bar": nil,
	}})
}

func TestXDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa[0:3])
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i2, err := NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := Dataset(map[string]InterpolatorInstance{"bar": i2})

	_, err = GetXDataset(dpi1, "bar", ds)
	require.Error(t, err)

	xd, err := GetXDataset(dpi1, "foo", ds)
	require.NoError(t, err)

	dp, err := xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1, Data: map[string]*pipescript.Datapoint{
		"foo": {1, "test0"},
		"bar": {2, "test1"},
	}})
	dp, err = xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 2, Data: map[string]*pipescript.Datapoint{
		"foo": {2, "test1"},
		"bar": {3, "test2"},
	}})
	dp, err = xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 3, Data: map[string]*pipescript.Datapoint{
		"foo": {3, "test2"},
		"bar": {4, "test3"},
	}})

	dp, err = xd.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}

func TestTDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa)
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i1, err := NewClosestInterpolator(dpi1)
	require.NoError(t, err)
	i2, err := NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := Dataset(map[string]InterpolatorInstance{"foo": i1, "bar": i2})

	td, err := GetTDataset(1, 2, 1, ds)
	require.NoError(t, err)

	dp, err := td.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1, Data: map[string]*pipescript.Datapoint{
		"foo": {1, "test0"},
		"bar": {2, "test1"},
	}})
	dp, err = td.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}
