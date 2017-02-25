
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

### edita  
```
Usage 'edita filepath [filepath...]'

'edita' opens the specified files in the specified editor, or the editor
referenced in the EDITOR environment variable, or exits with an
error if no editor was specified and the EDITOR environment variable is not set.

The absolute path of the text editor may be specified as an argument,
or just the name of the text editor if the text editor exists in any directory
specified within the PATH environment variable.

Example:

	EDITOR=$PLAN9PORT/bin/acme export EDITOR && edita /tmp/file1.txt

	go get github.com/aoeu/acme/A && edita -with A /tmp/file1.txt /tmp/file2.txt

	find . -name '*.go' | edita
  -with string
	The text editor to edit text files with.
```
### escribe  
```
Usage: escribe URL

escribe downloads the file at the specified web URL and converts any HTML to plain text.

Examples:

	escribe http://example.com/index.html

	escribe https://en.wikipedia.org/wiki/Readability | fmt --split-only --goal 50 | less

	echo 'function leamos() { escribe $1 | fmt -40 | pr -w 200 -5 | less; }' >> ~/.profile

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

imagebounds takes a list of PNG, GIF, and JPG files and prints their pixel boundary dimenions.

Examples:

	imagebounds *.png
	imagebounds cat.gif dog.png
	find . -name '*.jpg' | xargs imagebounds

```
### largest  
```
Usage: largest [-top 20] [-under /path/to/a/directory] [-in /path/to/a/directory]

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
### output  
```
Usage: output [file]...

'output' reads the contents of any provided files until "end-of-file" (EOF)
and sequentially copies the file contents to standard output.

Filepaths may also be provided by standard input.

Examples:

	output /dev/urandom

	output file1 file2 file3

	echo 'test Darwin = $(uname) && man cat | grep -B3 -A1 "Rob Pike"' > /tmp/file &&
	eval $(output /tmp/file)

	find . -name '*.txt' | output

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
	Lines accepted must match all regular expressions provided (instead of any).
```
### path  
```
Usage: path [DIRECTORY]...

path takes a space separated list of directory names of a valid directory tree and prints the
full path with separators specific to the host Operating System.

If the directory names do not create a complete path, a path under the user's home directory
is attempted, then a path derived from the root directory, and finally an error is printed if
none are found to be valid paths.

Example:

	path go src encoding json

Recipes:

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
### primero  
```
Usage: primero [ [ < filepath] | [-of <filepath> ] ]

"primero" prints the first line of text from standard input or a specified file
to standard output that is not empty when trimmed of leading and trailing
whitespace characters, and exits with a non-error status.

Otherwise, "primero"  exits with an error status and prints nothing to
standard output or standard error.

The emitted line of text printed to standard output has all leading and
trailing whitespace characters removed, followed by a newline.

Examples:
	$ find $GOPATH/src -name '*.go' | primero
	> /home/username/go/src/golang.org/x/tools/benchmark/parse/parse_test.go
	$ echo $?
	> 0

	$ primero < /tmp/empty_file
	$ echo $?
	> 1

	$ touch /tmp/arbitrary_file
	$ sam -d /tmp/arbitrary_file
	>  -. /tmp/arbitrary_file
	> a
	>
	>
	>	  golang
	> awk
	>	  sed
	>    grep
	>
	> .
	> w
	> /tmp/arbitrary_file: #30
	> q
	$ primero -of /tmp/arbitrary_file
	> golang

	$ cat << EOF | primero

			first
		second
		third
		fourth
		EOF
	> first

  -of string
	A filepath to print the first line of.
```
### println  
```
Usage: 'println text'

'println' prints any provided text to standard output followed by a newline character.

Example:

	println "Hello, $PWD"
	println "Why not use" the echo command?
	'http://www.in-ulm.de/~mascheck/various/echo+printf'

```
### replace  
```
Usage: replace -all [REGULAR EXPRESSION] -with [REPLACEMENT TEXT]

"replace" reads text from standard input, searches for all text that matches
a supplied regular expression, replaces the matching text with supplied replacement text,
and outputs the resulting text to standard output.

Example:

	$ echo '123. One Two Three' | replace -all '^\d+\.' -with 'Testing:'
	> Testing: One Two Three

	$ echo '123. One Two Three' | sed -E 's/[0-9]{1,}\.[ ]{1,}/Numbers: /'
	> Numbers: One Two Three
	$ echo '123. One Two Three' | replace -with 'Numbers: ' -all '\d+\.\s+'
	> Numbers: One Two Three

	echo 'sed -E "s/[0-9]{1,}\.[ ]{1,}/Numbers: /"' | replace -all 'sed.*'	-with 'replace
	-all "\d+\.\s+" -with "Numbers: "'

  -all string
	The regular expression to search for in the input text.
  -with string
	The literal text to replace any regular expression matches with.
```
### revela  
```
Usage: 'revela regexp'

'revela' uses a regular expression to locate files with a matching name under the current
working directory.

Example:

	revela 'example.*\.txt'

```
### run  
```
Usage 'run -with <program name> -commands "<command> [command; command...]" [-on <filepath>]

'run' executes commands with the specified program.


Example:

	run -with git -commands 'add run.go; commit -m "Adding command to execute subcommands
	of provided command"'

	run -with go -commands 'fmt; vet; install' -on run.go

	$ cat << EOF | run -with git
	> add run.go
	>
	> EOF

	$ run -with git -commands "checkout release; \
	> fetch; merge gerrit/release; \
	> branch cr/draft/shiny-feature; \
	> checkout cr/draft/shiny-feature; \
	> merge --squash shiny-feature; \
	> mergetool; commit -F $HOME/commitMessage.txt"

  -commands string
	The commands (sub-commands) to be executed as arguments to the executable file (command).
  -on string
	A target path to execute the subcommands on, appended as a final arugment in the list
	of subcommands to be executed by the executable file (command).
  -with string
	The executable file (command) to execute commands (sub-commands) with.
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

Examples:

	trash $HOME/Downloads/*

	trash -any -empty -dirs *

	trash -files *.yava

	find . -name '*.yava' | trash -files

	$ # List the current directory, date, and time.
	$ pwd
	/home/aoeu/Documents
	$ date
	Wed May 25 17:15:21 EDT 2016
	$ # Discard a file using trash
	$ trash Objection.yava
	$ # Examine that the file has been moved to the trash bin.
	$ ls $HOME/trash/2016-05-25T17:15:25-04:00/home/aoeu/Documents/
	Objection.yava
	$ # Restore the file from the trash bin.
	$ mv $HOME/trash/2016-05-25T17:15:25-04:00/home/aoeu/Documents/Objection.yava
	$ ls
	Objection.yava

Flags:

  -any
	Trash any possible arguments, ignoring any invalid arguments.
  -dirs
	Trash all valid directories supplied as arguments (or none if any arguments are invalid).
  -empty
	Trash only the arguments that are empty files or empty directories.
  -files
	Trash all valid files supplied as arguments (or none if any arguments are invalid).
  -into string
	Put all trash into a specific directory. (default "/home/aoeu/trash")
```
### ultimo  
```
Usage: ultimo [ [ < filepath] | [-of <filepath> ] ]

"ultimo" prints the last line of text from standard input or a specified file
to standard output that is not empty when trimmed of leading and trailing
whitespace characters, and exits with a non-error status.

Otherwise, "ultimo"  exits with an error status and prints nothing to
standard output or standard error.

The emitted line of text printed to standard output has all leading and
trailing whitespace characters removed, followed by a newline.

Examples:
	$ find $GOPATH/src -name '*.go' | ultimo
	> /home/username/go/src/9fans.net/go/draw/draw.go
	$ echo $?
	> 0

	$ ultimo < /tmp/empty_file
	$ echo $?
	> 1

	$ touch /tmp/arbitrary_file
	$ sam -d /tmp/arbitrary_file
	>  -. /tmp/arbitrary_file
	> a
	>
	>
	>	  golang
	> awk
	>	  sed
	>    grep
	>
	> .
	> w
	> /tmp/arbitrary_file: #30
	> q
	$ ultimo -of /tmp/arbitrary_file
	> grep

	$ cat << EOF | ultimo

			first
		second
		third
		fourth
		EOF
	> fourth

  -of string
	A filepath to print the first line of.
```
