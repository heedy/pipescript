package bytestreams

import (
	"strings"
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

func TestDatapointReader(t *testing.T) {
	testdata := strings.NewReader(`[{"t": 1448610379.159, "d": 2}, {"t": 1448610383.248, "d": 3}, {"t": 1448610387.3, "d": 4}]`)

	dpr, err := NewDatapointReader(testdata)
	require.NoError(t, err)

	dp, err := dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.EqualValues(t, &pipescript.Datapoint{Timestamp: 1448610379.159, Data: float64(2)}, dp)

	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.EqualValues(t, &pipescript.Datapoint{Timestamp: 1448610383.248, Data: float64(3)}, dp)

	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.EqualValues(t, &pipescript.Datapoint{Timestamp: 1448610387.3, Data: float64(4)}, dp)

	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Nil(t, dp)

}
