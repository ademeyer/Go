package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	fmt.Printf("Time type is %T\n", dt)
	fmt.Println("Current Time ", dt)
	fmt.Println("Fomated Date Time is ", dt.Format(time.ANSIC))
}
