package interpolator

import (
	"testing"

	"github.com/connectordb/pipescript"
	"github.com/stretchr/testify/require"
)

func TestUniformRange(t *testing.T) {
	_, err := NewUniformRange(5, 2, 3)
	require.Error(t, err)

	r, err := NewUniformRange(3, 5, 1)
	require.NoError(t, err)
	f, err := r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 3.0, f)
	f, err = r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 4.0, f)
	f, err = r.Timestamp()
	require.EqualError(t, err, ErrEOF.Error())
}

func TestIteratorRange(t *testing.T) {
	dpi := pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{3, ""},
		{8, ""},
		{8, ""},
		{9, ""},
	})

	r, err := NewIteratorRange(dpi)
	require.NoError(t, err)
	f, err := r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 3.0, f)
	f, err = r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 8.0, f)
	f, err = r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 8.0, f)
	f, err = r.Timestamp()
	require.NoError(t, err)
	require.Equal(t, 9.0, f)
	f, err = r.Timestamp()
	require.EqualError(t, err, ErrEOF.Error())
}
