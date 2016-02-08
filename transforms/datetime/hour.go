package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type hourTransform struct {
	timezone  *time.Location
	startTime time.Time
}

func (t hourTransform) Copy() (pipescript.TransformInstance, error) {
	return hourTransform{t.timezone, t.startTime}, nil
}

func (t hourTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(int64(tval.Sub(t.startTime) / (time.Hour)))
}

var Hour = pipescript.Transform{
	Name:        "hour",
	Description: "Returns the number of hours since Jan 1 1970 in the given time zone.",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		if err != nil {
			return nil, err
		}
		startTime := time.Date(1970, time.January, 1, 0, 0, 0, 0, l)
		return &pipescript.TransformInitializer{Transform: hourTransform{l, startTime}}, err
	},
}
