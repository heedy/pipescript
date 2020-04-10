package pipescript

import (
	"errors"
	"time"
)

type Datapoint struct {
	Timestamp float64     `json:"t"`
	Duration  float64     `json:"dt"`
	Data      interface{} `json:"d"`
}

func (dp *Datapoint) Float() (float64, error) {
	f, ok := Float(dp.Data)
	if !ok {
		return f, errors.New("Can't convert data value to number")
	}
	return f, nil
}

func (dp *Datapoint) Int() (int64, error) {
	f, ok := Int(dp.Data)
	if !ok {
		return f, errors.New("Can't convert data value to number")
	}
	return f, nil
}

func (dp *Datapoint) Bool() (bool, error) {
	b, ok := Bool(dp.Data)
	if !ok {
		return b, errors.New("Can't use data value as boolean")
	}
	return b, nil
}

func (dp *Datapoint) String() (string, error) {
	b, ok := String(dp.Data)
	if !ok {
		return b, errors.New("Can't use data value as string")
	}
	return b, nil
}

func (dp *Datapoint) ToString() string {
	return ToString(dp.Data)
}

// Time returns the unix time object
func (dp *Datapoint) Time() time.Time {
	return time.Unix(0, int64(1e9*dp.Timestamp))
}

func (dp *Datapoint) MapElement(k string) (interface{}, error) {
	v, ok := dp.Data.(map[string]interface{})
	if !ok {
		return v, errors.New("Can't get key from non-object value")
	}
	vv, _ := v[k]
	return vv, nil
}

type DatapointArrayIterator struct {
	Datapoints []Datapoint
	i          int
}

func (d *DatapointArrayIterator) Next(out *Datapoint) (*Datapoint, error) {
	if d.i < len(d.Datapoints) {
		dp := d.Datapoints[d.i]
		d.i++
		return &dp, nil
	}
	return nil, nil
}

func NewDatapointArrayIterator(dpa []Datapoint) *DatapointArrayIterator {
	return &DatapointArrayIterator{
		Datapoints: dpa,
	}
}
