
# gosh

Go programs to augment shell usage. Oh, my.

## Installation and Usage

* Install Go: `https://golang.org/doc/install`
* Open a terminal window and run a shell.
* Install an individual program: `go get github.com/aoeu/gosh/cmd/trash`
* Run the help page of the program: `trash -help`

Alternatively, install all the programs:  
`go get github.com/aoeu/gosh/cmd/...`


## Program Descriptions

### busca
```
usage: 'busca regexp'

'busca' uses a regular expression to locate files with a matching name under
the current working directory.

example:

	busca 'example.*\.txt'

```
### escribe
```
usage: escribe URL

escribe downloads the file at the specified web URL and converts any HTML
to plain text.

example:

	escribe http://example.com/index.html

	escribe https://en.wikipedia.org/wiki/Readability | fmt --split-only
	--goal 50 | less

	echo 'function leamos() { escribe $1 | fmt -40 | pr -w 200 -5 |
	less; }' >> ~/.profile

```
### filter
```
usage: filter [token]...

'filter' removes lines of text from standard input that contain any
of text tokens provided in a space-separated list. Any lines of
text do not contain the provided text tokens are printed to
standard output.

examples:

	find . -name '*.yava' | filter generated-sources target test
	cat works_of_shakespeare.txt | filter thou thee thine

flags:

  -all
	Lines ommitted must match all filters (instead of any filter).
```
### imagebounds
```

usage: imagebounds [image.png image.gif imagejpg ...]

imagebounds takes a list of PNG, GIF, and JPG files and prints their pixel
boundary dimenions.

```
### largest
```
usage: largest [-top 20] [-under /path/to/a/directory] [-in
/path/to/a/directory]

largest walks the current or provided directory, and prints out the top N
files by largest size, in descending order.

  -rightjustify
	Align file paths to the right in output
  -top int
	The top number of files to output. (default 10)
  -under string
	The directory under which to size and rank all files.
  -within string
	The directory within to size and rank all files.
```
### list
```
usage: list [file]...

'list' lists the files in the current directory in an actual list,
instead of columns, which is dissimilar from the 'ls' command in
Unix-like systems, but is similar to 'ls' of the Plan9
operating system (or the 'ls -1' command in Unix-like systems).

A glob expression or arbirtary list of files may be provided as arguments.

examples:
	list
	list a*
	list *.txt
	list foo.bar *.fiz qux.baz *.buz

```
### path
```
usage: path relative path to a directory

path takes a space separated list of directory names of a valid directory tree
and prints the full path with separators specific to the host Operating System.

If the directory names do not create a complete path, a path under the user's
home directory is attempted, then a path derived from the root directory,
and finally an error is printed if none are found to be valid paths.

example: path go src encoding json

In a Bourne-compatible shell:
go get github.com/aoeu/gosh/cmd/path
echo 'function goto { cd $(path $*); }' >> ~/.profile  && source ~/.profile
goto go src net

In fish:
go get github.com/aoeu/gosh/cmd/path
function goto
	cd (path $argv)
end
funcsave goto
goto go src net
```
### trash
```
Usage of trash:
  -any
	Trash any possible arguments, ignoring any invalid arguments.
  -dirs
	Trash all valid directories supplied as arguments (or none if any
	arguments are invalid).
  -empty
	Trash only the arguments that are empty files or empty directories.
  -files
	Trash all valid files supplied as arguments (or none if any arguments
	are invalid).
  -into string
	Put all trash into a specific directory. (default "/home/aoeu/trash")
```
