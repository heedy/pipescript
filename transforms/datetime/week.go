package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type weekTransform struct {
	timezone  *time.Location
	startTime time.Time
}

func (t weekTransform) Copy() (pipescript.TransformInstance, error) {
	return weekTransform{t.timezone, t.startTime}, nil
}

func (t weekTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(int64(tval.Sub(t.startTime) / (7 * 24 * time.Hour)))
}

var Week = pipescript.Transform{
	Name:        "week",
	Description: "Returns the number of weeks since Jan 5 1970 (First Monday after unix time) in the given time zone.",
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
		startTime := time.Date(1970, time.January, 5, 0, 0, 0, 0, l)
		return &pipescript.TransformInitializer{Transform: weekTransform{l, startTime}}, err
	},
}
