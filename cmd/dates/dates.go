package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage {{.}} -from <DATE> -to <DATE> [-with <DELIMITER>]

'{{.}}' outputs a range of dates, separated with newlines by default,
or separated by a delimiter string provided as an optional argument,
with a newline following the final outputted date (instead of the delimiter).

Examples:

	{{.}} -from 2007-01-01 -to 2007-02-01

	{{.}} -from 2019-10-01 -to 2019-10-31 -with ', '

	{{.}} -from 01/01/1970 -to 01/01/1972 -like '01/02/2006'
	

`

const (
	toDesc            = "the date to output the date range until (inclusive)"
	fromDesc          = "the date to start outputting the date range from (inclusive)"
	withDesc          = "an optional string to delimit the dates in the range with"
	likeDesc          = "an optional date format to print the dates as, examplified with January 1, 2006."
	argErr            = "Both a date to start from and a date to stop at must be provided by the '-from' and '-to' arguments, respectively."
	defaultDateFormat = "2006-01-02"
)

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		from       string
		to         string
		delimiter  string
		dateFormat string
	}{}
	flag.StringVar(&args.from, "from", "", fromDesc)
	flag.StringVar(&args.to, "to", "", toDesc)
	flag.StringVar(&args.delimiter, "with", "\n", withDesc)
	flag.StringVar(&args.dateFormat, "like", defaultDateFormat, likeDesc)
	flag.Parse()

	switch "" {
	case args.from, args.to:
		printErr(argErr)
	}

	t, err := time.Parse(args.dateFormat, args.dateFormat)
	if err != nil {
		s := "internal error: could not initialize the date '%v': %v"
		printErr(s, defaultDateFormat, err)
	}
	if t.Format(defaultDateFormat) != defaultDateFormat {
		s := "the default date format argument '%v' must be expressed as January 1st, 2006 (i.e. 2006-01-02 or 01/02/2006, etc.)"
		printErr(s, args.dateFormat)
	}

	from, err := time.Parse(args.dateFormat, args.from)
	if err != nil {
		s := "could not parse date specified with '-from' (%v) in format from '-like' (%v): %v"
		printErr(s, args.from, args.dateFormat, err)
	}
	to, err := time.Parse(args.dateFormat, args.to)
	if err != nil {
		s := "could not parse date specified with '-to' (%v) in format from '-like' (%v): %v"
		printErr(s, args.to, args.dateFormat, err)
	}
	numDays := int(to.Sub(from).Hours() / 24)
	for i := 0; i <= numDays; i++ {
		t := from.AddDate(0, 0, i)
		if i == numDays {
			args.delimiter = "\n"
		}
		fmt.Printf("%v%v", t.Format(args.dateFormat), args.delimiter)
	}
}

func printErr(format string, a ...interface{}) (n int, err error) {
	i := len(format) - 1
	if i > 0 && format[i] != '\n' {
		format += "\n"
	}
	return fmt.Fprintf(os.Stderr, format, a...)
}
