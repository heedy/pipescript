package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type dayTransform struct {
	timezone  *time.Location
	startTime time.Time
}

func (t dayTransform) Copy() (pipescript.TransformInstance, error) {
	return dayTransform{t.timezone, t.startTime}, nil
}

func (t dayTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(int64(tval.Sub(t.startTime) / (24 * time.Hour)))
}

var Day = pipescript.Transform{
	Name:        "day",
	Description: "Returns the number of days since Jan 1 1970 in the given time zone.",
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
		return &pipescript.TransformInitializer{Transform: dayTransform{l, startTime}}, err
	},
}
