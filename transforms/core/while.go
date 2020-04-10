package core

import (
	"github.com/heedy/pipescript"
	"github.com/heedy/pipescript/resources"
)

// The argIterator ends iterating once the first arg is false
type whileArgIterator struct {
	args    []*pipescript.Datapoint
	e       *pipescript.TransformEnv
	startDP *pipescript.Datapoint
	done    bool
}

func (ai *whileArgIterator) Next(out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	if ai.startDP != nil {
		out.Timestamp = ai.startDP.Timestamp
		out.Duration = ai.startDP.Duration
		out.Data = ai.startDP.Data
		ai.startDP = nil
		return out, nil
	}
	dp, args, err := ai.e.Next(ai.args)
	if err != nil || dp == nil {
		ai.done = true
		return dp, err
	}
	argval, err := args[0].Bool()
	if err != nil {
		return nil, err
	}
	if !argval {
		ai.done = true
		ai.startDP = dp
		return nil, nil // The value was false. Pretend that the stream is finished
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Data
	return out, nil
}

type whileIter struct {
	args    []*pipescript.Datapoint
	pipe    *pipescript.Pipe
	startDP *pipescript.Datapoint
	done    bool
}

func (w *whileIter) OneToOne() bool {
	return false
}

func (w *whileIter) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {

	if w.done {
		return nil, nil
	}
	for {

		p := w.pipe.Copy()
		wi := &whileArgIterator{
			args:    w.args,
			e:       e,
			startDP: w.startDP,
		}
		p.InputIterator(wi)

		// Loop until finished
		dp, err := p.Next(out)
		if err != nil {
			return nil, err
		}

		dp2 := dp
		for dp != nil {
			dp2 = dp
			dp, err = p.Next(out)
			if err != nil {
				return nil, err
			}
		}
		var tmp pipescript.Datapoint
		for !wi.done {
			// The pipe finished before the iterator actually got to a false arg. Finish looping
			_, err := wi.Next(&tmp)
			if err != nil {
				return nil, err
			}
		}

		w.startDP = wi.startDP
		if w.startDP == nil {
			w.done = true
		}
		if dp2 != nil {
			// Return the last datapoint returned from the pipe
			return dp2, err
		}
		// Otherwise, keep looping through subsequent values
	}

}

var While = &pipescript.Transform{
	Name:          "while",
	Description:   "Equivalent to a while loop that runs while the first argument is true. Restarts the loop when the argument is false.",
	Documentation: string(resources.MustAsset("docs/transforms/while.md")),
	Args: []pipescript.TransformArg{
		{
			Description: "The statement to check for truth",
			Type:        pipescript.TransformArgType,
			Schema: map[string]interface{}{
				"type": "boolean",
			},
		},
		{
			Description: "transform to run, and to reset when the first arg is false",
			Type:        pipescript.PipeArgType,
		},
	},
	Constructor: func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		return &whileIter{args: make([]*pipescript.Datapoint, 1), pipe: pipes[0]}, nil
	},
}
