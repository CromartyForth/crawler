package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string){
	fmt.Printf("Processing: %v\n", rawCurrentURL)
	// is rawCurrent still within rawBaseURL?

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error2: %v", err)
		return
	}

	if cfg.baseURL.Host != parsedCurrentURL.Host {
		fmt.Printf("error3: %v", err)
		return
	}

	normURL, err := normaliseURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error4: %v", err)
		return
	}

	// create and or increase page count
	_, ok := cfg.pages[normURL]
	if ok {
		cfg.pages[normURL] += 1
		fmt.Println("URL already mapped, increasing count")
		return
	} else {
		cfg.pages[normURL] = 1
		fmt.Println("URL is new, following link")
	}

	// get HTML from current url
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error5: %v", err)
		return
	}

	// get urls from currentHTML
	currentURLs, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("error6: %v", err)
		return
	}

	for _, instURL := range currentURLs{
		cfg.crawlPage(instURL)
	}
}