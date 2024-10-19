package main

import (
	"fmt"

	animal "github.com/ademeyer/Aninal"
	bird "github.com/ademeyer/Bird"
	cat "github.com/ademeyer/Cat"
	dog "github.com/ademeyer/Dog"
)

func main() {
	B := bird.Bird{
		Sound:    "sqeaky",
		Locomate: "fly",
		Reside:   "tree",
	}
	C := cat.Cat{
		Sound:    "meow",
		Locomate: "walk",
		Reside:   "house",
	}
	D := dog.Dog{
		Sound:    "woof",
		Locomate: "walk",
		Reside:   "house",
	}
	an := []animal.Animal{
		&B,
		&C,
		&D,
	}
	for _, a := range an {
		executeMethod(a)
		fmt.Println()
	}
}

func executeMethod(fn animal.Animal) {
	fn.Speak()
	fn.Move()
	fn.Live()
}
