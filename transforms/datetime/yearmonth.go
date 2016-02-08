package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type yearmonthTransform struct {
	timezone *time.Location
}

func (t yearmonthTransform) Copy() (pipescript.TransformInstance, error) {
	return yearmonthTransform{t.timezone}, nil
}

func (t yearmonthTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(tval.Month().String())
}

var Yearmonth = pipescript.Transform{
	Name:        "yearmonth",
	Description: "Returns the month name during which the datapoint happened ('January','February'...)",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		l, err := getTimezone(args[0])
		return &pipescript.TransformInitializer{Transform: yearmonthTransform{l}}, err
	},
}
