package main

import (
	"reflect"
	"testing"
)

func TestParsedUrl(t *testing.T) {
	type testParse struct {
		name, baseURL, input, expected string
	}
	var tests []testParse
	one := testParse{name: "Path test", baseURL: "https://blog.boot.dev", input: "/salmon", expected: "https://blog.boot.dev/salmon"}

	two := testParse{name: "full url test", baseURL: "https://blog.boot.dev", input: "https://not.boot.dev/salmon", expected: "https://not.boot.dev/salmon"}

	tests = append(tests, one)
	tests = append(tests, two)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := parsedUrl(tc.input, tc.baseURL)

			ok := reflect.DeepEqual(actual, tc.expected)
			if !ok {
				t.Errorf("Test %v - FAIL: Actual: %s /// Expected %s", tc.name, actual, tc.expected)
				return
			}
		})
	}

}
