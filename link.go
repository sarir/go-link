package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Parse will take is an HTML doc and will return a
// slick of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	nodes := linkNodes(doc)

	var links []Link

	for _, n := range nodes {
		links = append(links, buildLink(n))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var r Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			r.Href = attr.Val
			break
		}
	}

	r.Text = text(n)

	return r
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}

	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var r []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r = append(r, linkNodes(c)...)
	}

	return r
}
