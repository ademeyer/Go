package main

import (
	"fmt"
	"math"
)

type cartItem struct {
	name     string
	price    float64
	quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItem) float64 {
	num := len(cart)
	sum := 0.0
	for i := 0; i < num; i++ {
		sum += (cart[i].price * float64(cart[i].quantity))
	}
	return math.Round(sum*100) / 100
}

func main() {
	var cart []cartItem
	apples := cartItem{"apple", 2.99, 3}
	oranges := cartItem{"oranges", 1.5, 8}
	bananas := cartItem{"bananas", 0.49, 12}

	cart = append(cart, apples, oranges, bananas)

	fmt.Println(calculateTotal(cart))

}
