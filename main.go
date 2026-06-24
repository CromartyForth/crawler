package main
	
import (
	"os"
	"fmt"
	"sync"
	"time"
	"strconv"
	"net/url"
)

func main() {
	fmt.Println("Hello, World!")

	if len(os.Args) < 3{
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	} else if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} 
	fmt.Printf("starting crawl of: %v\n",os.Args[1])
	
	// make pages map
	myPages := make(map[string]PageData)
	
	// make url.URL
	myURL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Printf("error2: %v", err)
		return
	}

	// make mutex
	myMu := sync.Mutex{}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("couldn't convert argument to maxConcurrency int: %v", err)
	}

	// make channel
	myChan := make(chan struct{}, maxConcurrency)

	// make wait group
	myWg := sync.WaitGroup{}

	// make maxPages
	myMaxPages, err := strconv.Atoi(os.Args[3])

	// Construct cfg
	cfg := config{
		pages: myPages,
		baseURL: myURL,
		mu: &myMu,
		concurrencyControl: myChan,
		wg:&myWg,
		maxPages: myMaxPages,
	}

	start := time.Now()
	cfg.wg.Add(1)
	go cfg.crawlPage(os.Args[1])
	
	cfg.wg.Wait()
	since := time.Since(start)

	// print results
	for key, value := range cfg.pages {
		fmt.Printf("Key: %v\n	%+v\n\n", key, value)
	}
	fmt.Printf("Time taken: %v milliseconds\n", since.Milliseconds())
	os.Exit(0)
}