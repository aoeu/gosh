package main

import (
	"flag"
	"fmt"
	"strings"
)

func reverseText(text string) string {
	r := []rune(text)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func rotateRune(in rune) rune {
	out, ok := runeMap[in]
	if !ok {
		return in
	}
	return out
}

func initMap() {
	for key, value := range runeMap {
		runeMap[value] = key
	}
}

var flipIt = flag.Bool("f", false, "Rage flip it.")

func rotateText() {
	flag.Parse()
	input := *argInput
	if *argInput == "" {
		input = getInput()
	}
	initMap()
	rotated := strings.Map(rotateRune, reverseText(input))
	if *flipIt {
		rotated = "(╯°□°)╯︵" + rotated
	}
	fmt.Printf("%s\n", rotated)
}

var runeMap = map[rune]rune{
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
}
