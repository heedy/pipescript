package core

import "github.com/connectordb/pipescript"

type tshiftTransform struct {
	shiftby float64
}

func (t tshiftTransform) Copy() pipescript.TransformInstance {
	return tshiftTransform{t.shiftby}
}

func (t tshiftTransform) Next(ti *pipescript.TransformIterator) (*pipescript.Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	return &pipescript.Datapoint{Timestamp: te.Datapoint.Timestamp + t.shiftby, Data: te.Datapoint.Data}, nil
}

var tshift = pipescript.Transform{
	Name:        "tshift",
	Description: "Shift the datapoint timestamp by a constant number of seconds",
	OneToOne:    true,
	Stateless:   true,
	Args: []pipescript.TransformArg{
		{
			Description: "The number of seconds to shift the timestamp",
			Constant:    true,
		},
	},
	Generator: func(name string, args []*pipescript.Script) (*pipescript.TransformInitializer, error) {
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}
		val, err := dp.Float()
		return &pipescript.TransformInitializer{Transform: tshiftTransform{val}}, err
	},
}
