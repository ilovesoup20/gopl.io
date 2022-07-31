package main

import (
	"fmt"
	"math"
)

// Point ...
type Point struct{ X, Y float64 }

// Distance ... traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance ... same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Path ...
type Path []Point

// Distance returns the distance traveled along the path
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p1 := Point{5, 10}
	p2 := Point{8, 3}

	fmt.Println(Distance(p1, p2))

	fmt.Println(p1.Distance(p2))
	fmt.Println(p2.Distance(p1))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	p3 := &Point{1, 2}
	p3.ScaleBy(2)
	fmt.Printf("p3 = %v\n", p3)

	p4 := Point{2, 3}
	pptr := &p4
	pptr.ScaleBy(2)
	fmt.Printf("p4 = %v; %p\n", p4, pptr)

	p5 := Point{3, 4}
	(&p5).ScaleBy(2)
	fmt.Printf("p5 = %v\n", p5)

	p6 := Point{4, 5}
	p6.ScaleBy(3)
	fmt.Printf("p6 = %v\n", p6)
}
