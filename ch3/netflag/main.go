package main

import "fmt"

// Flags ...
type Flags uint

// foo
const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopBack
	FlagPointToPoint
	FlagMulticast
)

// IsUp ...
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// TurnDown ...
func TurnDown(v *Flags) { *v &^= FlagUp }

// SetBroadcast ...
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

// IsCast ...
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"

	fmt.Println(MiB)
}

//
const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)
