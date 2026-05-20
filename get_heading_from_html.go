package main

import (
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	
	// make a reader
	stringreader := strings.NewReader(html)

	// make a goquery document
	htmldoc, err := goquery.NewDocumentFromReader(stringreader)
	if err != nil {
		return ""
	}

	has_h1 := htmldoc.Selection.Find("h1")
	if has_h1.Length() != 0 {
		return has_h1.Text()
	}	
	has_h2 := htmldoc.Selection.Find("h2")
	if has_h2.Length() != 0 {
		return has_h2.Text()
	}

	return ""
	
}

func getFirstParagraphFromHTML(html string) string {

	// make a reader
	stringreader := strings.NewReader(html)

	// make a goquery document
	htmldoc, err := goquery.NewDocumentFromReader(stringreader)
	if err != nil {
		return ""
	}

	has_main := htmldoc.Selection.Find("main")

	if has_main.Length() != 0 {
		has_p := has_main.Find("p")
		if has_p.Length() != 0 {
			return has_p.Text()
		}
	} 

	has_pp := htmldoc.Selection.Find("p")
	if has_pp.Length() != 0 {
		return has_pp.Text()
	}
	
	return "nothing"
}
