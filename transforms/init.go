/*Package transforms imports all of the transforms that are available with PipeScript. The core PipeScript
only has an if statement and the identity operator, which are not nearly enough.

This package imports EVERYTHING
*/
package transforms

import (
	"github.com/heedy/pipescript/transforms/core"     // The core transforms
	"github.com/heedy/pipescript/transforms/datetime" // Manipulating timestamps
	"github.com/heedy/pipescript/transforms/misc"     // Miscellaneous transforms
	"github.com/heedy/pipescript/transforms/numeric"  // Statistical transforms
	"github.com/heedy/pipescript/transforms/strings"  // Text-based transforms
)

// Register ALL functions
func Register() {
	core.Register()
	numeric.Register()
	datetime.Register()
	strings.Register()
	misc.Register()
}
