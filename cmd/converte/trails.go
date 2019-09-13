package main

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
