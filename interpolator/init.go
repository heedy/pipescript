package interpolator

import (
	"github.com/connectordb/pipescript"

	// transform core is required for datasets to work. In particular, we need "if last"
	"github.com/connectordb/pipescript/transforms/core"
)

func init() {
	// We need Last, if, and identity for good measure
	core.IdentityTransform.Register()
	core.If.Register()
	core.Last.Register()

	// In order to initialize
	var err error
	iflast, err = pipescript.Parse("if last")
	if err != nil {
		panic("Could not generate script for 'if last' for ScriptInterpolator.")
	}

	// Register built-in interpolators
	after.Register()
	before.Register()
	closest.Register()
}
