package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type weekdayTransform struct {
	timezone *time.Location
}

func (t weekdayTransform) Copy() (pipescript.TransformInstance, error) {
	return weekdayTransform{t.timezone}, nil
}

func (t weekdayTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(tval.Weekday().String())
}

var Weekday = pipescript.Transform{
	Name:        "weekday",
	Description: "Returns the weekday during which the datapoint happened ('Monday','Tuesday'...)",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		return &pipescript.TransformInitializer{Transform: weekdayTransform{l}}, err
	},
}
