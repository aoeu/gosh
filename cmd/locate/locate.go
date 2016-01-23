package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var usage string = `
Usage: '%s regexp', where "regexp" is a regular expression to match a file name.
Example: %s 'example.*.txt'
`

var (
	filenameRegexp *regexp.Regexp
	paths          chan string
	errs           chan error
	errNotFound    = errors.New("File not found with name matching '%s'\n")
	done           bool
)

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
	errs = make(chan error)
	paths = make(chan string)
	go func() {
		errs <- Walk(wd, Mark)
	}()
	for i := 0; i < 2; i++ {
		select {
		case path := <-paths:
			fmt.Println(filepath.Dir(path))
			os.Exit(0)
		case err := <-errs:
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
}

func Mark(path string, info os.FileInfo, err error) error {
	switch {
	case done || info.IsDir():
		return nil
	case err != nil && err != os.ErrPermission:
		return err
	case filenameRegexp.Match([]byte(info.Name())):
		done = true
		paths <- path
		return nil
	}
	return errNotFound
}

// Code below this line is an excerpt from the standard library package named "filepath."

func Walk(root string, walkFn filepath.WalkFunc) error {
	info, err := os.Lstat(root)
	if err != nil {
		return walkFn(root, nil, err)
	}
	return walk(root, info, walkFn)
}

var SkipDir = errors.New("skip this directory")

func walk(path string, info os.FileInfo, walkFn filepath.WalkFunc) error {
	err := walkFn(path, info, nil)
	if err != nil && err != errNotFound {
		if info.IsDir() && err == SkipDir {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}

	names, err := readDirNames(path)
	if err != nil {
		return walkFn(path, info, err)
	}

	for _, name := range names {
		filename := filepath.Join(path, name)
		fileInfo, err := os.Lstat(filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil && err != SkipDir {
				return err
			}
		} else {
			err = walk(filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || err != SkipDir {
					return err
				}
			}
		}
	}
	return nil
}

func readDirNames(dirname string) ([]string, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	// Unlike filepath's readDirNames function, do not sort directory names here.
	if err != nil {
		return nil, err
	}
	return names, nil
}
