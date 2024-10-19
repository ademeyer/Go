package dog

import "fmt"

type Dog struct {
	Sound    string
	Locomate string
	Reside   string
}

func (b *Dog) Speak() {
	fmt.Println(b.Sound)
}

func (b *Dog) Move() {
	fmt.Println(b.Locomate)
}

func (b *Dog) Live() {
	fmt.Println(b.Reside)
}
