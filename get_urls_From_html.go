package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(input string, baseURL *url.URL ) ([]string, error) {

	foundURLs := []string{}

	// make a goquery document
	htmldoc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		return []string{}, fmt.Errorf("Error making goquery document: %v", err)
	}

	// Find method returns a selection type?
	htmldoc.Find("a[href]").Each(func(i int, s *goquery.Selection){
		attr, exists := s.Attr("href")
		if exists == false {
			return
		}
		// func HasPrefix(str, pre string) bool
		if strings.HasPrefix(attr, "/") {
			foundURLs = append(foundURLs, fmt.Sprintf("%v%v", baseURL, attr))
		} else {
			foundURLs = append(foundURLs, attr)
		}
	})
	return foundURLs, nil
}