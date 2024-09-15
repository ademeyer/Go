package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalf("error: can't call httpbin.org")
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("------------------------------------------------------------------------")
	// Make POST request
	job := &Job{
		User:   "Saitama",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer // what datatype is this?

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}

	resp, err = http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't make a post request - %s", err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

// Job description
type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}
