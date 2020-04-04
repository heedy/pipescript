package core

import (
	"errors"

	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

type arrIterator struct {
	Timestamp float64
	Duration  float64
	vals      []interface{}
	idx       int
}

func (ai *arrIterator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if ai.idx >= len(ai.vals) {
		return nil, nil
	}
	out.Timestamp = ai.Timestamp
	out.Duration = ai.Duration
	out.Data = ai.vals[ai.idx]
	ai.idx++
	return out, nil
}

var Reduce = &pipescript.Transform{
	Name:        "reduce",
	Description: "Takes a json object, and considers each field to be a separate datapoint's data. It then runs the transform in its argument over the elements",
	Args: []pipescript.TransformArg{
		{
			Description: "The transform to instantiate for each datapoint's values.",
			Type:        pipescript.PipeArgType,
		},
	},
	Documentation: string(resources.MustAsset("docs/transforms/reduce.md")),
	Constructor: pipescript.NewBasic(nil, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		ai := &arrIterator{
			Timestamp: dp.Timestamp,
			Duration:  dp.Duration,
		}
		switch d := dp.Data.(type) {
		case []interface{}:
			ai.vals = d
		case map[string]interface{}:
			ai.vals = make([]interface{}, 0, len(d))
			for _, v := range d {
				ai.vals = append(ai.vals, v)
			}
		default:
			return nil, errors.New("Can't reduce non-object/array datapoint")
		}
		p := pipes[0].Copy()
		p.InputIterator(ai)

		var data interface{}
		for {
			dd, err := p.Next(out)
			if err != nil {
				return nil, err
			}
			if dd != nil {
				data = dd.Data
			} else {
				break
			}
		}
		out.Timestamp = dp.Timestamp
		out.Duration = dp.Duration
		out.Data = data
		return out, nil
	}),
}
