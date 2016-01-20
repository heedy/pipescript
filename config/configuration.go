/*Package config provides the structures necessary to set up a PipeScript configuration.

PipeScript currently has a single global configuration that can be modified.
*/

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/interpolator"
	"github.com/connectordb/pipescript/transforms/core"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

var (
	defaultSplitMax = core.SplitMax
	defaultNextMax  = core.NextMax

	configHeader = fmt.Sprintf("// PipeScript v%s Configuration\n", pipescript.Version)

	// Maps of the disabled interpolators and transforms
	disabledInterpolators = make([]interpolator.Interpolator, 0)
	disabledTransforms    = make([]pipescript.Transform, 0)
)

// Configuration contains all of the things that can be modified in PipeScript.
type Configuration struct {
	// The maximum number of elements in a reduce
	ReduceSplitMax int `json:"reduce_split_max"`
	// The maximum depth of a next call
	NextMax int64 `json:"next_max"`

	// An array of transform names to disable
	DisableTransforms []string `json:"disable_transforms"`

	DisableInterpolators []string `json:"disable_interpolators"`
}

// Default uses the built-in values to set up a default "good enough" configuration
func Default() *Configuration {
	return &Configuration{
		ReduceSplitMax:       defaultSplitMax,
		NextMax:              defaultNextMax,
		DisableTransforms:    []string{},
		DisableInterpolators: []string{},
	}
}

// Set writes the PipeScript settings to the global PipeScript instance
func (c *Configuration) Set() error {
	core.SplitMax = c.ReduceSplitMax
	core.NextMax = c.NextMax

	// First, reenable all interpolators and transforms
	for k := range disabledInterpolators {
		disabledInterpolators[k].Register()
	}
	for k := range disabledTransforms {
		disabledTransforms[k].Register()
	}

	// Now, we set up the disabled and enabled transforms/interpolators
	pipescript.RegistryLock.Lock()
	defer pipescript.RegistryLock.Unlock()

	interpolator.RegistryLock.Lock()
	defer interpolator.RegistryLock.Unlock()

	// Reset the disabled lists
	disabledInterpolators = make([]interpolator.Interpolator, 0, len(c.DisableInterpolators))
	disabledTransforms = make([]pipescript.Transform, 0, len(c.DisableTransforms))

	// Put the disabled transforms/interpolators on the disabled lists
	for k := range c.DisableInterpolators {
		i, ok := interpolator.InterpolatorRegistry[c.DisableInterpolators[k]]
		if ok {
			disabledInterpolators = append(disabledInterpolators, i)
			delete(interpolator.InterpolatorRegistry, c.DisableInterpolators[k])
		}
	}
	for k := range c.DisableTransforms {
		i, ok := pipescript.TransformRegistry[c.DisableTransforms[k]]
		if ok {
			disabledTransforms = append(disabledTransforms, i)
			delete(pipescript.TransformRegistry, c.DisableTransforms[k])
		}
	}

	return nil
}

// Validate makes sure that the configuration is OK
func (c *Configuration) Validate() error {
	if c.NextMax <= 1 {
		return errors.New("The maximum peek depth for PipeScript must be > 1")
	}
	if c.ReduceSplitMax <= 1 {
		return errors.New("The maximum reduce split size in PipeScript must be >= 1")
	}
	return nil
}

// Load reads the configuration from a file
func Load(filename string) (*Configuration, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// To allow comments in the json, we minify the file with js minifer before parsing
	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)
	file, err = m.Bytes("text/javascript", file)
	if err != nil {
		return nil, err
	}

	d := Default()
	err = json.Unmarshal(file, d)
	if err != nil {
		return nil, err
	}

	return d, d.Validate()
}

// Save saves the configuration
func (c *Configuration) Save(filename string) error {
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(configHeader))
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	return err
}
