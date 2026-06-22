package main

import (
	"net/url"
	"sync"
)

type config struct{
	pages map[string]MyPage
	baseURL *url.URL
	mu *sync.Mutex
	concurrencyControl chan struct{}
	wg *sync.WaitGroup
	maxPages int
}

type MyPage struct{
	URL string
	Heading string
	FirstParagraph string
	OutgoingLinks []string
	ImageURLs []string
	linkCount int
}



