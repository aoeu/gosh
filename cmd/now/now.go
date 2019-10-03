package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}}

'{{.}}' outputs the current date, time, or datetime, optionally
with added or subtracted units of duration.

Examples:

	{{.}}

	{{.}} -date

	{{.}} -time

	{{.}} -date -time

	{{.}} -date -minus -days 30

	{{.}} -time -plus -minutes 9

	{{.}} -date -time -plus hours 23

Flags:

`

const (
	dateDesc      = `output the date, formatted as "YYYY-MM-DD"`
	timeDesc      = `output the time, formatted as "HH:MM:SS"`
	datetimeDesc  = `output the datetime, formatted as "YYYY-MM-DD_HH:MM:SS", equivalent to using both '-date' and '-time' flags or excluding all of the "-date", "-time", and "-datetime" flags).`
	shouldAddDesc = "add any durations specified with flags to the current time"
	shouldSubDesc = "subtract any durations specified with flags to the current time"

	daysDesc    = "the number of days to add or substract from the current time"
	hoursDesc   = "the number of hours to add or substract from the current time"
	minutesDesc = "the number of minutes to add or substract from the current time"
)

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		inclDate     bool
		inclTime     bool
		inclDatetime bool
		shouldAdd    bool
		shouldSub    bool
		days         uint
		hours        uint
		minutes      uint
	}{}
	flag.BoolVar(&args.inclDate, "date", false, dateDesc)
	flag.BoolVar(&args.inclTime, "time", false, timeDesc)
	flag.BoolVar(&args.inclDatetime, "datetime", false, datetimeDesc)
	flag.BoolVar(&args.shouldAdd, "plus", false, "")
	flag.BoolVar(&args.shouldSub, "minus", false, "")

	flag.UintVar(&args.days, "days", 0, daysDesc)
	flag.UintVar(&args.hours, "hours", 0, hoursDesc)
	flag.UintVar(&args.minutes, "minutes", 0, minutesDesc)
	flag.Parse()

	switch {
	case args.inclDatetime && (args.inclDate || args.inclTime):
		s := "use the ''date' and 'time' flags together or separately, "
		s += "xor 'datetime' on its own, "
		s += "xor none of the 'date', 'time', and 'datetime' flags whatsoever."
		printErr(s)
		flag.Usage()
	case args.shouldAdd && args.shouldSub:
		printErr("both the 'plus' and 'minus' flags may not be used at once")
		flag.Usage()
	case args.shouldAdd, args.shouldSub:
		if args.minutes < 1 && args.hours < 1 && args.days < 1 {
			printErr("a unit of time to add or subtract by should be specified")
			flag.Usage()
		}
	}

	var (
		days    time.Duration
		hours   time.Duration
		minutes time.Duration
		err     error
		errFmt  = "could not parse %v %v as duration: %v"
	)
	if days, err = parseDuration(args.days, 'd'); err != nil {
		printErr(errFmt, args.days, "days", err)
	}
	if hours, err = parseDuration(args.hours, 'h'); err != nil {
		printErr(errFmt, args.hours, "hours", err)
	}
	if minutes, err = parseDuration(args.minutes, 'm'); err != nil {
		printErr(errFmt, args.minutes, "minutes", err)
	}

	t := time.Now()
	switch {
	case args.shouldAdd:
		t = t.Add(days).Add(hours).Add(minutes)
	case args.shouldSub:
		t = t.Add(-days).Add(-hours).Add(-minutes)
	}

	timeFmt := "2006-01-02_15:04:05"
	if args.inclDate || args.inclDatetime {
		timeFmt = "2006-01-02"
	}
	if (args.inclDate && args.inclTime) || args.inclDatetime {
		timeFmt += "_"
	}
	if args.inclTime || args.inclDatetime {
		timeFmt += "15:04:05"
	}
	fmt.Println(t.Format(timeFmt))
}

func printErr(format string, a ...interface{}) (n int, err error) {
	i := len(format) - 1
	if i > 0 && format[i] != '\n' {
		format += "\n"
	}
	return fmt.Fprintf(os.Stderr, format, a...)
}

func parseDuration(quantity uint, unit rune) (time.Duration, error) {
	if quantity == 0 {
		return 0, nil
	}
	switch unit {
	case 'd':
		unit = 'h'
		quantity *= 24
	case 'h', 'm':
	default:
		return 0, fmt.Errorf("unit %v must be 'd', 'h', or 'm'", unit)
	}
	s := fmt.Sprintf("%v%c", quantity, unit)
	return time.ParseDuration(s)
}
