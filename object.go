package pipescript

import "errors"

// objectTransform is used when creating a json-like object in pipescript
type objectTransform struct {
	obj  map[string]*Script
	iter *SingleDatapointIterator
}

func (o objectTransform) Copy() (TransformInstance, error) {
	newmap := make(map[string]*Script)
	iter := &SingleDatapointIterator{}
	for key := range o.obj {
		val, err := o.obj[key].Copy()
		if err != nil {
			return nil, err
		}
		val.SetInput(iter)
		newmap[key] = val
	}
	return objectTransform{newmap, iter}, nil
}

// Next performs an and on two boolean datapoints
func (o objectTransform) Next(ti *TransformIterator) (*Datapoint, error) {
	te := ti.Next()
	if te.IsFinished() {
		// We need to clear the maps
		o.iter.Set(nil, nil)
		for key := range o.obj {
			_, err := o.obj[key].Next()
			if err != nil {
				return nil, err
			}
		}
		return te.Get()
	}

	result := make(map[string]interface{})
	o.iter.Set(te.Datapoint, nil)
	for key := range o.obj {
		dp, err := o.obj[key].Next()
		if err != nil {
			return nil, err
		}
		result[key] = dp.Data
	}

	return te.Set(result)
}

func newObjectTransform(input map[string]*Script) (*Script, error) {
	iter := &SingleDatapointIterator{}
	isConstant := true
	for key := range input {
		if input[key].Peek {
			return nil, errors.New("All transforms in object must not peek")
		}
		if !input[key].OneToOne {
			return nil, errors.New("All transforms in object must be one-to-one")
		}
		if !input[key].Constant {
			isConstant = false
		}
		input[key].SetInput(iter)
	}
	pe, err := NewPipelineElement([]*Script{}, objectTransform{input, iter})

	return &Script{
		input:    pe,
		output:   pe,
		OneToOne: true,
		Constant: isConstant,
	}, err

}
