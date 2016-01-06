/*
Package pipescript contains the PipeScript streaming time series query engine and standard transform
functions. It is the core data processing method used in ConnectorDB.
*/
package pipescript

import "errors"

const (
	// Version is the PipeScript version
	Version = "0.1"
)

// Script is the main type used in PipeScript
type Script struct {
	input      *PipelineElement
	output     *PipelineElement
	IsOneToOne bool
	Constant   bool
}

// Append takes a Script and appends another script to the end of its command chain. That is,
// if s = (a | b | c) and s2 = (d | e | f) then s.Append(s2) will make s = (a | b | c | d | e | f)
// WARNING: This only works for uninitialized scripts (ie, scripts that have not had Datapoints pass through yet)
// behavior for initialized scripts is undefined.
func (s *Script) Append(s2 *Script) error {

	if !s2.IsOneToOne {
		s.IsOneToOne = false
	}
	// Link the output of s with the input of s2
	dpi := NewDatapointPeekIterator(s2.output)
	s2.input.SetInput(dpi)

	// The total output is not s2's output
	s.output = s2.output

	if s2.Constant {
		if s.Constant || s.IsOneToOne {
			// If both are constants, or if s is one to one, and s2 is constant,
			// it means that we can replace the full query by a constant
			c, err := s2.GetConstant()
			if err != nil {
				return err
			}
			s.replaceWithConstant(c)
			return nil
		}
	} else {
		s.Constant = false
	}
	return nil
}

// SetInput sets the input DatapointIterator of the stream. It will automatically recognize PeekIterators.
func (s *Script) SetInput(d DatapointIterator) {
	pi, ok := d.(DatapointPeekIterator)
	if !ok {
		pi = NewDatapointPeekIterator(d)
	}

	// Set the PipelineElement input
	s.input.SetInput(pi)
}

// Next allows Script to conform to the DatapointIterator interface.
func (s *Script) Next() (*Datapoint, error) {
	return s.output.Next()
}

// constantIterator is used as an iterator with an always-constant value
type constantIterator struct {
	C interface{}
}

func (c constantIterator) Next() (*Datapoint, error) {
	return &Datapoint{0, c.C}, nil
}
func (c constantIterator) Peek(p int) (*Datapoint, error) {
	return c.Next()
}

// GetConstant passes through a dummy Datapoint through the transform to get the constant value it represents.
func (s *Script) GetConstant() (*Datapoint, error) {
	if !s.Constant {
		return nil, errors.New("GetConstant was called on a non-constant Script")
	}
	oldIterator := s.input.iter.iterator

	s.SetInput(constantIterator{0})
	dp, err := s.Next()
	if err != nil {
		return nil, err
	}

	// Replace the script with a newly created ConstantScript
	s.replaceWithConstant(dp.Data)

	// If the iterator was already initialized, switch it to the new constant
	if oldIterator != nil {
		s.SetInput(oldIterator)
	}

	return dp, nil
}

// replaceWithConstant replaces the Script with a constant script
func (s *Script) replaceWithConstant(c interface{}) {
	oldIterator := s.input.iter.iterator

	s2 := ConstantScript(c)
	s.input = s2.input
	s.output = s2.output
	s.IsOneToOne = s2.IsOneToOne
	s.Constant = s2.Constant

	if oldIterator != nil {
		s.SetInput(oldIterator)
	}
}
