package main

import (
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func justErrors() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.zouying123456789.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err != nil {
		log.Printf("some error happend")
	} else {
		log.Println("all url success")
	}
}

func main() {
	justErrors()
}
