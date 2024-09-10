package main

import (
    "fmt"
    "bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    input, _ := reader.ReadString('\n')
    fmt.Printf("Hello, %s", input)
}