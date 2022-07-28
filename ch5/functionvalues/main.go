package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func add1(r rune) rune     { return r + 1 }

func main() {
	f := square
	fmt.Println(f(3))     // "9"
	fmt.Printf("%T\n", f) // "func(int) int"

	// var f2 func(int) int
	// f2()	// panic: call of nil function

	// Function values can be compared to nil
	var f3 func(int) int
	if f3 != nil {
		f3(3)
	}

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111
}
