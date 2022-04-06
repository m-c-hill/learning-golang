package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// func Get(url string) (resp *Response, err error)
	resp, err := http.Get("http://127.0.0.1:5000/brands")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99_999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	lw := logWriter{}

	io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs), " bytes have been written to the terminal.")
	return len(bs), nil
}
