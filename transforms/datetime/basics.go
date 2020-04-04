package datetime

import (
	"time"

	"github.com/heedy/pipescript"
)

var Hour = &pipescript.Transform{
	Name:        "hour",
	Description: "Returns the number of hours since Jan 1 1970 in the given time zone.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = int64(tval.Sub(st) / time.Hour)
		return out, nil
	}),
}

var Day = &pipescript.Transform{
	Name:        "day",
	Description: "Returns the number of days since Jan 1 1970 in the given time zone.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = int64(tval.Sub(st) / (24 * time.Hour))
		return out, nil
	}),
}

var Dayhour = &pipescript.Transform{
	Name:        "dayhour",
	Description: "Returns the hour in a day during which the datapoint happened.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.Hour()
		return out, nil
	}),
}

var Week = &pipescript.Transform{
	Name:        "week",
	Description: "Returns the number of weeks since Jan 5 1970 (First Monday after unix time) in the given time zone.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		// The week actually starts later
		st = time.Date(1970, time.January, 5, 0, 0, 0, 0, tz)
		out.Data = int64(tval.Sub(st) / (7 * 24 * time.Hour))
		return out, nil
	}),
}

var Weekday = &pipescript.Transform{
	Name:        "weekday",
	Description: "Returns the weekday during which the datapoint happened ('Monday','Tuesday'...)",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.Weekday().String()
		return out, nil
	}),
}

var Month = &pipescript.Transform{
	Name:        "month",
	Description: "Returns the number of months since Jan 1 1970 in the given time zone.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = int64(12*(tval.Year()-1970) + int(tval.Month()))
		return out, nil
	}),
}

var Monthday = &pipescript.Transform{
	Name:        "monthday",
	Description: "Returns the number of day in the datapoint's month (int)",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.Day()
		return out, nil
	}),
}

var Year = &pipescript.Transform{
	Name:        "year",
	Description: "Returns the year of the datapoint in the given time zone.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.Year()
		return out, nil
	}),
}

var Yearmonth = &pipescript.Transform{
	Name:        "yearmonth",
	Description: "Returns the month name during which the datapoint happened ('January','February'...)",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.Month().String()
		return out, nil
	}),
}

var Yearday = &pipescript.Transform{
	Name:        "yearday",
	Description: "Returns the day of the year during which the datapoint happened [1,365] non-leap, and [1,366] for leap years.",
	Args: []pipescript.TransformArg{
		timezoneArg,
	},
	Constructor: NewTimeBasic(0, func(dp *pipescript.Datapoint, args []*pipescript.Datapoint, consts []interface{}, pipes []*pipescript.Pipe, tz *time.Location, st time.Time, out *pipescript.Datapoint) (*pipescript.Datapoint, error) {
		tval := dp.Time().In(tz)
		out.Data = tval.YearDay()
		return out, nil
	}),
}
