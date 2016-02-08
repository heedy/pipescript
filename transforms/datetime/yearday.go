package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type yeardayTransform struct {
	timezone *time.Location
}

func (t yeardayTransform) Copy() (pipescript.TransformInstance, error) {
	return yeardayTransform{t.timezone}, nil
}

func (t yeardayTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(tval.YearDay())
}

var Yearday = pipescript.Transform{
	Name:        "yearday",
	Description: "Returns the day of the year during which the datapoint happened [1,365] non-leap, and [1,366] for leap years.",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		return &pipescript.TransformInitializer{Transform: yeardayTransform{l}}, err
	},
}
