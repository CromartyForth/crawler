package main

import (
	"fmt"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	

	return fmt.Sprintln("nothing here!")
}


func getFirstParagraphFromHTML(html string) string {

	// make a reader
	stringreader := strings.NewReader(html)

	// make a goquery document
	htmldoc, err := goquery.NewDocumentFromReader(stringreader)
	if err != nil {
		return ""
	}

	has_main := htmldoc.Selection.Find("<main>")
	if has_main != nil {
		fmt.Println("has main!")
		has_para := has_main.Find("<p>")
		if has_para != nil {
			fmt.Println(has_para.Contents().Nodes)
		}
	} else {
		fmt.Println("does not have main")
	}



	return fmt.Sprintln("nothing here!")
}