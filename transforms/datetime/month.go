package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type monthTransform struct {
	timezone *time.Location
}

func (t monthTransform) Copy() (pipescript.TransformInstance, error) {
	return monthTransform{t.timezone}, nil
}

func (t monthTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(int64(12*(tval.Year()-1970) + int(tval.Month())))
}

var Month = pipescript.Transform{
	Name:        "month",
	Description: "Returns the number of months since Jan 1970 in the given time zone.",
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
		return &pipescript.TransformInitializer{Transform: monthTransform{l}}, err
	},
}
