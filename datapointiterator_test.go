package pipescript

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPeek(t *testing.T) {
	testarray := []Datapoint{
		Datapoint{1, 1}, Datapoint{2, 2}, Datapoint{3, 3}, Datapoint{4, 4},
	}
	pi := NewDatapointPeekIterator(NewDatapointArrayIterator(testarray))

	dp, err := pi.Next()
	assert.NoError(t, err)
	require.EqualValues(t, testarray[0], *dp)

	dp, err = pi.Peek(0)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[1], *dp)

	dp, err = pi.Next()
	assert.NoError(t, err)
	require.EqualValues(t, testarray[1], *dp)

	dp, err = pi.Peek(1)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[3], *dp)

	dp, err = pi.Peek(1)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[3], *dp)

	dp, err = pi.Peek(0)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[2], *dp)

	dp, err = pi.Peek(20)
	assert.NoError(t, err)
	require.Nil(t, dp)

	dp, err = pi.Peek(20)
	assert.NoError(t, err)
	require.Nil(t, dp)

	dp, err = pi.Next()
	assert.NoError(t, err)
	require.EqualValues(t, testarray[2], *dp)

	dp, err = pi.Next()
	assert.NoError(t, err)
	require.EqualValues(t, testarray[3], *dp)

	dp, err = pi.Next()
	assert.NoError(t, err)
	require.Nil(t, dp)

	dp, err = pi.Peek(0)
	assert.NoError(t, err)
	require.Nil(t, dp)
}

func TestPeek2(t *testing.T) {
	testarray := []Datapoint{
		Datapoint{1, 1}, Datapoint{2, 2}, Datapoint{3, 3}, Datapoint{4, 4}, Datapoint{5, 5},
	}
	pi := NewDatapointPeekIterator(NewDatapointArrayIterator(testarray))

	dp, err := pi.Peek(20)
	assert.NoError(t, err)
	require.Nil(t, dp)

	dp, err = pi.Peek(1)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[1], *dp)

	dp, err = pi.Peek(3)
	assert.NoError(t, err)
	require.EqualValues(t, testarray[3], *dp)

	pi.Err = errors.New("HadError")
	_, err = pi.Peek(20)
	assert.Error(t, err)

}

func TestVirtualPeek(t *testing.T) {
	testarray := []Datapoint{
		Datapoint{1, 1}, Datapoint{2, 2}, Datapoint{3, 3}, Datapoint{4, 4}, Datapoint{5, 5},
	}

	pi := NewDatapointPeekIterator(NewDatapointArrayIterator(testarray))

	pi2 := NewVirtualPeekIterator(pi)

	dp, err := pi.Peek(0)
	require.NoError(t, err)
	dp2, err := pi2.Peek(0)
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	dp, err = pi.Peek(80)
	require.NoError(t, err)
	dp2, err = pi2.Peek(80)
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	dp, err = pi.Peek(0)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	dp, err = pi.Peek(1)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	pi2.Reset()

	dp, err = pi.Peek(0)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	dp, err = pi.Peek(1)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	dp, err = pi.Peek(2)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	pi2.SetBack(2)

	dp, err = pi.Peek(1)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	pi2.SetBack(50)

	dp, err = pi.Peek(0)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)

	// Make sure that Next on the original will have the expected effect
	_, err = pi.Next()
	require.NoError(t, err)
	dp, err = pi.Peek(1)
	require.NoError(t, err)
	dp2, err = pi2.Next()
	require.NoError(t, err)
	require.EqualValues(t, dp, dp2)
}
