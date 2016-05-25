
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
Usage: 'busca regexp'

'busca' uses a regular expression to locate files with a matching name under
the current working directory.

Example:

	busca 'example.*\.txt'

```
### escribe
```
Usage: escribe URL

escribe downloads the file at the specified web URL and converts any HTML
to plain text.

Examples:

	escribe http://example.com/index.html

	escribe https://en.wikipedia.org/wiki/Readability | fmt --split-only
	--goal 50 | less

	echo 'function leamos() { escribe $1 | fmt -40 | pr -w 200 -5 |
	less; }' >> ~/.profile

```
### filtra
```
Usage: filtra [token]...

'filtra' removes lines of text from standard input that
match regular expressions provided in a space-separated list.
Any lines of text that match the regular expressions and constraints
are printed standard output.

Examples:

	find . -name '*.yava' | filtra generated-sources target test

	cat works_of_shakespeare.txt | filtra thou thee thine

	cat << EOF | filtra -all cat dog
		o
		cat
		dog
		cat dog
		dog cat
		tacocat
		dogmaomagod
		grep -v dog | grep -v cat | grep -v 'cat.*dog'
		EOF

Flags:

  -all
	Lines ommitted must match all filters (instead of any filter).
```
### imagebounds
```
Usage: imagebounds [FILE]...

imagebounds takes a list of PNG, GIF, and JPG files and prints their pixel
boundary dimenions.

Examples:

	imagebounds *.png
	imagebounds cat.gif dog.png
	find . -name '*.jpg' | xargs imagebounds

```
### largest
```
Usage: largest [-top 20] [-under /path/to/a/directory] [-in
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
Usage: list [file]...

'list' lists the files in the current directory in an actual list,
instead of columns, which is dissimilar from the 'ls' command in
Unix-like systems, but is similar to 'ls' of the Plan9
operating system (or the 'ls -1' command in Unix-like systems).

A glob expression or arbirtary list of files may be provided as arguments.

Examples:

	list
	list a*
	list *.txt
	list foo.bar *.fiz qux.baz *.buz

```
### pasa
```
Usage: pasa [regular expression]...

'pasa' prints lines of text from standard input that match
regular expressions provided in a space-separated list.

Any lines of text that do not match the regular expressions are
ommitted from standard output.

Examples:

	find . | pasa '.*\..ava' > yava_and_yavascript_filenames.txt

	cat << EOF || pasa dog > /tmp/no_cats.txt
		cat
		dog
		cat dog
		dog cat
		EOF

Flags:

  -all
	Lines accepted must match all regular expressions provided (instead
	of any).
```
### path
```
Usage: path [DIRECTORY]...

path takes a space separated list of directory names of a valid directory tree
and prints the full path with separators specific to the host Operating System.

If the directory names do not create a complete path, a path under the user's
home directory is attempted, then a path derived from the root directory,
and finally an error is printed if none are found to be valid paths.

Example:

	path go src encoding json

Recipes:

	In a Bourne-compatible shell:

		go get github.com/aoeu/gosh/cmd/path
		echo 'function goto { cd $(path $*); }' >> ~/.profile  &&
		source ~/.profile
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
Usage: trash [FILE]... [DIRECTORY]...

'trash' moves files or directories provided as arguments
to a folder, referred to as the "trash bin." The default
location of the trash bin is a directory named 'trash' in
the user's home directory.

Parent directories are created in the trash bin
corresponding to the absolute path that the specified file
or directories resided in until moved by the 'trash' command.
Additionally, a root folder is created in the trash bin
named after the date and time the trash command was run,
where the mentioned parent directories and arguments are
stored under.

This allows a user to run the 'trash' command on several
files or directories with the exact same name, even at
different points in time, with the context and absolute path
of each file represented by the final location within the
trash bin.

This enables users to restore files to their former locations
using just the 'mv' command.

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
