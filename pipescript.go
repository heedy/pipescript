/*
Package pipescript contains the PipeScript streaming time series query engine and standard transform
functions. It is the core data processing method used in ConnectorDB.
*/
package pipescript

import (
	"errors"
	"fmt"
)

const (
	// Version is the PipeScript version
	Version = "0.2"
)

// Script is the main type used in PipeScript
type Script struct {
	input     *PipelineElement
	output    *PipelineElement
	OneToOne  bool // Whether the script is one to one (for each input, gives an output)
	Constant  bool // Whether the script is constant (always returns the same answer)
	Stateless bool // Whether the script is stateless (given input, always returns the same value)
	Peek      bool // Whether the script peeks at future values.
}

func (s *Script) String() string {
	return fmt.Sprintf("<SCRIPT onetoone=%v constant=%v stateless=%v peek=%v >", s.OneToOne, s.Constant, s.Stateless, s.Peek)
}

// MarshalJSON allows us to implement the marshaller interface when using scripts as default args. It returns null,
// so that it seems empty to json
func (s *Script) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

// Append takes a Script and appends another script to the end of its command chain. That is,
// if
//	s = (a | b | c)
// and
//	s2 = (d | e | f)
// then
//	s.Append(s2) => s = (a | b | c | d | e | f)
//
// WARNING: This only works for uninitialized scripts (ie, scripts that have not had Datapoints pass through yet)
// behavior for initialized scripts is undefined.
func (s *Script) Append(s2 *Script) error {
	if s2 == s {
		return errors.New("Can't append self to self")
	}

	if !s2.OneToOne {
		s.OneToOne = false
	}

	if !s2.Stateless {
		s.Stateless = false
	}

	if s2.Peek {
		s.Peek = true
	}

	// Link the output of s with the input of s2
	dpi := NewDatapointPeekIterator(s.output)
	s2.input.SetInput(dpi)

	// The total output is now s2's output
	s.output = s2.output

	if s2.Constant && s.OneToOne || s.Constant && s2.Stateless {
		// if s is one to one, and s2 is constant,
		// it means that we can replace the full query by a constant

		// Similarly, if the input to a stateless transform is constant,
		// we can replace the full thing with a constant
		s.Constant = true
		c, err := s2.GetConstant()
		if err != nil {
			return err
		}
		s.replaceWithConstant(c.Data)
		return nil
	}

	s.Constant = false
	return nil
}

// SetInput sets the input DatapointIterator of the stream. It will automatically recognize PeekIterators.
// use this command to link a Script to data.
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
	s.OneToOne = s2.OneToOne
	s.Constant = s2.Constant
	s.Stateless = s2.Stateless
	s.Peek = s2.Peek

	if oldIterator != nil {
		s.SetInput(oldIterator)
	}
}

// Copy performs a script copy.
//
// Warning: Copy may only be used during script initialization. If it is used after
// datapoints started flowing through the script, copy will fail. This is because peeked
// data is not copied at this time.
//
func (s *Script) Copy() (*Script, error) {
	i, o, err := s.output.copyUntil(s.input)
	return &Script{i, o, s.OneToOne, s.Constant, s.Stateless, s.Peek}, err
}

// Parse parses the given transform, and returns the corresponding script object
func Parse(script string) (*Script, error) {
	lexer := parserLex{input: script}

	if parserParse(&lexer) != 0 {
		if lexer.errorString != "" {
			return nil, fmt.Errorf("Script '%s': %s", script, lexer.errorString)
		}
		return nil, fmt.Errorf("Script '%s': Unknown error", script)
	}

	return lexer.output, nil
}
