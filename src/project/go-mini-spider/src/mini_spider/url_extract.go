/* url_extract.go - for extract urls from html	*/
/*
modification history
--------------------
2015/5/17, by Zhang Miao, create
*/
/*
DESCRIPTION
*/
package mini_spider

import (
	"bytes"
	"fmt"
	"path"
	"strings"
)

import (
	"golang.org/x/net/html"
)

type HtmlLinks struct {
	links []string
}

// create new HtmlLinks
func NewHtmlLinks() *HtmlLinks {
	hl := new(HtmlLinks)
	hl.links = make([]string, 0)
	return hl
}

/*
get all href in given html node

Params:
	- n: html node
	- links: string slice for storing links in html
*/
func (hl *HtmlLinks) linksGet(n *html.Node, url string) {
	var link string

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if strings.HasPrefix(a.Val, "http://") {
					link = a.Val
				} else {
					link = path.Join(url, a.Val)
				}
			
				hl.links = append(hl.links, link)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hl.linksGet(c, url)
	}
}

/*
get url links in given html page

Params:
	- data: data for html page
	- url: url of this html page

Returns:
	(links, error)
*/
func urlExtract(data []byte, url string) ([]string, error) {
	// parse html
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("html.Parse():%s", err.Error())
	}

	// get all links
	hl := NewHtmlLinks()
	hl.linksGet(doc, url)

	return hl.links, nil
}
