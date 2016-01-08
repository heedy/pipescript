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
		input:     pe,
		output:    pe,
		OneToOne:  true,
		Constant:  true,
		Stateless: true,
	}
}
