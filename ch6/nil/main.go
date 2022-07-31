package main

import "fmt"

// An IntList is a linked list of integers
// A nil *IntList represents the empty list
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum ...
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
func main() {
	i1 := &IntList{1, nil}
	i2 := &IntList{2, i1}
	i3 := &IntList{3, i2}
	i4 := &IntList{4, i3}

	fmt.Println(i1.Sum())
	fmt.Println(i2.Sum())
	fmt.Println(i3.Sum())
	fmt.Println(i4.Sum())
}
