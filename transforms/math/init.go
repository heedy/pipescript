/*
Package math contains basic statistical transforms (such as mean)
*/
package math

import (
	"github.com/connectordb/pipescript"
	"github.com/connectordb/pipescript/transforms/core"
)

func init() {
	identity, err := core.IdentityTransform.Script([]*pipescript.Script{pipescript.ConstantScript(nil)})
	if err != nil {
		panic("Could not generate identity script: " + err.Error())
	}

	// Set up the default args for the transforms which use them
	Max.Args[0].Default = identity
	Min.Args[0].Default = identity
}

func Register() {
	Mean.Register()

	Max.Register()
	Min.Register()
}
