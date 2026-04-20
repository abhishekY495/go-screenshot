package utils

import (
	"net/url"
	"strings"
)

func ValidateURL(urlString string) bool {
	if !strings.HasPrefix(urlString, "http://") && !strings.HasPrefix(urlString, "https://") {
		return false
	}

	if !strings.Contains(urlString, ".") {
		return false
	}

	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}

	if u.Host == "" {
		return false
	}

	return true
}
