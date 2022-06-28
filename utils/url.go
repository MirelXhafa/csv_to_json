package utils

import (
	"log"
	"net/url"

	"github.com/goware/urlx"
)

func ParseURL(url string) *url.URL {
	normalizedUrl, err := urlx.NormalizeString(url)

	if err != nil {
		log.Fatal(err)
	}

	parsedUrl, err := urlx.Parse(normalizedUrl)

	if err != nil {
		log.Fatal(err)
	}

	return parsedUrl
}
