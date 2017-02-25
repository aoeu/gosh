package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [file]...

'{{.}}' reads the contents of any provided files until "end-of-file" (EOF) 
and sequentially copies the file contents to standard output.

Filepaths may also be provided by standard input.

Examples:

	{{.}} /dev/urandom
	
	{{.}} file1 file2 file3

	echo 'test Darwin = $(uname) && man cat | grep -B3 -A1 "Rob Pike"' > /tmp/file && eval $({{.}} /tmp/file)

	find . -name '*.txt' | {{.}}

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	var files []string
	switch len(os.Args) {
	case 1:
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if s := strings.TrimSpace(input.Text()); s != "" {
				files = append(files, s)
			}
		}
	default:
		files = os.Args[1:]
	}

	nonExistent := make([]string, 0)
	dirs := make([]string, 0)

	for _, f := range files {
		info, err := os.Stat(f)
		switch {
		case err != nil && os.IsNotExist(err):
			nonExistent = append(nonExistent, f)
		case err != nil:
			fmt.Fprintf(os.Stderr, "Error encountered when checking file paths are valid: %v", err)
		case info.IsDir():
			dirs = append(dirs, f)
		}
	}
	switch {
	case len(nonExistent) == 0 && len(dirs) == 0:
	case len(nonExistent) > 0:
		fmt.Fprintf(os.Stderr, "Filepaths were provided that do not exist, exiting: %v\n", nonExistent)
		fallthrough
	case len(dirs) > 0:
		fmt.Fprintf(os.Stderr, "Filepaths were provided that are directories, exiting: %v\n", dirs)
		fallthrough
	default:
		flag.Usage()
	}
	for _, f := range files {
		r, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error occurred when opening file %v : %v\n", f, err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, r)
	}
}