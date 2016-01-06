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

// NewPipelineElement creates a PipelineElement given an array of argument scripts and a TransformInstance
func NewPipelineElement(args []*Script, t TransformInstance) (*PipelineElement, error) {
	// Ensure the args are one to one
	for i := range args {
		if !args[i].IsOneToOne {
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
