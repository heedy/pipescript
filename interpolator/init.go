package interpolator

import (
	"github.com/connectordb/pipescript"

	// transform core is required for datasets to work. In particular, we need "if last"
	_ "github.com/connectordb/pipescript/transforms/core"
)

func init() {
	// In order to initialize
	var err error
	iflast, err = pipescript.Parse("if last")
	if err != nil {
		panic("Could not generate script for 'if last' for ScriptInterpolator.")
	}
}
