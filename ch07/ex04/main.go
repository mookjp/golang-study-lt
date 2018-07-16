package ex04

import (
	"fmt"
	"os"

	"io"

	"golang.org/x/net/html"
)

type Reader struct {
	str   string
	index int64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.index >= int64(len(r.str)) {
		return 0, io.EOF
	}
	n = copy(b, r.str[r.index:])
	r.index += int64(n)
	return
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func FromReader(r io.Reader, w io.Writer) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		w.Write([]byte(link))
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}