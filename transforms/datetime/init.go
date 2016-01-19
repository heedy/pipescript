/*Package datetime offers transforms useful for processing timestamps*/
package datetime

import (
	"time"

	"github.com/connectordb/pipescript"
)

var timezoneArg = pipescript.TransformArg{
	Description: "The time zone to use for determining timestamps, in IANA timezone database format (ex: 'America/New_York'). 'Local' uses the server time zone. 'UTC' uses UTC.",
	Optional:    true,
	Default:     "Local",
	Constant:    true,
}

func getTimezone(arg *pipescript.Script) (*time.Location, error) {
	dp, err := arg.GetConstant()
	if err != nil {
		return nil, err
	}
	v, err := dp.DataString()
	if err != nil {
		return nil, err
	}
	return time.LoadLocation(v)
}

func Register() {
	weekday.Register()
}
