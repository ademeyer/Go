package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

// Square
type Square struct {
	Length float64
}

// Function that calculate Area
func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shape []Shape) float64 {
	total := 0.0

	for _, s := range shape {
		total += s.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := Square{20}
	fmt.Println("Area of Square:", s.Area())
	c := Circle{10}
	fmt.Println("Area of Circle:", c.Area())
	shape := []Shape{s, c}
	fmt.Println(sumAreas(shape))

	// Capper Interface
	cap := &Capper{os.Stdout}
	fmt.Fprintln(cap, "timileyin")

}

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))

	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.wtr.Write(out)
}
