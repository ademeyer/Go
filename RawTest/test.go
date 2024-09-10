package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var y float64 = 3.14159
	x := int(y)
	var z uint8 = 0x45
	fmt.Printf("y: value is %f, and type is %T \n", y, y)
	fmt.Printf("x: value is %d, and type is %T \n", x, x)
	fmt.Printf("z: value is %d, and type is %T \n", z, z)

	const LEN = 25
	s := make([]int, LEN)
	fmt.Printf("LEN: %d, and type is %T \n", LEN, LEN)
	fmt.Printf("s: value is %v, and type is %T \n", s, s)
}