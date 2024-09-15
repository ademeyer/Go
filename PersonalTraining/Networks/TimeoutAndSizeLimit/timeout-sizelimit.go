package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond) // created for http breakout context timeout
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/ip", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	const mb = 1 << 20 // Limited data to read to 1mb
	r := io.LimitReader(resp.Body, mb)
	io.Copy(os.Stdout, r)
}
