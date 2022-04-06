package main

import (
	"fmt"
	"net/http"
	"os"
)

// func Get(url string) (resp *Response, err error)

func main() {
	resp, err := http.Get("http://127.0.0.1:5000/brands")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	bs := make([]byte, 99_999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
