package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	p1 := &Point{1, 1}
	p2 := &Point{2, 2}
	path1 := &Path{*p1, *p2}
	p3 := &Point{3, 3}

	fmt.Printf("before: %v\n", path1)
	path1.TranslateBy(*p3, true)
	fmt.Printf("after: %v\n", path1)

	fmt.Printf("before: %v\n", path1)
	path1.TranslateBy(*p3, false)
	fmt.Printf("after: %v\n", path1)

	p4 := Point{7, 19}
	fmt.Printf("p4: %v\n", p4)
	p4.ScaleBy(2)
	fmt.Printf("p4: %v\n", p4)

	// Method value
	scaleP := p4.ScaleBy
	scaleP(2)
	fmt.Printf("p4: %v\n", p4)
	scaleP(3)
	fmt.Printf("p4: %v\n", p4)

	// Method expression
	distance := Point.Distance
	fmt.Println(distance(p4, p4))

	scale := (*Point).ScaleBy
	scale(&p4, 2)
	fmt.Printf("p4: %v\n", p4)
}
