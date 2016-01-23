package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

var usageMessage = `
usage: %v [image.png image.gif imagejpg ...]

%v takes a list of PNG, GIF, and JPG files and prints their pixel boundary dimenions.

`

func usage() {
	p := os.Args[0]
	fmt.Fprintf(os.Stderr, usageMessage, p, p)
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	for _, name := range os.Args[1:] {
		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open %s : %v", name, err)
		}
		var i image.Image
		switch strings.ToLower(filepath.Ext(name)[1:]) {
		case "jpeg", "jpg":
			if i, err = jpeg.Decode(f); err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't decode %s : %v", name, err)
			}
		case "gif":
			if i, err = gif.Decode(f); err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't decode %s : %v", name, err)
			}
		case "png":
			if i, err = png.Decode(f); err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't decode %s : %v", name, err)
			}
		}
		p := i.Bounds().Max
		fmt.Printf("%s is %v x %v pixels\n", name, p.X, p.Y)
	}
}
