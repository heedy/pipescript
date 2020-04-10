package datasets

import (
	"math"

	"github.com/heedy/pipescript"
)

type mergeIterator struct {
	iterator  []pipescript.Iterator
	datapoint []*pipescript.Datapoint
}

func (m *mergeIterator) Next(out *pipescript.Datapoint) (dp *pipescript.Datapoint, err error) {
	//TODO: There are several inefficiencies in this implementation: First off, it is O(n), where
	//it can be made O(logn) by using a tree. Second, I just keep nulls in the array, which is
	//totally BS, the array could be made shorter when one iterator empties. But I just want to get this
	//thing working atm, so making it efficient is a task for later.
	mini := -1
	mint := math.Inf(-1)
	for i := range m.datapoint {
		//Iterators that are empty will be nil
		if m.datapoint[i] != nil {
			//Get the datapoint with smallest timestamp
			if m.datapoint[i].Timestamp < mint || mini == -1 {
				mini = i
				mint = m.datapoint[i].Timestamp
			}
		}
	}
	if mini == -1 {
		//There are no datapoints left
		return nil, nil
	}
	dp = m.datapoint[mini]
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Data

	m.datapoint[mini], err = m.iterator[mini].Next(dp)

	return out, err
}

// Merge takes multiple Iterators and merges them into one Iterator based upon
// increasing timestamp. Remember that everywhere in PipeScript, it is assumed that datapoints have
// increasing timestamp, that is, they are ordered by time.
func Merge(da []pipescript.Iterator) (pipescript.Iterator, error) {
	dpa := make([]*pipescript.Datapoint, 0, len(da))

	for i := range da {
		dp, err := da[i].Next(&pipescript.Datapoint{})
		if err != nil {
			return nil, err
		}
		dpa = append(dpa, dp)
	}

	return &mergeIterator{da, dpa}, nil
}
