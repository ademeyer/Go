package cat

import "fmt"

type Cat struct {
	Sound    string
	Locomate string
	Reside   string
}

func (b *Cat) Speak() {
	fmt.Println(b.Sound)
	fmt.Println(tel)
}

func (b *Cat) Move() {
	fmt.Println(b.Locomate)
}

func (b *Cat) Live() {
	fmt.Println(b.Reside)
}
