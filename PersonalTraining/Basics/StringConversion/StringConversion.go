package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.
	// Convert the first string to a float64
	aFloat1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	// Convert the second string to a float64
	aFloat2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	// Calculate and return the result
	if err1 == nil && err2 == nil {
		return aFloat1 + aFloat2
	} else {
		fmt.Println("err1 : ", err1, "aFloat1", aFloat1)
		fmt.Println("err2 : ", err2, "aFloat2", aFloat2)
	}
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Number: ")
	strInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(strInput), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The value in Float is: ", aFloat)
	}

	fmt.Println("result = ", calculate("10.0", "5.5"))

}
