package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
)

func scanInput() string {
	input := make([]string, 0)
	var token string
	for {
		_, err := fmt.Scan(&token)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}
		input = append(input, token)
	}
	return strings.Join(input, " ")
}

func main() {
	// TODO(aoeu): Implement
	args := struct {
		input    string
		funcName string
		amount   int
		decrypt  bool
		emote    bool
	}{}
	flag.StringVar(&args.funcName, "with", "rotate", "the name of the way to convert input text")
	flag.IntVar(&args.amount, "amount", 0, `czar: Amount of letter offsets to transpose by.
trails: The amount of trails to add`)
	flag.BoolVar(&args.decrypt, "decrypt", false, "czar: Decrypt text.")
	flag.BoolVar(&args.emote, "emote", false, "rotate: Flip text with added emoji")
	flag.Parse()
	args.input = strings.Join(flag.Args(), " ")
	if args.input == "" {
		args.input = scanInput()
	}
	var s string
	switch args.funcName {
	case "rotate", "flip":
		s = rotate(args.input, args.emote)
	case "czar", "caesar", "rot":
		s = czar(args.input, args.amount, args.decrypt)
	case "rot13":
		s = czar(args.input, 13, args.decrypt)
	case "shrink":
		s = shrink(args.input)
	case "strikethrough":
		s = strikethrough(args.input)
	case "trails":
		s = addTrails(args.input, args.amount)
	}
	fmt.Println(s)
}

func strikethrough(input string) string {
	output := make([]rune, 0)
	for _, r := range input {
		output = append(output, []rune{r, 822}...)
	}
	return string(output)
}

var miniatures = map[rune][]rune{
	'a': {' ', 867},
	'c': {' ', 872},
	'd': {' ', ' ', 873},
	'e': {' ', 868},
	'h': {' ', 874},
	'i': {' ', 869},
	'l': {' ', ' ', 7646},
	'm': {' ', 875},
	'o': {' ', ' ', 870},
	'r': {' ', 876},
	't': {' ', 877},
	'u': {' ', 871},
	'v': {' ', 878},
	'x': {' ', 879},
}

func shrink(in string) (out string) {
	for _, r := range in {
		r2, ok := miniatures[r]
		if !ok {
			r2 = []rune{r}
		}
		out += string(r2)
	}
	return
}

var trails = map[rune]rune{
	'a': 867,
	'c': 872,
	'd': 873,
	'e': 868,
	'h': 874,
	'i': 869,
	'm': 875,
	'o': 870,
	'r': 876,
	't': 877,
	'u': 871,
	'v': 878,
	'x': 879,
	'z': 7654,
}

func addTrails(in string, num int) (out string) {
	for _, r := range in {
		r2, ok := trails[r]
		if ok {
			runes := []rune{r}
			for i := 0; i < num; i++ {
				runes = append(runes, r2)
			}
			out += string(runes)
		} else {
			out += string(r)
		}
	}
	return out
}

func rot(b byte, amount int) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+byte(amount))%(z-a+1) + a
}

type rotReader struct {
	r      io.Reader
	amount int
}

func newRotReader(s string, amount int) *rotReader {
	r := rotReader{}
	r.r = strings.NewReader(s)
	r.amount = amount
	return &r
}

func (r rotReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot(p[i], r.amount)
	}
	return
}

func czar(input string, amount int, decrypt bool) string {
	if amount < 1 {
		amount = 13 // rot13 by default.
	}
	if decrypt {
		amount = 26 - amount
	}
	r := newRotReader(input, amount)
	b := new(bytes.Buffer)
	b.ReadFrom(r)
	return b.String()
}

func reverse(text string) string {
	r := []rune(text)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func rotate(s string, emote bool) string {
	rotated := strings.Map(func(r rune) rune {
		if rr, ok := rotations[r]; ok {
			return rr
		}
		return r
	}, reverse(s))
	if emote {
		rotated = "(╯°□°)╯︵" + rotated
	}
	return rotated
}

var rotations = map[rune]rune{
	'a': 'ɐ',
	'b': 'q',
	'c': 'ɔ',
	'd': 'p',
	'e': 'ǝ',
	'f': 'ɟ',
	'g': 'ƃ',
	'h': 'ɥ',
	'i': 'ɪ',
	'j': '𐅾',
	'k': 'ʞ',
	'l': 'l',
	'm': 'ɯ',
	'n': 'u',
	'o': 'o',
	'p': 'd',
	'q': 'b',
	'r': 'ɹ',
	's': 's',
	't': 'ʇ',
	'u': 'n',
	'v': 'ʌ',
	'w': 'ʍ',
	'x': 'x',
	'y': 'ʎ',
	'z': 'z',

	',':  'ʻ',
	'!':  '¡',
	'¡':  '!',
	'?':  '¿',
	'¿':  '?',
	'\'': ',',
	'"':  '«',
	'.':  '˙',
	'(':  ')',
	')':  '(',
	'[':  ']',
	']':  '[',
	'{':  '}',
	'}':  '{',

	'A': 'ᗄ',
	'B': 'ᗺ',
	'C': 'Ɔ',
	'D': 'ᗡ',
	'E': 'Ǝ',
	'F': 'ᖵ',
	'G': '⅁',
	'H': 'H',
	'I': 'I',
	'J': 'ᒋ',
	'K': 'ʞ',
	'L': 'ᒣ',
	'M': 'W',
	'N': 'N',
	'O': 'O',
	'P': 'Ԁ',
	'Q': 'ර',
	'R': 'ᖈ',
	'S': 'S',
	'T': '⊥',
	'U': 'ᑎ',
	'V': 'Ʌ',
	'W': 'M',
	'Y': '⅄',
	'Z': 'Z',

	// Reversed via Edit .s/('.*'): ('.*'),/\2: \1,/g
	'ɐ': 'a',
	// duplicate 'q': 'b',
	'ɔ': 'c',
	// duplicate 'p': 'd',
	'ǝ': 'e',
	'ɟ': 'f',
	'ƃ': 'g',
	'ɥ': 'h',
	'ɪ': 'i',
	'𐅾': 'j',
	'ʞ': 'k',
	// duplicate:  'l': 'l',
	'ɯ': 'm',
	// duplicate:  'u': 'n',
	// duplicate:  'o': 'o',
	// duplicate:  'd': 'p',
	// duplicate:  'b': 'q',
	'ɹ': 'r',
	// duplicate:  's': 's',
	'ʇ': 't',
	// duplicate:  'n': 'u',
	'ʌ': 'v',
	'ʍ': 'w',
	// duplicate:  'x': 'x',
	'ʎ': 'y',
	// duplicate:  'z': 'z',

	// duplicate:  ',': 'ʻ',
	// duplicate:  '!': '¡',
	// duplicate:  '¡': '!',
	// duplicate:  '?': '¿',
	// duplicate:  '¿': '?',
	// duplicate:  ',': '\'',
	// duplicate: 	'"': '«',
	// duplicate: 	'.': '˙',
	// duplicate: 	'(': ')',
	// duplicate: 	')': '(',
	// duplicate: 	'[': ']',
	// duplicate: 	']': '[',
	// duplicate:  '{': '}',
	// duplicate:  '}': '{',

	'ᗄ': 'A',
	'ᗺ': 'B',
	'Ɔ': 'C',
	'ᗡ': 'D',
	'Ǝ': 'E',
	'ᖵ': 'F',
	'⅁': 'G',
	// duplicate:	'H': 'H',
	// duplicate: 'I': 'I',
	'ᒋ': 'J',
	// duplicate 'ʞ': 'K',
	'ᒣ': 'L',
	// duplicate: 'W': 'M',
	// duplicate:	'N': 'N',
	// duplicate:	'O': 'O',
	'Ԁ': 'P',
	'ර': 'Q',
	'ᖈ': 'R',
	// duplicate: 'S': 'S',
	'⊥': 'T',
	'ᑎ': 'U',
	'Ʌ': 'V',
	// duplicate: 'M': 'W',
	'⅄': 'Y',
	// duplicate:	'Z': 'Z',
}
