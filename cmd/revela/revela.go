package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: '{{.}} regexp'

'{{.}}' uses a regular expression to locate files with a matching name under the current working directory.

Example:

	{{.}} 'example.*\.txt'

`

var (
	filenameRegexp *regexp.Regexp
	paths          chan string
	errs           chan error
	errNotFound    = errors.New("File not found with name matching '%s'\n")
	done           chan bool
)

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
	}
	var err error
	filenameRegexp, err = regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	errs = make(chan error)
	paths = make(chan string)
	done = make(chan bool)
	go func() {
		if err := Walk(wd, Mark); err != nil {
			errs <- err
		} else {
			done <- true
		}
	}()
	for {
		select {
		case path := <-paths:
			if fp, err := filepath.Abs(path); err != nil {
				errs <- err
			} else {
				fmt.Println(fp)
			}
		case err := <-errs:
			switch err {
			case nil:
				continue
			case errNotFound:
				fmt.Fprintf(os.Stderr, err.Error(), os.Args[1])
				os.Exit(0)
			default:
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		case <-done:
			os.Exit(0)
		}
	}
}

func Mark(path string, info os.FileInfo, err error) error {
	if info == nil { // TODO(aoeu): Why can this case happen (resulting in panics)?
		return nil
	}
	switch {
	case info.IsDir():
		return nil
	case err != nil && err != os.ErrPermission:
		return err
	case filenameRegexp.Match([]byte(info.Name())):
		paths <- path
		return nil
	default:
		return errNotFound
	}
}

// Code below this line is modified excerpt from
// the standard library package named "filepath."

func Walk(root string, walkFn filepath.WalkFunc) error {
	var e error
	info, err := os.Lstat(root)
	switch {
	case err == nil:
		e = walk(root, info, walkFn)
	case err == errNotFound:
		e = err
	default:
		e = walkFn(root, nil, err)
	}
	return e
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
