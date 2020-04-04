package pipescript

import (
	"encoding/json"
)

type AggregatorFunc func(e *TransformEnv, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error)
type Aggregator struct {
	ConstArgs []interface{}
	PipeArgs  []*Pipe

	isFinished bool
	f          AggregatorFunc
}

func (a *Aggregator) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	if a.isFinished {
		return nil, nil
	}
	a.isFinished = true
	return a.f(e, a.ConstArgs, a.PipeArgs, out)
}

func (a *Aggregator) OneToOne() bool {
	return false
}

func NewAggregator(f AggregatorFunc) TransformConstructor {
	return func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		return &Aggregator{
			ConstArgs:  consts,
			PipeArgs:   pipes,
			f:          f,
			isFinished: false,
		}, nil
	}
}

type BasicFunc func(dp *Datapoint, args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error)
type BasicInit func(consts []interface{}, pipes []*Pipe) ([]interface{}, []*Pipe, error)

type Basic struct {
	ConstArgs []interface{}
	PipeArgs  []*Pipe
	Args      []*Datapoint
	f         BasicFunc
}

func (s *Basic) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, arr, err := e.Next(s.Args)
	if err != nil || dp == nil {
		return nil, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	return s.f(dp, arr, s.ConstArgs, s.PipeArgs, out)
}

func (s *Basic) OneToOne() bool {
	return true
}

func (s *Basic) GetConst(dpi interface{}, args []interface{}) (interface{}, error) {
	ind := &Datapoint{Data: dpi}
	argd := make([]*Datapoint, len(args))
	for i := range args {
		argd[i] = &Datapoint{Data: args[i]}
	}
	dp, err := s.f(ind, argd, s.ConstArgs, s.PipeArgs, &Datapoint{})
	if err != nil {
		return nil, err
	}
	return dp.Data, nil
}

func NewBasic(i BasicInit, f BasicFunc) TransformConstructor {
	return func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		var err error
		if i != nil {
			consts, pipes, err = i(consts, pipes)
		}
		return &Basic{
			ConstArgs: consts,
			PipeArgs:  pipes,
			f:         f,
			Args:      make([]*Datapoint, len(transform.Args)),
		}, err
	}
}

type ArgBasicFunc func(args []*Datapoint, consts []interface{}, pipes []*Pipe, out *Datapoint) (*Datapoint, error)

type ArgBasic struct {
	ConstArgs []interface{}
	PipeArgs  []*Pipe
	Args      []*Datapoint
	f         ArgBasicFunc
}

func (s *ArgBasic) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, arr, err := e.Next(s.Args)
	if err != nil || dp == nil {
		return nil, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	return s.f(arr, s.ConstArgs, s.PipeArgs, out)
}

func (s *ArgBasic) OneToOne() bool {
	return true
}

func (s *ArgBasic) GetConst(args []interface{}) (interface{}, error) {
	argd := make([]*Datapoint, len(args))
	for i := range args {
		argd[i] = &Datapoint{Data: args[i]}
	}
	dp, err := s.f(argd, s.ConstArgs, s.PipeArgs, &Datapoint{})
	if err != nil {
		return nil, err
	}
	return dp.Data, nil
}

func NewArgBasic(f ArgBasicFunc) TransformConstructor {
	return func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
		return &ArgBasic{
			ConstArgs: consts,
			PipeArgs:  pipes,
			f:         f,
			Args:      make([]*Datapoint, len(transform.Args)),
		}, nil
	}
}

type ConstIterator struct {
	Value interface{}
}

func (c *ConstIterator) Next(e *TransformEnv, out *Datapoint) (*Datapoint, error) {
	dp, _, err := e.Next(nil)
	if dp == nil || err != nil {
		return dp, err
	}
	out.Data = c.Value
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	return out, nil
}

func (c *ConstIterator) OneToOne() bool {
	return true
}

func NewConstTransform(v interface{}) *Transform {
	b, _ := json.Marshal(v)
	return &Transform{
		Name:        string(b),
		Description: "Constant value",
		Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
			return &ConstIterator{Value: v}, nil
		},
	}
}
