package main

import "fmt"

func max(num1, num2 int) int {
	if num1 > num2 {
        return num1
    }
    return num2
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func makeArr(num1, num2, num3, num4 int) []int {
	return []int{num1, num2, num3, num4}
}

func main() {

	x := max(13, 18)
	fmt.Println("The maximum value is:", x)
	a, b := swap("Adebayo", "Timileyin")
	fmt.Println(a, b)
	arr := makeArr(1, 2, 3, 4)
	fmt.Println(arr)

}