package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Data struct {
	size int
	err  error
}

func downloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}
	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadSizeWithRoutine(url string, out chan Data) {
	res, er := downloadSize(url)
	out <- Data{res, er}
}

func main() {
	start := time.Now()

	urls := []string{
		"https://raw.githubusercontent.com/ademeyer/Gricd-Alarm/master/Core/Inc/Flash_Data.h",
		"https://raw.githubusercontent.com/ademeyer/Gricd-Alarm/master/Core/Inc/Utility.h",
		"https://services.explorecalifornia.org/json/tours.php",
		//"https://s3.amazonaws.com/nyc-tlc/trip+data/green_tripdata_2020-03.csv",
		//"https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_2020-03.csv",
	}

	for _, url := range urls {
		res, err := downloadSize(url)
		fmt.Println(res, err)
	}

	fmt.Println("elapased time:", time.Since(start))

	start = time.Now()
	ch := make(chan Data)
	for _, url := range urls {
		go downloadSizeWithRoutine(url, ch)
	}

	for range urls {
		d := <-ch
		fmt.Printf("size: %d, err: %v\n", d.size, d.err)
	}

	fmt.Println("elapased time with routine:", time.Since(start))

}
