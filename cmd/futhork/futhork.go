package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	var s string
	switch {
	case len(flag.Args()) == 0:
		in := bufio.NewScanner(os.Stdin)
		for in.Scan() {
			fmt.Print(strings.Map(transcribe, in.Text()))
		}
	default:
		s = strings.Join(flag.Args(), " ")
		fmt.Print(strings.Map(transcribe, s))
	}
	fmt.Print("\n")
}

func transcribe(in rune) rune {
	if out, ok := m[in]; ok {
		return out
	}
	return in
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
}
