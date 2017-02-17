package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage '{{.}} -with [program name] -commands [command; command...]

'{{.}}' executes commands with the specified program.


Example:

	{{.}} -with git -commands 'checkout release; fetch; merge gerrit/release; branch cr/draft/shiny-feature; merge --squash shiny-feature; commit -F $HOME/commitMessage.txt'

`

func main() {

	flag.Usage = gosh.UsageFunc(usageTemplate)
	args := struct {
		executableFile string
		commands       string
	}{}
	flag.StringVar(&args.executableFile, "with", "", "The executable file (command) to execute commands (sub-commands) with.")
	flag.StringVar(&args.commands, "commands", "", "The commands (sub-commands) to be executed as arguments to the executable file (command).")
	flag.Parse()
	switch {
	case args.executableFile == "":
		fmt.Fprintln(os.Stderr, "No executable file (command) specificed as an argument")
		flag.Usage()
	}
	commands := make([]string, 0)
	var in *bufio.Scanner
	switch {
	case len(args.commands) == 0:
		in = bufio.NewScanner(os.Stdin)
	default:
		in = bufio.NewScanner(strings.NewReader(args.commands))
	}
	for in.Scan() {
		if s := strings.Split(in.Text(), ";"); len(s) != 0 {
			commands = append(commands, s...)
		}
	}
	lastQuote := rune(0)
	f := func(r rune) bool {
		switch {
		case r == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			lastQuote = r
			return false
		default:
			return unicode.IsSpace(r)

		}
	}
	for _, c := range commands {
		cc := strings.FieldsFunc(c, f)
		for i := 0; i < len(cc); i++ {
			cc[i] = strings.Trim(cc[i], "\"'")
		}
		fmt.Println(c, strings.Join(cc, "|"))
		cmd := exec.Command(args.executableFile, cc...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error when running command %v %v : %v", args.executableFile, c, err)
			os.Exit(1)
		}
	}

}
