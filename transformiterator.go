package pipescript

import "container/list"

// PipelineElement is an element of the pipeline. It represents a statement between two pipes, and satisfies the DatapointIterator interface
type TransformIterator struct {
	args      []*Script
	argpeeker []*VirtualPeekIterator // The VirtualPeekIterator used for args
	iterator  DatapointPeekIterator  // The PeekIterator used to find the right datapoint
	peeklist  *list.List             // A cache that permits peeking forward in the sequence
	Err       error
}

// getNextEnvironment gets the TransformEnvironment that is to be passed to the TransformInstance
func (t *TransformIterator) getNextEnvironment() *TransformEnvironment {
	var err error

	// Set up the arguments. The ArgPeeker is assumed to be a VirtualPeekIterator based upon iterator
	// and is set up as the input to all of the args.
	args := make([]*Datapoint, len(t.args))
	for i := range t.args {
		t.argpeeker[i].SetBack(1)
		args[i], err = t.args[i].Next()
		if err != nil {
			t.Err = err
			return &TransformEnvironment{nil, nil, err}
		}
	}

	// All of the args are set up. Call Next() to get the Datapoint
	var dp *Datapoint
	dp, err = t.iterator.Next()

	// Return the formed TransformEnvironment
	return &TransformEnvironment{dp, args, err}
}

// Next gets the next TransformEnvironment
func (t *TransformIterator) Next() *TransformEnvironment {
	if t.peeklist.Len() > 0 {
		// There are datapoints in the cache
		te := t.peeklist.Remove(t.peeklist.Front()).(*TransformEnvironment)
		if te == nil {
			return &TransformEnvironment{Error: t.Err}
		}
		return te
	}

	// The peeklist does not contain the TransformEnvironment, so get a new one
	return t.getNextEnvironment()
}

// Peek looks at the datapoint/arg configuration n datapoints ahead of Next,
// but does not move the iterator forward (ie, Next is not called)
func (t *TransformIterator) Peek(forward int) (tenv *TransformEnvironment) {

	// Check if the peeklist has the element
	if forward < t.peeklist.Len() {
		// The data is on the peeklist! Now check which way will be faster for access:
		// forwards or backwards.
		if t.peeklist.Len()/2-1 >= forward {
			// Start from the front
			peekElement := t.peeklist.Front()
			for ; forward > 0; forward-- {
				peekElement = peekElement.Next()
			}
			return peekElement.Value.(*TransformEnvironment)
		}
		// Start from the back
		peekElement := t.peeklist.Back()

		for forward++; forward < t.peeklist.Len(); forward++ {
			peekElement = peekElement.Prev()
		}
		return peekElement.Value.(*TransformEnvironment)

	}

	//The element is not on the peeklist. Check if we are done iterating or had error
	if t.Err != nil {
		return &TransformEnvironment{nil, nil, t.Err}
	}

	if t.peeklist.Back() != nil && t.peeklist.Back().Value.(*TransformEnvironment).IsFinished() {
		return t.peeklist.Back().Value.(*TransformEnvironment)
	}

	// Extend the peeklist so that we get to the desired environment
	forward -= t.peeklist.Len()
	for ; forward >= 0; forward-- {
		tenv = t.getNextEnvironment()
		t.Err = tenv.Error
		t.peeklist.PushBack(tenv)
		if tenv.Error != nil || tenv.Datapoint == nil {
			return tenv
		}
	}

	return tenv
}
