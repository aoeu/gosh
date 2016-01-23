package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/jaytaylor/html2text"
)

func usage() string {
	return fmt.Sprintf("Usage: %v http://example.com", os.Args[0])
}

func download(u url.URL) ([]byte, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return make([]byte, 0, 0), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}
	u := os.Args[1]
	URL, err := url.Parse(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL %v : %v\n", u, err)
		os.Exit(1)
	}
	b, err := download(*URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %v : %v\n", u, err)
		os.Exit(1)
	}
	out, err := html2text.FromReader(bytes.NewReader(b))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error converting HTML to text: %v", err)
		os.Exit(1)
	}
	fmt.Println(out)
}
