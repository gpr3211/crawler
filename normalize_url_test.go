package main

import (
	"reflect"
	"testing"
)

type HTMltest struct {
	name      string
	inputURL  string
	inputBody string
	expected  []string
}

func TestURLfromHTML(t *testing.T) {
	tests := []HTMltest{}
	test := HTMltest{
		name:     "absolute/relative URLS",
		inputURL: "https://blog.boot.dev",
		inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
		expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
	}
	tests = append(tests, test)

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			actual := getURLSfromHTML(tc.inputBody, tc.inputURL)

			ok := reflect.DeepEqual(actual, tc.expected)
			if !ok {
				t.Errorf("Test - %v FAIL \n Actual: %v\n Expected: %v", tc.name, actual, tc.expected)
			}
		})
	}

}

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		// add more test cases here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
	return
}
