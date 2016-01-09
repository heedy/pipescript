/*Package transforms imports all of the transforms that are available with PipeScript. The core PipeScript
only has an if statement and the identity operator, which are not nearly enough.

This package imports EVERYTHING
*/
package transforms

import (
	_ "github.com/connectordb/pipescript/transforms/core"     // The core transforms
	_ "github.com/connectordb/pipescript/transforms/datetime" // Manipulating timestamps
)
