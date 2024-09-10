package main

import "fmt"

func main(){
	x := 12
	y := 13

	if x == y || y > x{
		fmt.Println("x is equal to or greater than y")
	}
	if x < y {
        fmt.Println("x is less than y")
    } else {
        fmt.Println("x is not less than y")
    }

	for x:= 0; x < 10; x++ {
		fmt.Println(x)
	}
}