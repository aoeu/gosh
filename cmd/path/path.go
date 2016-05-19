package main

// go get github.com/aoeu/gosh/cmd/path && echo "function goto { cd $(path $*); }" >> ~/.profile && source ~/.profile

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [DIRECTORY]...

{{.}} takes a space separated list of directory names of a valid directory tree and prints the full path with separators specific to the host Operating System.

If the directory names do not create a complete path, a path under the user's home directory is attempted, then a path derived from the root directory, and finally an error is printed if none are found to be valid paths.

Example:

	{{.}} go src encoding json

Recipes:

	In a Bourne-compatible shell:

		go get github.com/aoeu/gosh/cmd/{{.}}
		echo 'function goto { cd $({{.}} $*); }' >> ~/.profile  && source ~/.profile
		goto go src net

	In fish:

		go get github.com/aoeu/gosh/cmd/{{.}}
		function goto
			cd ({{.}} $argv)
		end
		funcsave goto
		goto go src net

`

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func prependIfExists(prefix string, path *string) bool {
	if _, err := os.Stat(prefix + *path); err != nil {
		return false
	}
	*path = prefix + *path
	return true
}

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	dir := strings.Join(os.Args[1:], "/")
	switch {
	case exists(dir) ||
		prependIfExists(os.Getenv("HOME")+"/", &dir) ||
		prependIfExists("/", &dir):
		fmt.Println(dir)
	default:
		fmt.Fprintf(os.Stderr, "Unknown path: %v\n", dir)
		os.Exit(1)
	}
}
