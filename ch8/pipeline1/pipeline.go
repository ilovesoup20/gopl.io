package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			time.Sleep(100 * time.Millisecond)
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}
