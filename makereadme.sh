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
wd=$PWD
cd cmd
for f in `/bin/ls`; do cd $f && go install && cd .. && echo "### $f  " && echo "\`\`\`" && $f -help 2>&1 | /usr/bin/fmt -s -w 80 && echo "\`\`\`" ; done >> ../$out
cd $wd
