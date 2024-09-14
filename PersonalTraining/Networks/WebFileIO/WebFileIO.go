package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://services.explorecalifornia.org/json/tours.php"

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}
	client := &http.Client{Transport: tr}
	//resp, err := client.Get(url)
	req, err := http.NewRequest("Get", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")
	CheckError(err)
	resp, err := client.Do(req)
	CheckError(err)
	fmt.Printf("Response Type: %T\n", resp)
	//fmt.Println("Web File IO")
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	CheckError(err)
	content := string(bytes)
	fmt.Println(content)

	tours := toursFromJson(content)

	for _, tour := range tours {
		fmt.Print(tour.Name)
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader((content)))
	_, err := decoder.Token()
	CheckError(err)
	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		CheckError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}
