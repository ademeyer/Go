package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s  ->  %s\n", url, ctype)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func siteConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func(ur string) {
			returnType(ur)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	urls := []string{
		"https://gobyexample.com/channels",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSerial(urls)
	fmt.Println("time taken:", time.Since(start))

	start = time.Now()
	siteConcurrent(urls)
	fmt.Println("time taken:", time.Since(start))
}
