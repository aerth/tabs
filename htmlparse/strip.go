package htmlparse

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func Parse(reader io.Reader) ([]byte, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	return getPreformatted(doc)
}

func getPreformatted(doc *html.Node) ([]byte, error) {
	var buf = new(bytes.Buffer)
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "pre" {
			html.Render(buf, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return buf.Bytes(), nil
}
