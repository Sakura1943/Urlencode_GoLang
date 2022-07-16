package cmd

import "net/url"

func UrlEscape(url_utf8 string) string {
	result := url.QueryEscape(url_utf8)
	return result
}

func UrlUnescape(url_utf_8 string) string {
	result, _ := url.QueryUnescape(url_utf_8)
	return result
}