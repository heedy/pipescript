package bytestreams

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonDatapointReader(t *testing.T) {
	jsonarraytext := strings.NewReader(`[
  {
    "t": "1974-08-11T01:37:45+00:00",
    "steps": 14,
    "activity": "walking"

  },
  {
    "t": "1974-08-12T03:44:25+00:00",
    "steps": 10,
    "activity": "running"

  },
  {
    "t": "1974-08-12T04:17:45+00:00",
    "steps": 12,
    "activity": "walking"
  },
  {
    "t": "1974-08-12T05:24:25+00:00",
    "steps": 5,
    "activity": "running"
  }
]`)
	dpr, err := NewJSONDatapointReader(jsonarraytext, "", false)
	require.NoError(t, err)
	dp, err := dpr.Next()
	require.NoError(t, err)
	require.Equal(t, float64(145417065), dp.Timestamp)
	dp, err = dpr.Next()
	require.NoError(t, err)
	require.Equal(t, float64(145511065), dp.Timestamp)
	dp, err = dpr.Next()
	require.NoError(t, err)
	require.Equal(t, float64(145513065), dp.Timestamp)
	dp, err = dpr.Next()
	require.NoError(t, err)
	require.Equal(t, float64(145517065), dp.Timestamp)
	dp, err = dpr.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}
