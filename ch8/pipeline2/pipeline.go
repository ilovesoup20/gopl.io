package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			// time.Sleep(10 * time.Millisecond)
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			// time.Sleep(10 * time.Millisecond)
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// for {
	// 	fmt.Println(<-squares)
	// }
	for x := range squares {
		fmt.Println(x)
	}
}
