package main

import (
	"fmt"
)

func Exists(s string, p map[string]int) bool {
	_, ok := p[s]
	return ok
}

// crawlPage recursively crawls an entire page coutning each instance of unique url
func crawlPage(baseUrl, rawcurrURL string, pages map[string]int, htmlbody map[string]string) {

	curr := newParsedURL(rawcurrURL)

	base := newParsedURL(baseUrl)

	// skip other sites
	if curr.hostname != base.hostname {
		return
	}

	normalizedCurrent, err := normalizeURL(rawcurrURL)
	if err != nil {
		fmt.Println("error normalizing URL")
		return
	}

	if Exists(normalizedCurrent, pages) {
		pages[normalizedCurrent]++
		return
	}

	pages[normalizedCurrent] = 1

	body, err := getHTML(parsedUrlString(rawcurrURL, baseUrl))
	if err != nil {
		htmlbody[normalizedCurrent] = body
		fmt.Println(err, normalizedCurrent)
		return
	}

	NextURLS := getURLSfromHTML(body, baseUrl)

	for _, next := range NextURLS {
		crawlPage(baseUrl, next, pages, htmlbody)
	}

}
