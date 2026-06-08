package main
	
import (
	"os"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	if len(os.Args) < 2{
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} 
	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %v\n",baseURL)

	// make pages map
	pages := make(map[string]int)

	crawlPage(baseURL, baseURL, pages)

	os.Exit(0)
}