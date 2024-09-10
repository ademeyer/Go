package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is a I/O test string"
	file, err := os.Create("./textFile.txt")
	CheckError(err)
	length, err := io.WriteString(file, content)
	CheckError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer ReadFile("./textFile.txt")
}

func ReadFile(FileName string) {
	data, err := os.ReadFile(FileName)
	CheckError(err)
	fmt.Println("", string(data))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
