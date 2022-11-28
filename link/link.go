package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ShowLinks(links []Link) {
	for _, l := range links {
		fmt.Printf("Href: %s, Text: %s\n", l.Href, l.Text)
	}
}

func findText(n *html.Node) string {
	str := ""
	if n.Type == html.TextNode {
		str = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str += findText(c)
	}
	return str
}

func findRef(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := findText(n)
				*links = append(*links, Link{a.Val, strings.TrimSpace(text)})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findRef(c, links)
	}
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	links := make([]Link, 0)
	findRef(doc, &links)
	return links, nil
}
