package gosh

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)


// UsageFunc takes a text/template string,
// and uses it to create a flag.Usage function
// that outputs recommended usage and default
// flags (as determined by the flag package).
func UsageFunc(usageTemplate string) func() {
	return func() {
		var t *template.Template
		var err error
		if t, err = template.New("usage").Parse(usageTemplate); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := t.Execute(os.Stdout, os.Args[0]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// TODO(aoeu): Print default flags before printing usage examples.
		flag.PrintDefaults()
		os.Exit(2)
	}
}
