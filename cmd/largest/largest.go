package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"

	"github.com/aoeu/gosh"
)

type byteSize int64

const (
	_           = iota
	KB byteSize = 1 << (10 * iota)
	MB
	GB
)

func (b byteSize) String() string {
	switch {
	case b >= GB:
		return b.fmtString("GB", GB)
	case b >= MB:
		return b.fmtString("MB", MB)
	case b >= KB:
		return b.fmtString("KB", KB)
	}
	return b.fmtString("B", b)
}

func (b byteSize) fmtString(name string, divisor byteSize) string {
	return fmt.Sprintf("%.2f"+name, float64(b)/float64(divisor))
}

type fileSize struct {
	path string
	byteSize
}

func (f fileSize) String() string {
	p, _ := filepath.Abs(f.path) // TODO(aoeu): Handle potential errors.
	return fmt.Sprintf("%s\t%s\t%s\t", f.byteSize.String(), filepath.Ext(f.path), p)
}

type fileSizes []fileSize

func (f fileSizes) Len() int           { return len(f) }
func (f fileSizes) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f fileSizes) Less(i, j int) bool { return f[i].byteSize > f[j].byteSize }

var allFileSizes fileSizes
var tabw *tabwriter.Writer

var usageTemplate = `usage: {{.}} [-top 20] [-under /path/to/a/directory] [-in /path/to/a/directory]

{{.}} walks the current or provided directory, and prints out the top N 
files by largest size, in descending order.

`

func main() {
	args := struct {
		under        string
		within       string
		num          int
		rightjustify bool
	}{}
	flag.StringVar(&args.under, "under", "", "The directory under which to size and rank all files.")
	flag.StringVar(&args.within, "within", "", "The directory within to size and rank all files.")
	//TODO(aoeu):
	// flag.StringVar(&args.within, "within", "", "The directory within to size and rank all files.")
	flag.IntVar(&args.num, "top", 10, "The top number of files to output.")
	flag.BoolVar(&args.rightjustify, "rightjustify", false, "Align file paths to the right in output")
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	if wd, err := os.Getwd(); args.under == "" && err == nil {
		args.under = wd
	}

	tabw = new(tabwriter.Writer)
	tabw.Init(os.Stdout, 8, 0, 1, ' ', tabwriter.AlignRight)
	allFileSizes = make(fileSizes, 0)

	err := filepath.Walk(args.under, mark)

	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(allFileSizes)
	for i := 0; i < len(allFileSizes) && i < args.num; i++ {
		fmt.Fprintln(tabw, allFileSizes[i])
		if !args.rightjustify {
			tabw.Flush()
		}
	}
	if args.rightjustify {
		tabw.Flush()
	}
}

func mark(path string, info os.FileInfo, err error) error {
	if info == nil || info.IsDir() {
		return nil
	}
	if err == os.ErrPermission {
		fmt.Fprintf(os.Stderr, "No permission to determine size of: '%v'\n", path)
		return nil
	}
	if err != nil {
		return err
	}
	allFileSizes = append(allFileSizes, fileSize{path, byteSize(info.Size())})
	return nil
}
