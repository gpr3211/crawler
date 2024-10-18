package main

import (
	"fmt"
)

func Exists(s string, p map[string]int) bool {
	_, ok := p[s]
	return ok
}

// crawlPage recursively crawls an entire page coutning each instance of unique url
// func crawlPage(baseUrl, rawcurrURL string, pages map[string]int, htmlbody map[string]string) {
func (cfg *Config) crawlPage(rawCurrentUrl string) {

	defer cfg.wg.Done()
	curr := newParsedURL(rawCurrentUrl)

	base := newParsedURL(*cfg.baseURL)

	// skip other sites
	if curr.hostname != base.hostname {
		return
	}

	normalizedCurrent, err := normalizeURL(rawCurrentUrl)
	if err != nil {
		fmt.Println("error normalizing URL")
		return
	}

	if Exists(normalizedCurrent, cfg.pages) {
		cfg.pages[normalizedCurrent]++
		return
	}
	// create
	cfg.pages[normalizedCurrent] = 1

	// grab HTLM body from link
	body, err := getHTML(parsedUrlString(rawCurrentUrl, *cfg.baseURL))
	if err != nil {
		fmt.Println(err, normalizedCurrent)
		return
	}
	// save to map
	cfg.bodies[normalizedCurrent] = body

	// proceeds to grab links from current page
	NextURLS := getURLSfromHTML(body, *cfg.baseURL)

	//iterate through current page
	for _, next := range NextURLS {
		cfg.wg.Add(1)

		cfg.crawlPage(next)

	}

}
