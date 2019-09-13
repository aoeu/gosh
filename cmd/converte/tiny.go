package main

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
