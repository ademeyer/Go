package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Microsecond)

	return Bid{
		AdURL: url, //"https://adsЯus.com/19",
		Price: 0.05,
	}
}

var defaultBid = Bid{
	AdURL: "https://adsЯus.com/default",
	Price: 0.02,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)

	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Microsecond)
	defer cancel()
	url := "https://adsЯus.com/19" //"https://http..cat/418"
	bid := findBid(ctx, url)
	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Microsecond)
	defer cancel()
	url = "https://http.cat/404"
	bid = findBid(ctx, url)
	fmt.Println(bid)
}
