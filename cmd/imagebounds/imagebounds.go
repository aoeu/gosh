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

	"github.com/aoeu/gosh"
)

var usageTemplate = `Usage: {{.}} [FILE]...

{{.}} takes a list of PNG, GIF, and JPG files and prints their pixel boundary dimenions.

Examples:

	{{.}} *.png
	{{.}} cat.gif dog.png
	find . -name '*.jpg' | xargs {{.}}

`

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	f := flag.Args()
	if len(f) == 0 {
		flag.Usage()
	}
	for _, name := range f {
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
