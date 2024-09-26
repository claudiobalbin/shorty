package utils

import (
	"net/url"
	"regexp"
)

func ValidKey(key string) bool {
	if key == "" || len(key) < 3 {
		return false
	}

	return true
}

func ValidURL(urlStr string) bool {
	urlRegex := regexp.MustCompile(`^(http|https)://([\w-]+\.)+[\w-]+/`)

	if !urlRegex.MatchString(urlStr) {
		return false
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}
	if parsedURL.Path == "" {
		return false
	}

	return true
}
