package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC 	Celsius = -273.15
	FreezingC 		Celsius = 0
	BoilingC		Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c * 9 / 5 + 32)}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	var tmp1 Celsius = 0.0
	fmt.Println(tmp1)

	tmp1F := CToF(tmp1)
	fmt.Println(tmp1F)

	var tmp2 Fahrenheit = 72.0
	fmt.Println(tmp2)
	tmp2C := FToC(tmp2)
	fmt.Println(tmp2C)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) // compile error
	fmt.Println("c=", c, "; f=", f)
	fmt.Println("Celsius(f): ", Celsius(f))
	fmt.Println(c == Celsius(f))

	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
}