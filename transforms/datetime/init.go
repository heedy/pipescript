/*Package datetime offers transforms useful for processing timestamps*/
package datetime

import (
	"errors"
	"time"

	"github.com/heedy/pipescript"
)

type TimeBasicFunc func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error)

type timeBasic struct {
	ConstArgs []interface{}
	Args      []*pipescript.Datapoint
	PipeArgs  []*pipescript.Pipe
	f         TimeBasicFunc
	tz        *time.Location
	st        time.Time
}

func (s *timeBasic) Next(e *pipescript.TransformEnv, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
	dp, arr, err := e.Next(s.Args)
	if err != nil || dp == nil {
		return nil, err
	}
	out.Timestamp = dp.Timestamp
	out.Duration = dp.Duration

	return s.f(dp, arr, s.ConstArgs, s.PipeArgs, s.tz, s.st, out)
}

func (s *timeBasic) OneToOne() bool {
	return true
}

func NewTimeBasic(tzarg int, f TimeBasicFunc) pipescript.TransformConstructor {
	return func(transform *pipescript.Transform, consts []interface{}, pipes []*pipescript.Pipe) (pipescript.TransformIterator, error) {
		tz, err := getTimezone(consts[tzarg])
		if err != nil {
			return nil, err
		}

		return &timeBasic{
			ConstArgs: consts,
			PipeArgs:  pipes,
			f:         f,
			tz:        tz,
			st:        time.Date(1970, time.January, 1, 0, 0, 0, 0, tz),
			Args:      make([]*pipescript.Datapoint, len(transform.Args)),
		}, nil
	}
}

var timezoneArg = pipescript.TransformArg{
	Description: "The time zone to use for determining timestamps, in IANA timezone database format (ex: 'America/New_York'). 'Local' uses the server time zone. 'UTC' uses UTC.",
	Optional:    true,
	Default:     pipescript.MustPipe(pipescript.NewConstTransform("Local"), nil),
	Type:        pipescript.ConstArgType,
	Schema: map[string]interface{}{
		"type": "string",
	},
}

func getTimezone(arg interface{}) (*time.Location, error) {
	s, ok := arg.(string)
	if !ok {
		return nil, errors.New("Timezone must be a string")
	}
	return time.LoadLocation(s)
}

func Register() {

	Tshift.Register()

	Hour.Register()
	Day.Register()
	Week.Register()
	Month.Register()
	Year.Register()

	Dayhour.Register()

	Weekday.Register()
	Monthday.Register()
	Yearday.Register()

	Yearmonth.Register()

}
