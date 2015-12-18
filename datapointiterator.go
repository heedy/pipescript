package pipescript

import "container/list"

// DatapointIterator is assumed to return datapoints ordered by increasing timestamp.
// At any point in time, there can be a read error, which will cause the iterator to fail.
// In that case, it returns an error value, at which point the iterator is assumed to be invalid.
// If there is no error, the iterator returns Datapoints until the stream is finished, at which point
// it returns `nil` (without an error), signalling EOF.
type DatapointIterator interface {
	Next() (*Datapoint, error)
}

// DatapointArrayIterator is a DatapointIterator which iterates through the given array one datapoint
// at a time.
type DatapointArrayIterator struct {
	Datapoints []Datapoint

	i int // i is the current locatino in the array
}

func NewDatapointArrayIterator(dp []Datapoint) *DatapointArrayIterator {
	return &DatapointArrayIterator{dp, 0}
}

// Next returns the next datapoint in the array
func (d *DatapointArrayIterator) Next() (*Datapoint, error) {
	if d.i < len(d.Datapoints) {
		dp := d.Datapoints[d.i]
		d.i++
		return &dp, nil
	}
	return nil, nil
}

// DatapointPeekIterator permits peeking ahead in the sequence of Datapoints given in an iterator
type DatapointPeekIterator struct {
	Iterator DatapointIterator // The iterator used to find the correct datapoint
	PeekList *list.List        // A cache to permit peeking forward in the sequence
	Err      error             // If the iterator returns an error, cache it here
}

// NewDatapointPeekIterator creates a new lookahead cache
func NewDatapointPeekIterator(iter DatapointIterator) *DatapointPeekIterator {
	return &DatapointPeekIterator{iter, list.New(), nil}
}

// Next allows use of DatapointCache as a DatapointIterator
func (c *DatapointPeekIterator) Next() (*Datapoint, error) {
	if c.PeekList.Len() > 0 {
		// There are datapoints in the cache
		dp := c.PeekList.Remove(c.PeekList.Front()).(*Datapoint)
		if dp == nil {
			return dp, c.Err
		}
		return dp, nil
	}

	return c.Iterator.Next()
}

// Peek allows to look forward into the data sequence without losing its place for Next. Peek(0) is
// The value that would be returned from Next. Must be >=0.
func (c *DatapointPeekIterator) Peek(forward int) (dp *Datapoint, err error) {

	// Check if the peeklist has the element
	if forward < c.PeekList.Len() {
		// The data is on the peeklist! Now check which way will be faster for access:
		// forwards or backwards.
		if c.PeekList.Len()/2 >= forward {
			// Start from the front
			peekElement := c.PeekList.Front()
			for ; forward > 0; forward-- {
				peekElement = peekElement.Next()
			}
			return peekElement.Value.(*Datapoint), nil
		}
		// Start from the back
		peekElement := c.PeekList.Back()
		for ; forward < c.PeekList.Len(); forward++ {
			peekElement = peekElement.Prev()
		}
		return peekElement.Value.(*Datapoint), nil

	}

	//The element is not on the PeekList. Check if we are done iterating or had error
	if c.Err != nil || (c.PeekList.Back() != nil && c.PeekList.Back().Value.(*Datapoint) == nil) {
		return nil, c.Err
	}

	// Extend the peeklist so that we get to the desired datapoint
	forward -= c.PeekList.Len()
	for ; forward >= 0; forward-- {
		dp, err = c.Iterator.Next()
		c.Err = err
		c.PeekList.PushBack(dp)
		if err != nil || dp == nil {
			return nil, err
		}
	}

	return dp, err
}
