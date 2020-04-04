package pipescript

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TransformEnv struct {
	Iter     *BufferIterator
	ArgIters []*BufferIterator
}

func (te *TransformEnv) Next(args []*Datapoint) (*Datapoint, []*Datapoint, error) {
	dp, err := te.Iter.Next()
	if err == nil && dp != nil {
		for i := range te.ArgIters {
			args[i], err = te.ArgIters[i].Next()
			if err != nil {
				return dp, args, err
			}
			if args[i] == nil {
				return nil, args, nil
			}
		}
	}
	return dp, args, err
}

func (te *TransformEnv) Peek(idx int, args []*Datapoint) (*Datapoint, []*Datapoint, error) {
	dp, err := te.Iter.Peek(idx)
	if err == nil && dp != nil {
		for i := range te.ArgIters {
			args[i], err = te.ArgIters[i].Peek(idx)
			if err != nil {
				return dp, args, err
			}
			if args[i] == nil {
				return nil, args, nil
			}
		}
	}
	return dp, args, err
}

type PipeElement struct {
	Env  *TransformEnv
	Args []*Pipe

	Iter TransformIterator

	// These three are here to allow copying a pipe by re-generating it from scratch
	ConstArgs []interface{}
	PipeArgs  []*Pipe
	Transform *Transform
}

func NewPipeElement(t *Transform, args []*Pipe) (*PipeElement, error) {
	targs := make([]*Pipe, 0)
	consts := make([]interface{}, 0)
	pipes := make([]*Pipe, 0)

	if len(args) > len(t.Args) {
		return nil, fmt.Errorf("Transform '%s' takes %d arguments, but %d given", t.Name, len(t.Args), len(args))
	}

	for i := range t.Args {
		if len(args) <= i {
			if t.Args[i].Optional {
				args = append(args, t.Args[i].Default)
			} else {
				return nil, fmt.Errorf("Transform '%s' requires additional arguments", t.Name)
			}
		}
		args[i].Simplify()

		// Now split the args up
		switch t.Args[i].Type {
		case ConstArgType:
			cv, err := args[i].GetConst()
			if err != nil {
				return nil, err
			}
			consts = append(consts, cv)
		case TransformArgType:
			targs = append(targs, args[i])
		case OneToOnePipeArgType:
			pipes = append(pipes, args[i])
			if !args[i].OneToOne() {
				return nil, fmt.Errorf("Transform %s arg %d (%s) is not one-to-one", t.Name, i, args[i].String())
			}
		case PipeArgType:
			pipes = append(pipes, args[i])
		}
	}

	pe := &PipeElement{
		Transform: t,
		Args:      targs,
		ConstArgs: consts,
		PipeArgs:  pipes,
		Env: &TransformEnv{
			ArgIters: make([]*BufferIterator, len(targs)),
		},
	}

	// Set up the arg iterators
	for i := range targs {
		pe.Env.ArgIters[i] = NewBuffer(targs[i]).Iterator()
	}

	var err error
	pe.Iter, err = t.Constructor(t, consts, pipes)

	return pe, err
}

func (pe *PipeElement) String() string {
	s := pe.Transform.Name
	if len(pe.Transform.Args) > 0 {
		s += "("
		tai := 0
		pai := 0
		cai := 0
		for i := range pe.Transform.Args {
			switch pe.Transform.Args[i].Type {
			case ConstArgType:
				b, _ := json.Marshal(pe.ConstArgs[cai])
				s += string(b)
				cai++
			case TransformArgType:
				s += pe.Args[tai].String()
				tai++
			case PipeArgType, OneToOnePipeArgType:
				s += pe.PipeArgs[pai].String()
				pai++
			}
			s += ","
		}
		s = s[:len(s)-1] + ")"
	}
	return s
}

func (pe *PipeElement) GetConst(in interface{}) (interface{}, error) {
	ci, ok := pe.Iter.(*ConstIterator)
	if ok {
		return ci.Value, nil
	}
	inputargs := make([]interface{}, len(pe.Args))
	var err error
	for i := range pe.Args {
		inputargs[i], err = pe.Args[i].GetConst()
		if err != nil {
			return nil, fmt.Errorf("Transform '%s' arg %d: %w", pe.String(), i, err)
		}
	}

	bt, ok := pe.Iter.(*Basic)
	if !ok {
		abt, ok := pe.Iter.(*ArgBasic)
		if !ok {
			return nil, fmt.Errorf("Output of transform '%s' cannot be used as a constant", pe.String())
		}

		return abt.GetConst(inputargs)
	}

	return bt.GetConst(in, inputargs)
}

func (pe *PipeElement) Copy() *PipeElement {
	newArgs := make([]*Pipe, len(pe.Args))
	for i := range pe.Args {
		newArgs[i] = pe.Args[i].Copy()
	}
	newPipeArgs := make([]*Pipe, len(pe.PipeArgs))
	for i := range pe.PipeArgs {
		newPipeArgs[i] = pe.PipeArgs[i].Copy()
	}
	pnew := &PipeElement{
		Transform: pe.Transform,
		Args:      newArgs,
		ConstArgs: pe.ConstArgs,
		PipeArgs:  newPipeArgs,
		Env: &TransformEnv{
			ArgIters: make([]*BufferIterator, len(newArgs)),
		},
	}

	// Set up the arg iterators
	for i := range newArgs {
		pnew.Env.ArgIters[i] = NewBuffer(newArgs[i]).Iterator()
	}
	var err error
	pnew.Iter, err = pe.Transform.Constructor(pnew.Transform, pnew.ConstArgs, pnew.PipeArgs)
	if err != nil {
		panic(err)
	}
	return pnew
}

func (pe *PipeElement) Input(b *Buffer) {
	// Set the root iterator
	pe.Env.Iter = b.Iterator()

	// Set the root iterators of all args
	for i := range pe.Args {
		pe.Args[i].Input(b)
	}
}

func (pe *PipeElement) IsBasic() bool {
	for _, ap := range pe.Args {
		if !ap.IsBasic() {
			return false
		}
	}
	return true
}

func (pe *PipeElement) Next(out *Datapoint) (*Datapoint, error) {
	return pe.Iter.Next(pe.Env, out)
}

type Pipe struct {
	Arr []*PipeElement
}

func NewPipe() *Pipe {
	return &Pipe{
		Arr: make([]*PipeElement, 0),
	}
}

func NewElementPipe(t *Transform, args []*Pipe) (*Pipe, error) {
	p := NewPipe()
	el, err := NewPipeElement(t, args)
	p.Append(el)
	return p, err
}

func NewTransformPipe(tname string, args []*Pipe) (*Pipe, error) {
	RegistryLock.RLock()
	defer RegistryLock.RUnlock()
	t, ok := TransformRegistry[tname]
	if !ok {
		return nil, fmt.Errorf("Could not find transform '%s'", tname)
	}
	return NewElementPipe(t, args)
}

// Input sets the pipe's input to the given buffer
func (p Pipe) Input(b *Buffer) {
	if len(p.Arr) > 0 {
		p.Arr[0].Input(b)
	}
}

// InputNexter uses an iterator interface as input
func (p *Pipe) InputIterator(n Iterator) {
	p.Input(NewBuffer(n))
}

func (p *Pipe) Next(dp *Datapoint) (*Datapoint, error) {
	return p.Arr[len(p.Arr)-1].Next(dp)
}

// Append returns a pipe with the given PipeElement added
func (p *Pipe) Append(e *PipeElement) {
	plen := len(p.Arr)
	if plen > 0 {
		e.Input(NewBuffer(p.Arr[plen-1]))
	}
	p.Arr = append(p.Arr, e)
}

func (p *Pipe) Copy() *Pipe {
	p2 := NewPipe()
	for i := range p.Arr {
		p2.Append(p.Arr[i].Copy())
	}
	return p2
}

func (p *Pipe) Join(p2 *Pipe) {
	plen := len(p.Arr)
	if plen > 0 {
		p2.InputIterator(p.Arr[plen-1])
	}
	p.Arr = append(p.Arr, p2.Arr...)
}

func (p *Pipe) OneToOne() bool {
	for _, e := range p.Arr {
		if !e.Iter.OneToOne() {
			return false
		}
	}
	return true
}

func (p *Pipe) String() string {
	if len(p.Arr) == 0 {
		return ""
	}
	s := p.Arr[0].String()
	for i := 1; i < len(p.Arr); i++ {
		s += ":" + p.Arr[i].String()
	}
	return s
}

func (p *Pipe) IsBasic() bool {
	for i := range p.Arr {
		switch p.Arr[i].Iter.(type) {
		case *ConstIterator, peekIterator:
			// Do nothing
		case *Basic, *ArgBasic:
			if !p.Arr[i].IsBasic() {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func (p *Pipe) Simplify() *Pipe {
	arr2 := make([]*PipeElement, 0, len(p.Arr))

	// Check which ArgBasics can be turned into consts
	for i := 0; i < len(p.Arr); i++ {
		switch p.Arr[i].Iter.(type) {
		case *ArgBasic:
			v, err := p.Arr[i].GetConst(nil)
			if err == nil {
				pe, err := NewPipeElement(NewConstTransform(v), nil)
				if err != nil {
					panic(err)
				}
				p.Arr[i] = pe
			}
		default:
			// Nothing
		}
	}

	// Remove redundant transforms from the pipe
	for i := 0; i < len(p.Arr); i++ {
		switch v := p.Arr[i].Iter.(type) {
		case *ConstIterator:
			// A const replaces all previous basic transforms
			j := len(arr2) - 1
			for ; j >= 0; j-- {
				switch v := arr2[j].Iter.(type) {
				case *ConstIterator:
				case *Basic, *ArgBasic:
					if !arr2[j].IsBasic() {
						break
					}
				case peekIterator:
					if v.Peek != 0 {
						// Need to stop here, since must include the peek
						break
					}
				default:
					break
				}
			}
			arr2 = append(arr2[:j+1], p.Arr[i])
		case peekIterator:
			// If not peeking, can remove it entirely
			if v.Peek != 0 {
				arr2 = append(arr2, p.Arr[i])
			}
		default:
			arr2 = append(arr2, p.Arr[i])
		}

	}
	p.Arr = make([]*PipeElement, 0, len(arr2))

	// Second pass: propagate constants forward where possible. No longer
	// need to worry about peeks anymore
	for i := 0; i < len(arr2); i++ {
		switch v := arr2[i].Iter.(type) {
		case *ConstIterator:
			cv := v.Value
			for i++; i < len(arr2); i++ {
				v, err := arr2[i].GetConst(cv)
				if err != nil {
					break
				}
				cv = v
			}
			pe, err := NewPipeElement(NewConstTransform(cv), nil)
			if err != nil {
				panic(err)
			}
			p.Append(pe)
			if i < len(arr2) {
				i--
			}
		default:
			p.Append(arr2[i])
		}
	}

	if len(p.Arr) == 0 {
		// If all elements were removed, add a basic 0 peek iterator back
		pe, err := NewPipeElement(&Transform{
			Name: "$",
			Constructor: func(transform *Transform, consts []interface{}, pipes []*Pipe) (TransformIterator, error) {
				return peekIterator{0}, nil
			},
		}, nil)
		if err != nil {
			panic(err)
		}
		p.Arr = []*PipeElement{pe}
	}
	return p
}

func (p *Pipe) GetConst() (interface{}, error) {
	// assumes that was already simplified
	if len(p.Arr) == 1 {
		ci, ok := p.Arr[0].Iter.(*ConstIterator)
		if ok {
			return ci.Value, nil
		}
	}
	return nil, fmt.Errorf("Cannot use %s as constant value", p.String())
}

type chanIterator struct {
	Receiver chan *Datapoint
}

func (ci chanIterator) Next(out *Datapoint) (*Datapoint, error) {
	dp, ok := <-ci.Receiver
	if !ok {
		return nil, errors.New("Input channel closed")
	}
	if dp == nil {
		return nil, nil
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration
	out.Data = dp.Data
	return out, nil
}

type ChanResult struct {
	DP  *Datapoint
	Err error
}

type ChannelPipe struct {
	Sender   chan *Datapoint
	Receiver chan ChanResult
}

func (cp *ChannelPipe) Close() {
	close(cp.Sender)
}

func NewChannelPipe(p *Pipe) *ChannelPipe {
	sender := make(chan *Datapoint)
	receiver := make(chan ChanResult)
	cp := &ChannelPipe{sender, receiver}

	go func() {
		p.InputIterator(chanIterator{sender})
		for {
			dp, err := p.Next(&Datapoint{}) // Make sure to create new datapoint each time
			receiver <- ChanResult{dp, err}
			if dp == nil || err != nil {
				break
			}
		}
		close(receiver)
	}()

	return cp
}

// Parse parses the given transform, and returns the corresponding pipe
func Parse(script string) (*Pipe, error) {
	lexer := parserLex{input: script}

	if parserParse(&lexer) != 0 {
		if lexer.errorString != "" {
			return nil, fmt.Errorf("'%s': %s", script, lexer.errorString)
		}
		return nil, fmt.Errorf("'%s': Unknown error", script)
	}
	lexer.output.Simplify()

	return lexer.output, nil
}

func MustPipe(t *Transform, args []*Pipe) *Pipe {
	p, err := NewElementPipe(t, args)
	if err != nil {
		panic(err)
	}
	return p
}

func MustParse(script string) *Pipe {
	p, err := Parse(script)
	if err != nil {
		panic(err)
	}
	return p
}
