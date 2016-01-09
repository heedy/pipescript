package pipescript

// SingleDatapointIterator is made for use in transforms, when transforms want to keep internal scripts.
// it behaves as an iterator, but its value is set at each iteration. it does not support peek for obvious reasons.
// To see an example of its usage, look at the implementation of "split" in transforms/core
type SingleDatapointIterator struct {
	Datapoint *Datapoint
	Error     error
}

// Next returns the datapoint and error
func (s *SingleDatapointIterator) Next() (*Datapoint, error) {
	return s.Datapoint, s.Error
}

// Set sets the next datapoint to return. You can also just do it manually - your choice
func (s *SingleDatapointIterator) Set(d *Datapoint, e error) {
	s.Datapoint = d
	s.Error = e
}
