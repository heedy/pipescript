package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

type yearTransform struct {
	timezone *time.Location
}

func (t yearTransform) Copy() (pipescript.TransformInstance, error) {
	return yearTransform{t.timezone}, nil
}

func (t yearTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	tval := te.Datapoint.Time().In(t.timezone)
	return te.Set(int64(tval.Year()))
}

var Year = pipescript.Transform{
	Name:        "year",
	Description: "Returns the year of the current timestamp",
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
		return &pipescript.TransformInitializer{Transform: yearTransform{l}}, err
	},
}
