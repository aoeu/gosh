package main

import (
	"strings"
)

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
