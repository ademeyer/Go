package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

// Create new square
func NewSquare(x, y, length int) (*Square, error) {
	if x == 0 {
		return nil, fmt.Errorf("x can not be zero")
	}
	if y == 0 {
		return nil, fmt.Errorf("y can not be zero")
	}
	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than zero")
	}
	return &Square{x, y, length}, nil
}

// Move square
func (s *Square) Move(dx, dy int) {
	s.X += dx
	s.Y += dy
}

// Area of square
func (s *Square) Area() int {
	return (s.Length * s.Length)
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}
	fmt.Printf("%#v\n", s)
	s.Move(2, 3)
	fmt.Printf("%#v\n", s)
	fmt.Println("Area:", s.Area())
}
