package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	path := flag.String("file", "urls.txt", "")
	flag.Parse()
	file, err := os.ReadFile(*path)
	if err != nil {
		panic(err)
	}
	urls := strings.Split(string(file), "\n")
	respCh := make(chan int)
	errCh := make(chan error)
	for _, url := range urls {
		go ping(url, respCh, errCh)
	}
	for range urls {
		select{
			case err := <-errCh:
				fmt.Println(err)
			case resp := <-respCh:
				fmt.Println(resp)
		}
		// resp := <-respCh
		// fmt.Println(resp)
		// err = <-errCh
		// fmt.Println(err)
	}
}

func ping(url string, respCh chan int, errorCh chan error) {
	resp, err := http.Get(url)
	if err != nil {
		errorCh <- err
		return
	}
	respCh <- resp.StatusCode
}
