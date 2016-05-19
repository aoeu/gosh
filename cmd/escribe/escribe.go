package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/aoeu/gosh"
	"github.com/jaytaylor/html2text"
)

var usageTemplate = `Usage: {{.}} URL

{{.}} downloads the file at the specified web URL and converts any HTML to plain text.

Examples:

	{{.}} http://example.com/index.html

	{{.}} https://en.wikipedia.org/wiki/Readability | fmt --split-only --goal 50 | less

	echo 'function leamos() { {{.}} $1 | fmt -40 | pr -w 200 -5 | less; }' >> ~/.profile

`

func download(u url.URL) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return make([]byte, 0, 0), err
	}
	req.Header.Set("User-Agent", "Lynx/2.8.8dev.3 libwww-FM/2.14 SSL-MM/1.4.1")
	resp, err := client.Do(req)
	if err != nil {
		return make([]byte, 0, 0), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	flag.Usage = gosh.UsageFunc(usageTemplate)
	flag.Parse()
	if len(os.Args) != 2 {
		flag.Usage()
	}
	u := os.Args[1]
	URL, err := url.Parse(u)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid URL %v : %v\n", u, err)
		os.Exit(1)
	}
	if URL.Host == "www.nytimes.com" {
		URL.RawQuery += "&pagewanted=print"
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
