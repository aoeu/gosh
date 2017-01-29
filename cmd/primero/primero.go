package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [ [ < filepath] | [-of <filepath> ] ]

"{{.}}" prints the first line of text from standard input or a specified file
to standard output that is not empty when trimmed of leading and trailing 
whitespace characters, and exits with a non-error status.

Otherwise, "{{.}}"  exits with an error status and prints nothing to
standard output or standard error.

The emitted line of text printed to standard output has all leading and
trailing whitespace characters removed, followed by a newline.

Examples:
	$ find $GOPATH/src -name '*.go' | {{.}}
	> /home/username/go/src/golang.org/x/tools/benchmark/parse/parse_test.go
	$ echo $?
	> 0

	$ {{.}} < /tmp/empty_file
	$ echo $?
	> 1

	$ touch /tmp/arbitrary_file
	$ sam -d /tmp/arbitrary_file
	>  -. /tmp/arbitrary_file
	> a
	> 
	>   
	>         golang
	> awk
	>         sed
	>    grep
	> 
	> .
	> w
	> /tmp/arbitrary_file: #30
	> q
	$ {{.}} -of /tmp/arbitrary_file 
	> golang

	$ cat << EOF | {{.}}

			first 
                second 
                third
                fourth
                EOF
	> first

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		filepath string
	}{}
	flag.StringVar(&args.filepath, "of", "", "A filepath to print the first line of.")
	flag.Parse()
	f := os.Stdin
	var err error
	if args.filepath != "" {
		f, err = os.Open(args.filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open file at the path '%v' due to error '%v', exitng.\n", args.filepath, err)
		}
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		t := strings.Trim(input.Text(), " \t\n")
		switch {
		case t == "":
			continue
		default:
			fmt.Fprintln(os.Stdout, t)
			os.Exit(0)
		}

	}
	os.Exit(1)
}
