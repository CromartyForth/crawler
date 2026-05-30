package main

import (
	"fmt"
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
)


func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	images := []string{}

	// make a goquery document
	htmlreader := strings.NewReader(htmlBody)

	// find all the image tags
	htmlDoc, err := goquery.NewDocumentFromReader(htmlreader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse html into goquery: %v", err)
	}

	// get the attributes for the image tags
	tags := htmlDoc.Selection.Find("img")

	// get attributes from each tag
	tags.Each(func(i int, s *goquery.Selection){
		// get tag attribute
		attr, hasAttr := s.Attr("src")
		if hasAttr != true{
			return
		}
		if strings.HasPrefix(attr, "/"){
			url := fmt.Sprintf("%v%v", baseURL, attr)
			images = append(images, url) 
		} else {
			images = append(images, attr) 
		}
	})
	return images, nil
}