package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	c := make(chan string) // Channel which commicates with strings

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		msg := "Error accessing" + url
		c <- msg
	}
	msg := "Successfully made call to " + url
	c <- msg
}
