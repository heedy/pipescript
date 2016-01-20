package interpolator

import (
	"fmt"
	"sync"

	"github.com/connectordb/pipescript"
)

// InterpolatorInstance is the core interpolation interface. Given an ordered stream of timestamps inserted into its
// next method, it returns the interpolated data. Usually it is assumed that a DatapointIterator is inserted
// into the interpolator during creation.
// The interpolator works by "interpolating" the stream of datapoints to fit the given timestamps. For each (increasing)
// timestamp that is given to the interpolator, the interpolator figures out what to return based upon the underlying data stream.
// It is assumed that the interpolator's decision is fairly greedy, and can work on arbitrarily large data without running out of memory.
type InterpolatorInstance interface {
	Interpolate(timestamp float64) (*pipescript.Datapoint, error)
}

// InterpolatorRegsitry is a map of the registered interpolators. Only use this structure when displaying documentation, do not manually
// add interpolators. Instead, use the Interpolator.Register() method.
var (
	InterpolatorRegistry = make(map[string]Interpolator)

	RegistryLock = &sync.RWMutex{}
)

// InterpolatorGenerator creates a new InterpolatorInstance.
type InterpolatorGenerator func(name string, dpi pipescript.DatapointIterator) (i InterpolatorInstance, err error)

// Interpolator is the struct which holds documentation and generator for an Interpolation method.
type Interpolator struct {
	Name         string `json:"name"`        // The name of the interpolator
	Description  string `json:"description"` // The docstring of the interpolator
	InputSchema  string `json:"ischema"`     // The schema of datapoints that are expected (optional)
	OutputSchema string `json:"oschema"`     // The schema of datapoints that are output (optional)

	Generator InterpolatorGenerator `json:"-"` // The generator function of the interpolator
}

// Unregister removes the given named interpolator from the registry
func Unregister(name string) {
	RegistryLock.Lock()
	delete(InterpolatorRegistry, name)
	RegistryLock.Unlock()
}

// Regsiter registers the interpolator with the system. See documentation for pipescript.Transform.Regsiter()
// as the two are implemented in the same way.
func (i Interpolator) Register() error {
	if i.Name == "" || i.Generator == nil {
		err := fmt.Errorf("Attempted to register invalid interpolator: '%s'", i.Name)
		return err
	}

	RegistryLock.Lock()
	defer RegistryLock.Unlock()

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
	RegistryLock.RLock()
	ireg, ok := InterpolatorRegistry[interpolator]
	RegistryLock.RUnlock()
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

// InterpolationIterator allows turning an InterpolatorInstance into a DatapointIterator
// by interpolating over a TimeRange object
type InterpolationIterator struct {
	ipltr InterpolatorInstance
	tr    TimeRange
}

// Next performs the interpolation based upon the given TimeRange and returns
// the next result in the sequence
func (ii *InterpolationIterator) Next() (*pipescript.Datapoint, error) {
	ts, err := ii.tr.Timestamp()
	if err != nil {
		if err == ErrEOF {
			return nil, nil
		}
		return nil, err
	}
	return ii.ipltr.Interpolate(ts)
}

// NewInterpolationIterator returns a DatapointIterator based upon the given Interpolator
// and the TimeRange over which to interpolate
func NewInterpolationIterator(ipltr InterpolatorInstance, tr TimeRange) *InterpolationIterator {
	return &InterpolationIterator{ipltr, tr}
}
