package main

import (
	"flag"
	"io"
	"os"
	"strings"
)

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

func newRotReader(s string, amount int) rotReader {
	r := rotReader{}
	r.r = strings.NewReader(s)
	r.amount = amount
	return r
}

func (r rotReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot(p[i], r.amount)
	}
	return
}

var argAmount = flag.Int("r", 0, "Amount of letters to rotate (transpose) by.")
var argDecrypt = flag.Bool("d", false, "Set flag to decrypt text.")

func czar() {
	flag.Parse()
	amount := *argAmount
	if amount == 0 {
		amount = 13 // rot13 by default.
	}
	decrypt := *argDecrypt
	if decrypt {
		amount = 26 - amount
	}
	input := *argInput
	if *argInput == "" {
		input = getInput()
	}
	r := newRotReader(input, amount)
	io.Copy(os.Stdout, &r)
}
