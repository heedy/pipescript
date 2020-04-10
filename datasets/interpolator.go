package datasets

import (
	"fmt"
	"sync"

	"github.com/heedy/pipescript"
)

// InterpolatorRegsitry is a map of the registered interpolators. Only use this structure when displaying documentation, do not manually
// add interpolators. Instead, use the Interpolator.Register() method.
var (
	InterpolatorRegistry = map[string]*Interpolator{}

	RegistryLock = &sync.RWMutex{}
)

// InterpolatorConstructor creates a new Iterator,
// which will return a datapoint for each element of reference.
type InterpolatorConstructor func(name string, options map[string]interface{}, reference *pipescript.BufferIterator, stream pipescript.Iterator) (pipescript.Iterator, error)

// Interpolator is the struct which holds documentation and generator for an Interpolation method.
type Interpolator struct {
	Name          string                 `json:"name"`          // The name of the interpolator
	Description   string                 `json:"description"`   // The docstring of the interpolator
	Documentation string                 `json:"documentation"` // The full documentation for the interpolator in markdown
	Options       map[string]interface{} `json:"options"`       // The options available to the interpolator

	Constructor InterpolatorConstructor `json:"-"` // The generator function of the interpolator

	optionsSchema *JSONSchema `json:"-"` // The json schema of the options, initialized during registration
}

// Unregister removes the given named interpolator from the registry
func Unregister(name string) {
	RegistryLock.Lock()
	delete(InterpolatorRegistry, name)
	RegistryLock.Unlock()
}

// Regsiter registers the interpolator with the system. See documentation for pipescript.Transform.Regsiter()
// as the two are implemented in the same way.
func (i *Interpolator) Register() error {
	if i.Name == "" || i.Constructor == nil {
		err := fmt.Errorf("Attempted to register invalid interpolator: '%s'", i.Name)
		return err
	}
	s, err := NewSchema(i.Options)
	if err != nil {
		return err
	}
	i.optionsSchema = s

	RegistryLock.Lock()
	defer RegistryLock.Unlock()

	InterpolatorRegistry[i.Name] = i

	return nil
}

// GetInterpolator parses the interpolator given and returns everything initialized. If the given string cannot be
// parsed as an interpolator, it is assumed to be PipeScript, and a ScriptInterpolator is returned based
// upon the pipescript. If both these methods fail, returns an error.
func GetInterpolator(interpolator string, options map[string]interface{}, reference *pipescript.BufferIterator, stream pipescript.Iterator) (pipescript.Iterator, error) {
	if options == nil {
		options = make(map[string]interface{})
	}
	RegistryLock.RLock()
	ireg, ok := InterpolatorRegistry[interpolator]
	RegistryLock.RUnlock()
	if ok {
		// The given interpolator was found - make sure that the options conform to the json schema
		err := ireg.optionsSchema.ValidateWithDefaults(options)
		if err != nil {
			return nil, err
		}

		return ireg.Constructor(interpolator, options, reference, stream)
	}

	// Not found. Try creating a TransformInterpolator of the given string
	options["transform"] = interpolator
	err := TransformInterpolator.optionsSchema.ValidateWithDefaults(options)
	if err != nil {
		return nil, err
	}
	return TransformInterpolator.Constructor(interpolator, options, reference, stream)
}
