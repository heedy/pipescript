package bytestreams

import (
	"strings"
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

func TestCsvDatapointReader(t *testing.T) {
	csvtext := strings.NewReader(`time,steps,activity
1974-08-11T01:37:45+00:00,14,walking
1974-08-12T03:44:25+00:00,10,running
1974-08-12T04:17:45+00:00, 12,walking
1974-08-12T05:24:25+00:00,5,running
`)
	dpr, err := NewCSVDatapointReader(csvtext, "", false)
	require.NoError(t, err)
	dp, err := dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Equal(t, float64(145417065), dp.Timestamp)
	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Equal(t, float64(145511065), dp.Timestamp)
	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Equal(t, float64(145513065), dp.Timestamp)
	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Equal(t, float64(145517065), dp.Timestamp)
	dp, err = dpr.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Nil(t, dp)
}
