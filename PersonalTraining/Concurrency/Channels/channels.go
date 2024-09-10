package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// spine up a goroutine
	go func() {
		ch <- 333 // send
	}()

	// recieve in the main thread
	val := <-ch // receive

	fmt.Printf("got %d\n", val)

	// send multiple channel
	const count = 3

	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	// receive multiple
	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("Received %d\n", val)
	}

	fmt.Println("\n-----------------------------\n")
	// Send multiple but also indicate there are no more value in the channel
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// receive multiple
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}

	fmt.Println("\n-----------------------------\n")

	// Example of buffered channel
	ch2 := make(chan int, 1)

	ch2 <- 19

	val2 := <-ch2

	fmt.Println(val2)

	fmt.Println("\n-----------------------------\n")
}
