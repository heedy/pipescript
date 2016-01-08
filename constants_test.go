package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstant(t *testing.T) {
	cs := ConstantScript(1337)

	require.True(t, cs.Constant)
	require.True(t, cs.IsOneToOne)

	testarray := []Datapoint{
		Datapoint{1, 1}, Datapoint{2, 2}, Datapoint{3, 3}, Datapoint{4, 4},
	}
	cs.SetInput(NewDatapointArrayIterator(testarray))

	for i := 1; i < 5; i++ {
		dp, err := cs.Next()
		require.NoError(t, err)
		require.EqualValues(t, &Datapoint{float64(i), 1337}, dp)
	}

	dp, err := cs.Next()
	require.NoError(t, err)
	require.Nil(t, dp)
}
