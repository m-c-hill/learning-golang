package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	c := make(chan string) // Channel which commicates with strings

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for u := range c {
		// Wait for channel to return a value and assign it to u. Then pass into checkUrl.
		go func(url string) { // Function literal
			time.Sleep(time.Second * 2)
			checkUrl(url, c)
		}(u)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error accessing" + url)
		c <- url
	}
	fmt.Println("Successfully made call to " + url)
	c <- url
}
