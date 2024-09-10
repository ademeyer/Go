package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)

	if err != nil {
		out <- fmt.Sprintf("%s => error: %s\n", url, err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s  ->  %s\n", url, ctype)
}

func main() {
	urls := []string{
		"https://gobyexample.com/channels",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	// create a response channel
	ch := make(chan string)

	for _, url := range urls {
		go returnType(url, ch)
	}

	// Run number of URLs times
	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}
