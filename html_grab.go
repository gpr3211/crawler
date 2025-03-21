package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// getHTML takes a full url and grabs the entire response body, retured as a string
func getHTML(baseUrl string) (string, error) {

	resp, err := http.Get(baseUrl)
	// HEALTH CHECKS
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error during GET request  (func getHTML) -- %v", err))
	}
	ok := resp.StatusCode >= 200 && resp.StatusCode <= 299
	if !ok {
		log.Println(errors.New(fmt.Sprintf("Bad Response (func getHTML) -- %v", resp.StatusCode)))
		return "", nil
	}

	head := resp.Header.Get("Content-Type")
	HeadCheck := strings.Contains(head, "text/html") || strings.Contains(head, "xml") || strings.Contains(head, "json")

	if !HeadCheck {
		return "", errors.New(fmt.Sprintf("Bad Header (func getHTML) -- %v", head))
	}
	// OK
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	slog.Info("finished", slog.Group("req",
		slog.String("method", "Get"),
		slog.String("url", resp.Request.URL.Hostname())),
		slog.Int("status", resp.StatusCode),
		slog.Duration("Duration", time.Millisecond))

	return string(b), nil

}
