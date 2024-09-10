package main

import (
	"fmt"
)

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("min of empty slice")
	}

	compare := values[0]

	for _, v := range values[1:] {
		if v < compare {
			compare = v
		}
	}
	return compare, nil
}

func main() {
	fmt.Println(min([]float64{1.1, 1.3, 2.4}))
	fmt.Println(min([]int{2, 5, 2, 5, 1}))
	fmt.Println(min([]string{"W", "Z", "P", "Q"}))
}
