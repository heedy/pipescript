package interpolator

import (
	"github.com/heedy/pipescript"

	// transform core is required for datasets to work. In particular, we need "if last"
	"github.com/heedy/pipescript/transforms/core"
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

}
