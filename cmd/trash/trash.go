package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [FILE]... [DIRECTORY]...

'{{.}}' moves files or directories provided as arguments
to a folder, referred to as the "trash bin." The default
location of the trash bin is a directory named 'trash' in
the user's home directory.

Parent directories are created in the trash bin
corresponding to the absolute path that the specified file
or directories resided in until moved by the '{{.}}' command.
Additionally, a root folder is created in the trash bin
named after the date and time the trash command was run,
where the mentioned parent directories and arguments are
stored under.

This allows a user to run the '{{.}}' command on several
files or directories with the exact same name, even at
different points in time, with the context and absolute path
of each file represented by the final location within the
trash bin.

This enables users to restore files to their former locations
using just the 'mv' command.

Examples:

	{{.}} $HOME/Downloads/*

	{{.}} -any -empty -dirs *

	{{.}} -files *.yava

	find . -name '*.yava' | {{.}} -files

	$ # List the current directory, date, and time.
	$ pwd
	/home/aoeu/Documents
	$ date
	Wed May 25 17:15:21 EDT 2016
	$ # Discard a file using {{.}}
	$ {{.}} Objection.yava
	$ # Examine that the file has been moved to the trash bin.
	$ ls $HOME/trash/2016-05-25T17:15:25-04:00/home/aoeu/Documents/
	Objection.yava
	$ # Restore the file from the trash bin.
	$ mv $HOME/trash/2016-05-25T17:15:25-04:00/home/aoeu/Documents/Objection.yava
	$ ls
	Objection.yava

Flags:

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := newArguments()
	args.parse()
	if args.none() {
		os.Exit(0)
	}
	f := args.flags
	switch {
	case !f.any:
		if len(args.nonExistent) > 0 {
			fmt.Fprintf(os.Stderr, "Invalid paths: %v\n", args.nonExistent)
			os.Exit(1)
		}
	}
	if f.empty {
		filepaths := args.files
		args.files = make([]file, 0)
		for _, f := range filepaths {
			if f.isEmpty {
				args.files = append(args.files, f)
			} else {
				args.nonEmpty = append(args.nonEmpty, f.path)
			}
		}
	}
	if !f.any && (f.files || f.dirs || f.empty) {
		files, dirs := args.pathsByType()
		switch {
		case f.empty && len(args.nonEmpty) > 0:
			logFatal(files, args.nonEmpty, "empty")
		case f.files && !f.dirs && len(dirs) > 0:
			logFatal(files, dirs, "files")
		case f.dirs && !f.files && len(files) > 0:
			logFatal(dirs, files, "dirs")
		default:
			break
		}
	}
	trashbin, err := filepath.Abs(args.into)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not determine absolute file path of the destination trash bin due to error: %v", err)
		os.Exit(1)
	}
	dest := fmt.Sprintf("%v/%v", trashbin, time.Now().Format(time.RFC3339))
	must(os.MkdirAll, dest)
	if err := trash(args.files, dest); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (a *arguments) parse() error {
	if flag.Parsed() {
		return nil
	}
	flag.BoolVar(&a.any, "any", false, "Trash any possible arguments, ignoring any invalid arguments.")
	flag.BoolVar(&a.empty, "empty", false, "Trash only the arguments that are empty files or empty directories.")
	flag.BoolVar(&a.dirs, "dirs", false, "Trash all valid directories supplied as arguments (or none if any arguments are invalid).")
	flag.BoolVar(&a.flags.files, "files", false, "Trash all valid files supplied as arguments (or none if any arguments are invalid).")
	trashBin := fmt.Sprintf("%v/%v", os.Getenv("HOME"), "trash")
	flag.StringVar(&a.into, "into", trashBin, "Put all trash into a specific directory.")
	flag.Parse()
	paths := flag.Args()
	if len(paths) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			if s := strings.Trim(input.Text(), "\t\n"); s != "" {
				paths = append(paths, s)
			}
		}
	}
	return a.parsePaths(paths)
}

func must(mkdirFunc func(string, os.FileMode) error, dirname string) {
	if err := mkdirFunc(dirname, os.ModePerm); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func trash(files []file, trashbin string) error {
	for _, f := range files {
		source := f.path
		dest := fmt.Sprintf("%v%v", trashbin, source)
		must(os.MkdirAll, fmt.Sprintf("%v%v", trashbin, filepath.Dir(source)))
		if err := os.Rename(source, dest); err != nil {
			return err
		}
	}
	return nil
}

func (a *arguments) pathsByType() (files, dirs []string) {
	files, dirs = make([]string, 0), make([]string, 0)
	for _, f := range a.files {
		switch {
		case f.isDir:
			dirs = append(dirs, f.path)
		default:
			files = append(files, f.path)
		}
	}
	return files, dirs
}

type arguments struct {
	flags
	files       []file
	nonExistent []string
	nonEmpty    []string
}

func (a arguments) none() bool {
	return len(a.files) == 0 && len(os.Args[1:]) == 0
}

type flags struct {
	any   bool
	empty bool
	dirs  bool
	files bool
	into  string
}

type file struct {
	path    string
	isEmpty bool
	isDir   bool
}

func (f *file) setIsEmpty(info os.FileInfo) error {
	switch {
	case !f.isDir && info.Size() == 0:
		f.isEmpty = true
	case f.isDir:
		ff, err := os.Open(f.path)
		if err != nil {
			return err
		}
		names, err := ff.Readdirnames(-1)
		if err != nil {
			return err
		}
		if len(names) == 0 {
			f.isEmpty = true
		}

	}
	return nil
}

func newArguments() *arguments {
	return &arguments{
		files:       make([]file, 0, len(os.Args[1:])),
		nonExistent: make([]string, 0),
		nonEmpty:    make([]string, 0),
	}

}

func (a *arguments) parsePaths(paths []string) error {
	for _, p := range paths {
		info, err := os.Stat(p)
		switch {
		case err != nil && os.IsNotExist(err):
			a.nonExistent = append(a.nonExistent, p)
		case err != nil:
			return err
		default:
			p, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			f := file{path: p, isDir: info.IsDir()}
			f.setIsEmpty(info)
			a.files = append(a.files, f)
		}
	}
	return nil
}

func logFatal(goodPaths, badPaths []string, pathType string) {
	fmt.Fprintf(os.Stderr, "Invalid %v: %v\n", pathType, badPaths)
	if len(goodPaths) > 0 {
		fmt.Fprintf(os.Stderr, "With only valid %v:\n	%v -%v %v\n",
			pathType, os.Args[0], pathType, strings.Join(goodPaths, " "))
	}
	os.Exit(1)
}
