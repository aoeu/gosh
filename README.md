
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

### escribe
```
usage: escribe http://example.com/index.html

escribe downloads the file at the specified web URL and converts any HTML
to plain text.

example: escribe https://en.wikipedia.org/wiki/Readability | fmt --split-only
--goal 50 | less

echo 'function leamos() { escribe $1 | fmt -40 | pr -w 200 -5 | less; }'
>> ~/.profile

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
### locate
```
usage: 'locate regexp'

locate uses a regular expression to locate files with a matching name under
the current working directory.

example: locate 'example.*.txt'
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
	Put all trash into a specific directory. (default "/home/tasm/trash")
```
