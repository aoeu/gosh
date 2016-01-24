#!/bin/sh
out=README.md

cat > $out << EOF

# gosh

Go programs to augment shell usage. Oh, my.

## Installation and Usage

* Install Go: \`https://golang.org/doc/install\`
* Open a terminal window and run a shell.
* Install an individual program: \`go get github.com/aoeu/gosh/cmd/trash\`
* Run the help page of the program: \`trash -help\`

Alternatively, install all the programs:  
\`go get github.com/aoeu/gosh/cmd/...\`


## Program Descriptions

EOF
for f in `lc cmd/`; do echo "### $f  " && echo "\`\`\`" && $f -help 2>&1 || echo "\`\`\`" ; done | fmt -s -w 80 >> $out

