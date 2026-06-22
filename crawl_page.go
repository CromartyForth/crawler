package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string){

	cfg.concurrencyControl <- struct{}{}
	defer cfg.wg.Done()
	defer func() {
		<- cfg.concurrencyControl
	}()

	if cfg.checkPageLen() == true{
		fmt.Println("Max page limit reached")
		return
	}

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
	if cfg.addPageVisit(normURL) == false {
		return
	}

	// get HTML from current url
	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error5: %v", err)
		return
	}

	// get urls from currentHTML
	outgoingURLs, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("error6: %v", err)
		return
	}

	for _, instURL := range outgoingURLs{
		cfg.wg.Add(1)
		go cfg.crawlPage(instURL)
	}
}

func (cfg *config) checkPageLen() (isTrue bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if len(cfg.pages) >= cfg.maxPages {
		return true
	}
	return false
}


func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool){
	
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	page, ok := cfg.pages[normalizedURL]
	if ok {
		page.linkCount += 1
		fmt.Println("URL already mapped, increasing count")
		cfg.pages[normalizedURL] = page
		return false
	} else {
		page := MyPage{linkCount: 1,}
		fmt.Println("URL is new, following link")
		cfg.pages[normalizedURL] = page
		return true
	}
}

