package jsparse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/aerth/tabs"
	"golang.org/x/net/html"
)

var ErrNotParsed = errors.New("not parsed")

func Parse(reader io.Reader) ([]byte, error) {
	resp := &tabs.TabJSON{}
	oneline := grep(reader, "window.UGAPP.store.page")
	if oneline == "" {
		return nil, ErrNotParsed
	}
	oneline = strings.TrimSpace(oneline)
	oneline = oneline[strings.Index(oneline, " = "):]
	oneline = html.UnescapeString(oneline)
	if err := json.NewDecoder(strings.NewReader(oneline)).Decode(resp); err != nil {
		fmt.Println(oneline)
		return nil, err
	}
	tab := resp.Data.TabView.WikiTab.Content
	tab = strings.Replace(tab, "[ch]", "[", -1)
	tab = strings.Replace(tab, "[/ch]", "]", -1)
	return []byte(tab), nil
}

func grep(haystack io.Reader, needle string) string {
	scanner := bufio.NewScanner(haystack)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, needle) {
			return text
		}
	}
	return ""
}

func getScript(doc *html.Node) []string {
	var scripts []string
	var buf = new(bytes.Buffer)
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			html.Render(buf, n)
			scripts = append(scripts, buf.String())
			buf.Reset()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return scripts
}
