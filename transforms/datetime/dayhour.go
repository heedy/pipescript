package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type dayhourTransform struct {
	timezone *time.Location
}

func (t dayhourTransform) Copy() (pipescript.TransformInstance, error) {
	return dayhourTransform{t.timezone}, nil
}

func (t dayhourTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(tval.Hour())
}

var Dayhour = pipescript.Transform{
	Name:        "dayhour",
	Description: "Returns the hour in a day during which the datapoint happened.",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		return &pipescript.TransformInitializer{Transform: dayhourTransform{l}}, err
	},
}
