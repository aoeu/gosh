package main

// go get github.com/aoeu/gosh/cmd/path && echo "function goto { cd $(path $*); }" >> ~/.profile && source ~/.profile

import (
	"fmt"
	"os"
	"strings"
)

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
