package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int){
	fmt.Printf("Processing: %v\n", rawCurrentURL)
	// is rawCurrent still within rawBaseURL?
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return 
	}

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if parsedBaseURL.Host != parsedCurrentURL.Host {
		return
	}

	normURL, err := normaliseURL(rawCurrentURL)
	if err != nil {
		return
	}

	// create and or increase page count
	_, ok := pages[normURL]
	if ok {
		pages[normURL] += 1
		fmt.Println("URL already mapped, increasing count")
		return
	} else {
		pages[normURL] = 1
		fmt.Println("URL is new, following link")
	}

	// get HTML from current url
	currentHTML, err := getHTML(normURL)
	if err != nil {
		return
	}

	// get urls from currentHTML
	currentURLs, err := getURLsFromHTML(currentHTML, parsedBaseURL)
	if err != nil {
		return
	}

	for _, instURL := range currentURLs{
		crawlPage(rawBaseURL, instURL, pages)
	}
}