package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type monthdayTransform struct {
	timezone *time.Location
}

func (t monthdayTransform) Copy() (pipescript.TransformInstance, error) {
	return monthdayTransform{t.timezone}, nil
}

func (t monthdayTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(tval.Day())
}

var Monthday = pipescript.Transform{
	Name:        "monthday",
	Description: "Returns the number of day in the datapoint's month (int)",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		return &pipescript.TransformInitializer{Transform: monthdayTransform{l}}, err
	},
}
