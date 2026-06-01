package main

import (
	"net/url"
)

type PageData struct{
		URL string
		Heading string
		FirstParagraph string
		OutgoingLinks []string
		ImageURLs []string
	}


func extractPageData(html, pageURL string) PageData {
	scraped_data := PageData{}

	// get url
	url_obj, err := url.Parse(pageURL)
	if err != nil {
		return PageData{}
	}
	scraped_data.URL = pageURL

	// get heading
	scraped_data.Heading = getHeadingFromHTML(html)

	// get first paragraph
	scraped_data.FirstParagraph = getFirstParagraphFromHTML(html)
	
	// get outgoing links
	scraped_links, err := getURLsFromHTML(html, url_obj)
	if err != nil {
		return PageData{}
	}
	scraped_data.OutgoingLinks = append(scraped_data.OutgoingLinks, scraped_links...)

	// get image urls
	scraped_img_urls, err := getImagesFromHTML(html, url_obj)
	if err != nil {
		return PageData{}
	}

	scraped_data.ImageURLs = append(scraped_data.ImageURLs, scraped_img_urls...)

	return scraped_data

}