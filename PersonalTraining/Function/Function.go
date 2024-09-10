package main

import (
	"fmt"
	"strconv"
)

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println((d.Sound))
}

func main() {
	puppy := Dog{"Pitbull", 234, "WWharf!!!"}
	puppy.Speak()

	fmt.Println(calculate("10", "5.5", "+"))

}

type arith func(string) float64

// func createMathMap() map[string]arith {
// 	MathMap := make(map[string]arith)
// 	MathMap["+"] = addValues()
// 	return MathMap
// }

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// load arithmetic map
	//MathMap := createMathMap()
	// Your code goes here.
	aFloat1 := convertInputToValue(input1)
	aFloat2 := convertInputToValue(input2)

	if operation == "+" {
		return addValues(aFloat1, aFloat2)
	} else if operation == "-" {
		return subtractValues(aFloat1, aFloat2)
	} else if operation == "*" {
		return multiplyValues(aFloat1, aFloat2)
	} else if operation == "/" {
		return divideValues(aFloat1, aFloat2)
	}
	return 0
}

func convertInputToValue(input string) float64 {
	aFloat, err := strconv.ParseFloat(input, 64)
	if err == nil {
		return aFloat
	}
	fmt.Println(err)
	return 0
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
