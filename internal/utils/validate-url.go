package utils

import (
	"net/url"
	"strings"
)

func ValidateURL(urlString string) bool {
	if !strings.HasPrefix(urlString, "http") && !strings.HasPrefix(urlString, "https") {
		return false
	}

	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}

	return (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}
