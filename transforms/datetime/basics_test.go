package datetime

import (
	"testing"
	"time"

	"github.com/heedy/pipescript"
)

func TestDay(t *testing.T) {
	Day.Register()

	pipescript.TestCase{
		Pipescript: "day('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "day('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: 1454907600, Data: 1},
			{Timestamp: 1454821200, Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1454907600, Data: int64(16839)},
			{Timestamp: 1454821200, Data: int64(16838)},
		},
	}.Run(t)
}

func TestDayhour(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "dayhour('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "dayhour('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: time.Now().UTC().Hour()},
		},
	}.Run(t)
}

func TestHour(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "hour('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(1454907600), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1454907600, Data: int64(404141)}, // Takes time zone into account
		},
	}.Run(t)
}

func TestMonth(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "month('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "month('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(1454907600), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1454907600, Data: int64(554)},
		},
	}.Run(t)
}

func TestMonthday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "monthday('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "monthday",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: time.Now().Day()},
		},
	}.Run(t)
}

func TestWeek(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "week('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "week('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(1454907600), Data: 1},
			{Timestamp: 1454821200, Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1454907600, Data: int64(2405)},
			{Timestamp: 1454821200, Data: int64(2404)},
		},
	}.Run(t)
}

func TestWeekday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "weekday('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "weekday",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: time.Now().Weekday().String()},
		},
	}.Run(t)
}

func TestYear(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "year('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "year('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: 1454821200, Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: 1454821200, Data: int(2016)},
		},
	}.Run(t)
}
func TestYearday(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "yearday('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "yearday",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: time.Now().YearDay()},
		},
	}.Run(t)
}

func TestYearMonth(t *testing.T) {
	Register()

	pipescript.TestCase{
		Pipescript: "yearmonth('lol')",
		Parsed:     "error",
	}.Run(t)
	pipescript.TestCase{
		Pipescript: "yearmonth('UTC')",
		Input: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: 1},
		},
		Output: []pipescript.Datapoint{
			{Timestamp: float64(time.Now().Unix()), Data: time.Now().Month().String()},
		},
	}.Run(t)
}
