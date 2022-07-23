package main

import "fmt"

func f() {}

var g = "g"

func test1() {
	f := "f"
	fmt.Println(f) // "f"; local var 'f' shadows package-level func f
	fmt.Println(g) // "g"; package-level var
	// fmt.Println(h)	// compile error: undefined h
}

func test2() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println("")
}

func main() {
	test1()
	test2()
}
