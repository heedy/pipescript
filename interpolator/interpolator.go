package interpolator

import (
	"fmt"

	"github.com/connectordb/pipescript"
)

// InterpolatorInstance is the core interpolation interface. Given an ordered stream of timestamps inserted into its
// next method, it returns the interpolated data. Usually it is assumed that a DatapointIterator is inserted
// into the interpolator during creation.
// The interpolator works by "interpolating" the stream of datapoints to fit the given timestamps. For each (increasing)
// timestamp that is given to the interpolator, the interpolator figures out what to return based upon the underlying data stream.
// It is assumed that the interpolator's decision is fairly greedy, and can work on arbitrarily large data without running out of memory.
type InterpolatorInstance interface {
	Next(timestamp float64) (*pipescript.Datapoint, error)
}

// InterpolatorRegsitry is a map of the registered interpolators. Only use this structure when displaying documentation, do not manually
// add interpolators. Instead, use the Interpolator.Register() method.
var InterpolatorRegistry = make(map[string]Interpolator)

// InterpolatorGenerator creates a new InterpolatorInstance.
type InterpolatorGenerator func(name string, dpi pipescript.DatapointIterator) (i InterpolatorInstance, err error)

// Interpolator is the struct which holds documentation and generator for an Interpolation method.
type Interpolator struct {
	Name        string
	Description string

	Generator InterpolatorGenerator `json:"-"` // The generator function of the interpolator
}

// Regsiter registers the interpolator with the system. See documentation for pipescript.Transform.Regsiter()
// as the two are implemented in the same way.
func (i Interpolator) Register() error {
	if i.Name == "" || i.Generator == nil {
		err := fmt.Errorf("Attempted to register invalid interpolator: '%s'", i.Name)
		return err
	}
	_, ok := InterpolatorRegistry[i.Name]
	if ok {
		err := fmt.Errorf("An Interpolator with the name '%s' already exists.", i.Name)
		return err
	}

	InterpolatorRegistry[i.Name] = i

	return nil
}

// Parse parses the interpolator given and returns everything initialized. If the given string cannot be
// parsed as an interpolator, it is assumed to be PipeScript, and a ScriptInterpolator is returned based
// upon the pipescript. If both these methods fail, returns an error.
func Parse(interpolator string, dpi pipescript.DatapointIterator) (InterpolatorInstance, error) {
	ireg, ok := InterpolatorRegistry[interpolator]
	if ok {
		// The given interpolator was found - return the InterpolatorInstance
		return ireg.Generator(ireg.Name, dpi)
	}

	// Not found. Try creating a PipeScript of the given string
	s, err := pipescript.Parse(interpolator)
	if err != nil {
		return nil, err
	}

	// Now set up the ScriptInterpolator based on the script
	return NewScriptInterpolator(s, dpi, nil)
}
