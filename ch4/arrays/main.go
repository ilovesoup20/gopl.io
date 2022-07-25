package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r[2])

	// q = [4]int{1, 2, 3, 4}	// compile error: cannot assign [4]int to [3]int

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "E", GBP: "G", RMB: "R"}
	fmt.Println(RMB, symbol[RMB])

	t := [...]int{99: -1}
	fmt.Println(t)

	aa := [2]int{1, 2}
	bb := [...]int{1, 2}
	cc := [2]int{1, 3}
	fmt.Println(aa == bb, aa == cc, bb == cc) // "true false false"
	// dd := [3]int{1, 2}
	// fmt.Println(aa == dd)	// compile error: cannot compare [2]int == [3]int
}
