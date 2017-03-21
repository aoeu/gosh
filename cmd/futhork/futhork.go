package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [TOKEN]...

'{{.}}' transliterates Latin letters to futhork (Medieval) runes
of text tokens provided as arguments or via standard input and
prints the translation to standard output.

Any letter that does not have a known translation is copied verbatim
to output (amidst any translated letters).

Examples:

	{{.}} HELLO world

	{{.}} < beowulf.txt

	escribe 'https://en.wikipedia.org/wiki/Medieval_runes' | {{.}}

`

type re struct {
	regexp *regexp.Regexp
	repl   string
}

func (r *re) replace(src string) string {
	return r.regexp.ReplaceAllString(src, r.repl)
}

type res []re

func (rr res) replace(src string) string {
	out := src
	for _, r := range rr {
		out = r.replace(out)
	}
	return out
}

var thornRegexp = res{
	re{regexp: regexp.MustCompile(`([^h])[Tt][Hh]`), repl: "${1}Þ"},
	re{regexp: regexp.MustCompile(`^[Tt][Hh]`), repl: "Þ"},
}

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	var s string
	switch {
	case len(flag.Args()) == 0:
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			fmt.Print(transcribe(in.Text()))
		}
	default:
		s = strings.Join(flag.Args(), " ")
		fmt.Print(transcribe(s))
	}
	fmt.Print("\n")
}

func transcribe(s string) string {
	return strings.Map(
		func(in rune) rune {
			if out, ok := m[in]; ok {
				return out
			}
			return in
		}, thornRegexp.replace(s))
}

var m = map[rune]rune{
	'A': 'ᛆ',
	'B': 'ᛒ',
	'C': 'ᛍ',
	'D': 'ᛑ',
	'Ð': 'ᚧ',
	'E': 'ᛂ',
	'F': 'ᚠ',
	'G': 'ᚵ',
	'H': 'ᚼ',
	'I': 'ᛁ',
	'K': 'ᚴ',
	'L': 'ᛚ',
	'M': 'ᛘ',
	'N': 'ᚿ',
	'O': 'ᚮ',
	'P': 'ᛔ',
	'Q': 'ᛩ',
	'R': 'ᚱ',
	'S': 'ᛌ',
	'T': 'ᛐ',
	'U': 'ᚢ',
	'V': 'ᚡ',
	'W': 'ᚥ',
	'X': 'ᛪ',
	'Y': 'ᛦ',
	'Z': 'ᛎ',
	'Þ': 'ᚦ',
	'Æ': 'ᛅ',
	'Ø': 'ᚯ',
	'Ä': 'ᛅ',
	'Ö': 'ᚯ',

	'a': 'ᛆ',
	'b': 'ᛒ',
	'c': 'ᛍ',
	'd': 'ᛑ',
	'ð': 'ᚧ',
	'e': 'ᛂ',
	'f': 'ᚠ',
	'g': 'ᚵ',
	'h': 'ᚼ',
	'i': 'ᛁ',
	'k': 'ᚴ',
	'l': 'ᛚ',
	'm': 'ᛘ',
	'n': 'ᚿ',
	'o': 'ᚮ',
	'p': 'ᛔ',
	'q': 'ᛩ',
	'r': 'ᚱ',
	's': 'ᛌ',
	't': 'ᛐ',
	'u': 'ᚢ',
	'v': 'ᚡ',
	'w': 'ᚥ',
	'x': 'ᛪ',
	'y': 'ᛦ',
	'z': 'ᛎ',

	'ᛆ': 'a',
	'ᛒ': 'b',
	'ᛍ': 'c',
	'ᛑ': 'd',
	'ᚧ': 'ð',
	'ᛂ': 'e',
	'ᚠ': 'f',
	'ᚵ': 'g',
	'ᚼ': 'h',
	'ᛁ': 'i',
	'ᚴ': 'k',
	'ᛚ': 'l',
	'ᛘ': 'm',
	'ᚿ': 'n',
	'ᚮ': 'o',
	'ᛔ': 'p',
	'ᛩ': 'q',
	'ᚱ': 'r',
	'ᛌ': 's',
	'ᛐ': 't',
	'ᚢ': 'u',
	'ᚡ': 'v',
	'ᚥ': 'w',
	'ᛪ': 'x',
	'ᛦ': 'y',
	'ᛎ': 'z',
}
