package main

import (
	"bytes"
	"io"
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

func newRotReader(s string, amount int) *rotReader {
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
