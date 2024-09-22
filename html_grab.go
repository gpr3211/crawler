package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(baseUrl string) (string, error) {

	resp, err := http.Get(baseUrl)
	// HEALTH CHECKS
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error during GET request  (func getHTML) -- %v", err))
	}
	ok := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if !ok {
		return "", errors.New(fmt.Sprintf("Bad Response (func getHTML) -- %v", resp.Status))
	}

	head := resp.Header.Get("Content-Type")
	HeadCheck := strings.Contains(head, "text/html")

	if !HeadCheck {
		return "", errors.New(fmt.Sprintf("Bad Header (func getHTML) -- %v", head))
	}
	// OK
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil

}
