package pipescript

// The Identity ($)
type iTransform struct{}

func (t iTransform) Copy() (TransformInstance, error) {
	return iTransform{}, nil
}

func (t iTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	return ti.Next().Get()
}

// A transform that returns a sub-object (assumes that object is struct or array)
type subobjectTransform struct {
	Obj interface{}
}

func (t *subobjectTransform) Copy() (TransformInstance, error) {
	return &subobjectTransform{t.Obj}, nil
}

func (t *subobjectTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		return te.Get()
	}

	//Datapoint.Get does exactly what we want here
	v, err := te.Datapoint.Get(t.Obj)
	if err != nil {
		return nil, err
	}
	return te.Set(v)
}

var identityTransform = Transform{
	Name:        "$",
	Description: "Identity transform - gives the current datapoint in sequence (or if given argument, a sub-object). Useful when performing comparisons ($ < 5)",
	OneToOne:    true,
	Stateless:   true,
	Args: []TransformArg{
		{
			Description: "The subobject to return. For use in json-object type data.",
			Optional:    true,
			Default:     nil,
			Constant:    true,
		},
	},
	Generator: func(name string, args []*Script) (*TransformInitializer, error) {
		// Get the subobject
		dp, err := args[0].GetConstant()
		if err != nil {
			return nil, err
		}

		if dp.Data == nil {
			// This means that the optional argument was not given
			return &TransformInitializer{Transform: iTransform{}}, nil
		}
		return &TransformInitializer{Transform: &subobjectTransform{dp.Data}}, nil
	},
}
