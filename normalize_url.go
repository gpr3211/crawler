package main

import (
	"golang.org/x/net/html"
	"log"
	"net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	parsed, err := url.ParseRequestURI(s)
	if err != nil {
		return "", err
	}
	return parsed.Host + parsed.Path, nil
}

func getURLSfromHTML(s, baseUrl string) []string {
	var out []string

	parsed, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Println(err)
		return out
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					norm := parsedUrlString(a.Val, baseUrl)
					out = append(out, norm)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(parsed)
	return out

}
