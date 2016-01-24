
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

### bigbig
```
usage: bigbig [-top 20] [-under /path/to/a/directory]

bigbig walks the current or provided directory, and prints out the top N
files by largest size, in descending order.

  -rightjustify
	Align file paths to the right in output
  -top int
	The top number of files to output. (default 10)
  -under string
	The root directory to run from.
```
### imagebounds
```

usage: imagebounds [image.png image.gif imagejpg ...]

imagebounds takes a list of PNG, GIF, and JPG files and prints their pixel
boundary dimenions.

```
### leamos
```
usage: leamos http://example.com/index.html

leamos downloads the file at the specified web URL and converts any HTML to
plain text.

example: leamos
https://www.reddit.com/r/NoStupidQuestions/comments/1t4niz/why_are_so_many_tumblr_blogs_so_unreadable/
| fmt | less

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
usage: path path to some dir

path takes a space separated list of directory names of a valid directory
tree and navigates and prints the full path with separators specific to the
host Operating System.

If the directory names do not create a complete path, a path under the user's
home directory is attempted, then a path derived from the root directory,
and finally an error is printed if none are found to be valid paths.

example: path go src encoding json

go get github.com/aoeu/gosh/cmd/path
echo "function goto { cd $(path $*); }" >> ~/.profile
source ~/.profile
goto go src net

```
### trash
```
Usage of trash:
  -any
	Trash any possible arguments, ignoring any invalid arguments
  -dirs
	Trash all valid directories supplied as arguments (or none if any
	arguments are invalid).
  -empty
	Use the arguments that are empty files or empty directories
  -files
	Trash
  -into string
	Put all trash into a specific directory. (default "/home/tasm/trash")
  -usage
	Trash all valid files supplied as arguments (or none if any arguments
	are invalid).
```
