package pipescript

// ConstantTransform is the transform that always returns a constant
type ConstantTransform struct {
	C interface{}
}

// Copy generates another
func (c ConstantTransform) Copy() TransformInstance {
	return ConstantTransform{c.C}
}

// Next returns a Constant datapoint whose timestamp is the current timestamp, but whose data is
func (c ConstantTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	return ti.Next().Set(c.C)
}

// Creates a script with the given constant in a ConstantTransform
func ConstantScript(c interface{}) *Script {
	pe, _ := NewPipelineElement(nil, ConstantTransform{c})
	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   true,
	}
}

// andTransform is used when given (a and b)
type andTransform struct{}

func (a andTransform) Copy() TransformInstance {
	return andTransform{}
}

// Next performs an and on two boolean datapoints
func (a andTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	b1, err := te.Args[0].Bool()
	if err != nil {
		return nil, err
	}
	b2, err := te.Args[1].Bool()
	if err != nil {
		return nil, err
	}
	return te.Set(b1 && b2)
}

func andScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, andTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err

}

// orTransform is used when given (a or b)
type orTransform struct{}

func (o orTransform) Copy() TransformInstance {
	return andTransform{}
}

// Next performs an or on two boolean datapoints
func (o orTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}
	b1, err := te.Args[0].Bool()
	if err != nil {
		return nil, err
	}
	b2, err := te.Args[1].Bool()
	if err != nil {
		return nil, err
	}
	return te.Set(b1 || b2)
}

func orScript(a1, a2 *Script) (*Script, error) {
	pe, err := NewPipelineElement([]*Script{a1, a2}, orTransform{})

	return &Script{
		input:      pe,
		output:     pe,
		IsOneToOne: true,
		Constant:   a1.Constant && a2.Constant,
	}, err

}
