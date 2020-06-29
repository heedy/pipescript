package pipescript

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testIterator struct {
	curi int
	maxi int
}

func (tn *testIterator) Next(dp *Datapoint) (*Datapoint, error) {
	if tn.curi >= tn.maxi && tn.maxi > 0 {
		return nil, nil
	}
	dp.Data = tn.curi
	dp.Timestamp = float64(tn.curi)
	tn.curi++
	return dp, nil
}

func TestBuffer(t *testing.T) {
	b := NewBuffer(&testIterator{})
	it1 := b.Iterator()
	it2 := b.Iterator()
	dp1, err := it1.Next()
	require.NoError(t, err)
	dp2, err := it2.Peek(0)
	require.NoError(t, err)
	require.Equal(t, dp1, dp2)
	require.Equal(t, 0, dp1.Data)
	dp1, err = it1.Next()
	require.NoError(t, err)
	require.Equal(t, 1, dp1.Data)

	dp1, err = it1.Peek(5000)
	require.NoError(t, err)
	require.Equal(t, 5002, dp1.Data)

	_, err = it1.Peek(-500)
	require.EqualError(t, err, ErrBeforeStart.Error())

	dp1, err = it1.Peek(-2)
	require.NoError(t, err)
	require.Equal(t, 0, dp1.Data)

	for i := 0; i < 1000; i++ {
		dp1, err = it1.Next()
		require.NoError(t, err)
		require.Equal(t, i+2, dp1.Data)
	}

	dp2, err = it2.Peek(0)
	require.NoError(t, err)
	require.Equal(t, 0, dp2.Data)

	dp1, err = it1.Peek(-992)
	require.NoError(t, err)
	require.Equal(t, 10, dp1.Data)

	// Now move the second iterator, so that old pages can be reused
	for i := 0; i < 7000; i++ {
		dp1, err = it2.Next()
		require.NoError(t, err)
		require.Equal(t, i, dp1.Data)
	}

	// Make sure that the old pages were reused
	dp1, err = it1.Peek(-992)
	require.Error(t, err)
	// But still make sure that some old values are still available
	dp1, err = it1.Peek(-2)
	require.NoError(t, err)
}

func TestBufferClose(t *testing.T) {
	b := NewBuffer(&testIterator{})
	it1 := b.Iterator()
	it2 := b.Iterator()
	require.Len(t, b.Iterators, 2)
	it1.Close()
	require.Len(t, b.Iterators, 1)
	require.Equal(t, b.Iterators[0], it2)
	it2.Close()
	require.Len(t, b.Iterators, 0)
}

func BenchmarkBuffer(b *testing.B) {
	buf := NewBuffer(&testIterator{maxi: b.N})
	it := buf.Iterator()
	dp, _ := it.Next()
	for dp != nil {
		dp, _ = it.Next()
	}
}
