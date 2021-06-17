package numeric

import (
	"testing"

	"github.com/heedy/pipescript"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	p, err := pipescript.NewElementPipe(Sum, nil)
	require.NoError(t, err)

	p.InputIterator(pipescript.NewDatapointArrayIterator([]pipescript.Datapoint{
		{Data: 3},
		{Data: 2},
		{Data: 6},
		{Data: -1.0},
	}))

	dp, err := p.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.NotNil(t, dp)
	require.EqualValuesf(t, 10.0, dp.Data, "Sum must be summing")
	dp, err = p.Next(&pipescript.Datapoint{})
	require.NoError(t, err)
	require.Nil(t, dp)
}

func TestSum2(t *testing.T) {
	Sum.Register()
	pipescript.TestCase{
		Pipescript: "{'s': sum, 'i': d}",
		Input: []pipescript.Datapoint{
			{Timestamp: 1, Data: int64(2)},
			{Timestamp: 2, Data: 3},
			{Timestamp: 3, Data: 4},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1, Duration: 2, Data: map[string]interface{}{"s": float64(9), "i": 4}},
		},
	}.Run(t)

	pipescript.TestCase{
		Pipescript: "sum",
		Input:      []pipescript.Datapoint{},
		Output: []pipescript.Datapoint{
			{Timestamp: 0, Duration: 0, Data: float64(0)},
		},
	}.Run(t)
}

type testIterator struct {
	curi int
	maxi int
}

func (tn *testIterator) Next(dp *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if tn.curi >= tn.maxi && tn.maxi > 0 {
		return nil, nil
	}
	dp.Data = tn.curi
	dp.Timestamp = float64(tn.curi)
	tn.curi++
	return dp, nil
}

func BenchmarkSum(b *testing.B) {
	p, _ := pipescript.NewElementPipe(Sum, nil)

	p.InputIterator(&testIterator{maxi: b.N})
	p.Next(&pipescript.Datapoint{})
}
