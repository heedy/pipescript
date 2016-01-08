package pipescript

import (
	"container/list"
	"errors"
)

// PipelineElement is an element which holds all the necessary information to perform a transform.
// It conforms to the DatapointIterator interface.
// Internally, it uses a TransformIterator to hold the information.
// The main difficulty here is how to set up the datapoint propagation. This is difficult due to
// the fact that while the overall "program" is a chain of commands:
//	a | b | c | d
// each element of the chain can have a whole subtree of chains (as arguments to the core element)
// In order to construct this chain, we need to start from bottom up. That is, start at the leaf
// chains and build up to the root chain. Due to the construction of the chain, when a PipelineElement
// is created, it does not yet have access to the DatapointIterator
type PipelineElement struct {
	iter      TransformIterator
	transform TransformInstance
}

// Next allows a PipelineElement to conform to the DatapointIterator interface
func (p *PipelineElement) Next() (*Datapoint, error) {
	return p.transform.Next(&p.iter)
}

// SetInput links this PipelineElement to a DatapointPeekIterator from which all future data will come.
// Note that PipelineElements extensively use DatapointPeekIterators and caching, so changing the Input Iterator
// during runtime is NOT supported. The only case where this is possible is when the input iterator is finished,
// meaning that it does not have any data left.
func (p *PipelineElement) SetInput(d DatapointPeekIterator) {
	p.iter.iterator = d

	// Next, we create the virtual PeekIterator, and set it as input for all of the arguments
	p.iter.argpeeker = NewVirtualPeekIterator(d)
	for i := range p.iter.args {
		p.iter.args[i].SetInput(p.iter.argpeeker)
	}
}

// copyUntil generates a recursive copy of a pipelineElement (going back through its parent iterators) until it
// reaches the passed in pipelineElement (inclusive), and returns a copy of itself, as well as a pointer to the
// starting element (equivalent to the passed in parameter for the copied pipelineElement)
func (p *PipelineElement) copyUntil(pe *PipelineElement) (input, output *PipelineElement, err error) {
	// First copy the TransformInstance
	tInstance := p.transform.Copy()

	// Next get copies of all of the args
	args := make([]*Script, 0, len(p.iter.args))
	for i := range p.iter.args {
		argscript, err := p.iter.args[i].Copy()
		if err != nil {
			return nil, nil, err
		}
		args = append(args, argscript)
	}

	// Now generate a new PipelineElement from this element
	output, err = NewPipelineElement(args, tInstance)
	if err != nil {
		return nil, nil, err
	}

	// We have the current PipelineElement, but we actually have to perform a recursive step
	// since the DatapointIterator inside the TransformIterator is usually the previous element in a script chain
	if p == pe {
		// The base case is when the chain is just one element
		return output, output, nil
	}

	errIncompatible := errors.New("Pipeline copy failed due to incompatible type")

	// We call copyUntil recursively
	peeker, ok := p.iter.iterator.(*datapointPeekIterator)
	if !ok {
		return nil, nil, errIncompatible
	}

	prevPipeline, ok := peeker.Iterator.(*PipelineElement)
	if !ok {
		return nil, nil, errIncompatible
	}

	var o *PipelineElement
	input, o, err = prevPipeline.copyUntil(pe)
	if err != nil {
		return nil, nil, err
	}

	// Now link the copied parent to the copied element
	output.SetInput(NewDatapointPeekIterator(o))

	// The script is now linked together
	return input, output, nil
}

// NewPipelineElement creates a PipelineElement given an array of argument scripts and a TransformInstance
func NewPipelineElement(args []*Script, t TransformInstance) (*PipelineElement, error) {
	// Ensure the args are one to one
	for i := range args {
		if !args[i].OneToOne {
			return nil, errors.New("All transform arguments must be OneToOne")
		}
	}

	return &PipelineElement{
		iter: TransformIterator{
			args:      args,
			argpeeker: nil,
			iterator:  nil,
			peeklist:  list.New(),
			Err:       nil,
		},
		transform: t,
	}, nil
}
