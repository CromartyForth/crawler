package main

import (
	"fmt"
	"strings"
	"net/url"
)

func normaliseURL(value string) (string, error) {
	parsed_url, err := url.Parse(value)
	if err != nil {
		return "", fmt.Errorf("Error parsing url - %v", err)
	}

	path, _:= strings.CutSuffix(parsed_url.Path, "/")

	// check for and add "www" onto host
	host := parsed_url.Host
	_, is_found := strings.CutPrefix(host, "www")

	if is_found == false{
		host = fmt.Sprintf("www.%v", host)
	}

	host_path := fmt.Sprintf("%v%v", host, path)
	return host_path, nil
}