package main
	
import (
	"os"
	"net/url"
	"fmt"
	"sync"
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
	fmt.Printf("starting crawl of: %v\n",os.Args[1])
	
	// make pages map
	myPages := make(map[string]int)
	
	// make url.URL
	myURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("error2: %v", err)
		return
	}

	// make mutex
	myMu := sync.Mutex{}

	// make channel
	myChan := make(chan struct{})

	// make wait group
	myWg := sync.WaitGroup{}

	// Construct cfg
	cfg := config{
		pages: myPages,
		baseURL: myURL,
		mu: &myMu,
		concurrencyControl: myChan,
		wg:&myWg,
	}

	cfg.crawlPage(os.Args[1])
	

	// print results
	for key, value := range cfg.pages {
		fmt.Printf("Key: %v -- Value: %v\n", key, value)
	}
	os.Exit(0)
}