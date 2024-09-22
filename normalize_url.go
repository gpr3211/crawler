package main

import (
	"net/url"
)

func normalizeURL(s string) (string, error) {

	parsed, err := url.ParseRequestURI(s)
	if err != nil {
		return "", err
	}
	return parsed.Host + parsed.Path, nil

}
