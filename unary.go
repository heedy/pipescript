package pipescript

type notTransform struct{}

func (t notTransform) Copy() TransformInstance {
	return notTransform{}
}

func (t notTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	v, err := te.Args[0].Bool()
	if err != nil {
		return nil, err
	}
	return te.Set(!v)
}

func notScript(a *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a}, notTransform{})
	return &Script{
		input:     pe,
		output:    pe,
		OneToOne:  true,
		Constant:  a.Constant,
		Stateless: true,
	}, err
}

// negativeTransform performs the unary minus
type negativeTransform struct{}

func (t negativeTransform) Copy() TransformInstance {
	return negativeTransform{}
}

func (t negativeTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	// There is a special case: if we do negative boolean, we behave as a not
	b, ok := te.Args[0].Data.(bool)
	if ok {
		return te.Set(!b)
	}
	// We just deal with numbers, so float is the same thing as int for our purposes
	v, err := te.Args[0].Float()
	if err != nil {
		return nil, err
	}
	return te.Set(-v)
}

func negativeScript(a *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a}, negativeTransform{})
	return &Script{
		input:     pe,
		output:    pe,
		OneToOne:  true,
		Constant:  a.Constant,
		Stateless: true,
	}, err
}
