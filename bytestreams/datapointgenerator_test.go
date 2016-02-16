package bytestreams

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	sample := map[string]interface{}{"foo": "2006-01-02T15:04:05-07:00", "bar": "hi!"}
	gen, err := NewDatapointGenerator(sample, "bla")
	require.Error(t, err)
	gen, err = NewDatapointGenerator(sample, "bar")
	require.Error(t, err)
	gen, err = NewDatapointGenerator(sample, "foo")
	require.NoError(t, err)
	dp, err := gen.Generate(sample)
	require.NoError(t, err)
	require.Equal(t, float64(1136239445), dp.Timestamp)

	gen, err = NewDatapointGenerator(sample, "")
	require.NoError(t, err)
	dp, err = gen.Generate(sample)
	require.NoError(t, err)
	require.Equal(t, float64(1136239445), dp.Timestamp)

}
