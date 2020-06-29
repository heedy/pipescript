package datasets

import (
	"fmt"

	"github.com/heedy/pipescript"
)

// Dataset allows to create a tabular structure
type Dataset struct {
	buf  *pipescript.Buffer
	iter *pipescript.BufferIterator

	data map[string]pipescript.Iterator
}

func NewDataset(stream pipescript.Iterator) *Dataset {
	ds := &Dataset{
		buf:  pipescript.NewBuffer(stream),
		data: make(map[string]pipescript.Iterator),
	}
	ds.iter = ds.buf.Iterator()
	return ds
}

func (d *Dataset) Reference() *pipescript.BufferIterator {
	return d.buf.Iterator()
}

func (d *Dataset) Add(key string, it pipescript.Iterator) {
	d.data[key] = it
}

func (d *Dataset) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	ref, err := d.iter.Next()
	if err != nil || ref == nil {
		return ref, err
	}
	data := make(map[string]interface{})
	for k, it := range d.data {
		dp, err := it.Next(out)
		if err != nil {
			return nil, err
		}
		if dp == nil {
			return nil, fmt.Errorf("Dataset interpolator for '%s' finished early. This is a bug!", k)
		}
		data[k] = dp.Data
	}
	out.Timestamp = ref.Timestamp
	out.Duration = ref.Duration
	out.Data = data
	return out, nil
}

type timeRangeIterator struct {
	T  float64
	Dt float64
	T2 float64
}

func (tri *timeRangeIterator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if tri.T >= tri.T2 {
		return nil, nil
	}
	out.Timestamp = tri.T
	// out.Duration = tri.Dt
	tri.T += tri.Dt
	return out, nil

}

func NewTDataset(t1, t2, dt float64) *Dataset {
	return NewDataset(&timeRangeIterator{
		T:  t1,
		T2: t2,
		Dt: dt,
	})
}
