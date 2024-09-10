package main

import (
	"fmt"
)

func main() {
	aInt := 42
	var ptr1 = &aInt
	fmt.Println("value of p1 = ", *ptr1)

	aFloat := 54.2

	ptr2 := &aFloat

	fmt.Println("value of p2 = ", *ptr2)

	*ptr2 = *ptr2 / 231.1

	fmt.Println("value of aFloat = ", aFloat)

}
