package stdlib

import (
	"fmt"
	"testing"
)

func Test_url(t *testing.T) {
	url := "https://x.com:90/some/path.html?p=123&z=%5B1%2C2%5D#section-1"
	tests := []struct {
		fn  string
		out string
	}{
		{"url_scheme", "https"},
		{"url_host", "x.com:90"},
		{"url_port", "90"},
		{"url_path", "/some/path.html"},
		{"url_fragment", "section-1"},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT %s('%s')", test.fn, url), test.out)
		assertQuery(t, fmt.Sprintf("SELECT %s('z kljsdfabsdf ://')", test.fn), "")
	}

	assertQuery(t, fmt.Sprintf("SELECT url_param('%s', 'z')", url), "[1,2]")
	assertQuery(t, fmt.Sprintf("SELECT url_param(' kljz sdf ://', 'z')"), "")
}
