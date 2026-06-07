package main

import (
	"io"
	"fmt"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	
	request, err := http.NewRequest(http.MethodGet, rawURL, http.NoBody)
	if err != nil{
		return "", fmt.Errorf("Couldn't create request: , %v", err)
	}
	 // set User-Agent to "BootCrawler/1.0"
	 request.Header.Set("User-Agent", "BootCrawler/1.0")

	 // make a client ??
	 var Client = &http.Client{}

	 resp, err := Client.Do(request)
	 if err != nil{
		return "", fmt.Errorf("Error making get request: %v", err)
	 }

	 if resp.StatusCode >= 400 {
		return "", fmt.Errorf("External Error: %v", err)
	 }

	 if resp.Header.Values("content-type")[0] != "text/html" {
		return "", fmt.Errorf("Response is not a website: %v", err)
	 }

	 // read body
	 html, err := io.ReadAll(resp.Body)
	 if err != nil{
		return "", fmt.Errorf("Could not ready body of website: %v", err)
	 }

	 return string(html), nil

}