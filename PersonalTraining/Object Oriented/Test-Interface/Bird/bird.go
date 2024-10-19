package bird

import (
	"fmt"
)

type Bird struct {
	Sound    string
	Locomate string
	Reside   string
}

func (b *Bird) Speak() {
	fmt.Println(b.Sound)
}

func (b *Bird) Move() {
	fmt.Println(b.Locomate)
}

func (b *Bird) Live() {
	fmt.Println(b.Reside)
}
