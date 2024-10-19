package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func getRootPath(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Hello, Client!")
	} else {
		fmt.Fprintf(w, "Invalid")
	}
}

func main() {

	go func() {
		http.HandleFunc("/", getRootPath)
		fmt.Println("Server running at http://localhost:8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("error starting server: %s\n", err)
		}
	}()

	time.Sleep(2 * time.Second)

	// Make request
	for i := 0; i < 10; i++ {
		// GET request
		resp, err := http.Get("http://localhosts:8080/")
		if err != nil {
			fmt.Println("error: can't call localhost:8080")
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Printf("status code: %v, Response from server: %s\n", resp.Status, string(body))
		time.Sleep(2 * time.Second)
	}
}
