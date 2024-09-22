package main

import "net/url"

type ParsedURL struct {
	protocol string
	username string
	password string
	hostname string
	port     string
	pathname string
	search   string
	hash     string
}

// newParsedURL takes a string and returns a nicely formatted URL struct
// to be used inside parseUrl
func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}
	pas, _ := parsedUrl.User.Password()

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: pas,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}

// Takes htmlnode value and turns into a proper URL, depending if it is internal or external link
func parsedUrl(s, baseUrl string) string {

	p := newParsedURL(s)
	if p.hostname == "" {
		return baseUrl + p.pathname
	} else {
		return p.protocol + "://" + p.hostname + p.pathname
	}
}
