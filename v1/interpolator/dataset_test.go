package interpolator_test

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/interpolator"
	"github.com/heedy/pipescript/interpolator/interpolators"
	"github.com/stretchr/testify/require"
)

func TestDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa)
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i1, err := interpolators.NewClosestInterpolator(dpi1)
	require.NoError(t, err)
	i2, err := interpolators.NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := interpolator.Dataset(map[string]interpolator.InterpolatorInstance{"foo": i1, "bar": i2})

	dp, err := ds.Interpolate(1.1)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1.1, Data: map[string]interface{}{
		"foo": "test0",
		"bar": "test1",
	}})
	dp, err = ds.Interpolate(5)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 5, Data: map[string]interface{}{
		"foo": "test4",
		"bar": "test5",
	}})

	dp, err = ds.Interpolate(10)
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 10, Data: map[string]interface{}{
		"foo": "test8",
		"bar": nil,
	}})
}

func TestXDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa[0:3])
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i2, err := interpolators.NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := interpolator.Dataset(map[string]interpolator.InterpolatorInstance{"bar": i2})

	_, err = interpolator.GetXDataset(dpi1, "bar", ds)
	require.Error(t, err)

	xd, err := interpolator.GetXDataset(dpi1, "foo", ds)
	require.NoError(t, err)

	dp, err := xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1, Data: map[string]interface{}{
		"foo": "test0",
		"bar": "test1",
	}})
	dp, err = xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 2, Data: map[string]interface{}{
		"foo": "test1",
		"bar": "test2",
	}})
	dp, err = xd.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 3, Data: map[string]interface{}{
		"foo": "test2",
		"bar": "test3",
	}})

	dp, err = xd.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}

func TestTDataset(t *testing.T) {
	dpi1 := pipescript.NewDatapointArrayIterator(testDpa)
	dpi2 := pipescript.NewDatapointArrayIterator(testDpa)
	i1, err := interpolators.NewClosestInterpolator(dpi1)
	require.NoError(t, err)
	i2, err := interpolators.NewAfterInterpolator(dpi2)
	require.NoError(t, err)

	ds := interpolator.Dataset(map[string]interpolator.InterpolatorInstance{"foo": i1, "bar": i2})

	td, err := interpolator.GetTDataset(1, 2, 1, ds)
	require.NoError(t, err)

	dp, err := td.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, &pipescript.Datapoint{Timestamp: 1, Data: map[string]interface{}{
		"foo": "test0",
		"bar": "test1",
	}})
	dp, err = td.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}
