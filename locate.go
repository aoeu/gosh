package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var usage string = `
Usage: '%s regexp', where "regexp" is a regular expression to match a file name.
Example: %s 'example.*.txt'
`

var filenameRegexp *regexp.Regexp
var paths chan string

func main() {
	if len(os.Args) != 2 {
		s := os.Args[0]
		fmt.Fprintf(os.Stderr, usage, s, s)
		os.Exit(1)
	}
	filenameRegexp = regexp.MustCompile(os.Args[1])
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	errors := make(chan error)
	paths = make(chan string)
	dirs, err := ioutil.ReadDir(wd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	var wg sync.WaitGroup
	for _, info := range dirs {
		switch {
		case info.IsDir():
			wg.Add(1)
			go func() {
				if err := filepath.Walk(wd, Mark); err != nil {
					errors <- err
				}
				wg.Done()
			}()
		case filenameRegexp.Match([]byte(info.Name())):
			paths <- fmt.Sprintf("%v/%v", wd, info.Name())
		}
	}
	go func() {
		wg.Wait()
		errors <- errNotFound
	}()
	select {
	case path := <-paths:
		fmt.Println(filepath.Dir(path))
		os.Exit(0)
	case err := <-errors:
		switch err {
		case errNotFound:
			fmt.Fprintf(os.Stderr, err.Error(), os.Args[1])
			os.Exit(0)
		default:
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		case nil:
		}

	}
}

var (
	errNotFound = errors.New("File not found with name matching '%s'\n")
)

func Mark(path string, info os.FileInfo, err error) error {
	switch {
	case info.IsDir():
		return nil
	case err != nil && err != os.ErrPermission:
		return err
	case filenameRegexp.Match([]byte(info.Name())):
		paths <- path
	}
	return nil
}
