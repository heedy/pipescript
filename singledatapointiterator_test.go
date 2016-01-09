package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleDatapointIterator(t *testing.T) {
	dp := &Datapoint{1, 1}
	dpi := SingleDatapointIterator{}
	dpi.Set(dp, nil)
	dp2, err := dpi.Next()
	require.NoError(t, err)
	require.Equal(t, dp, dp2)
}
