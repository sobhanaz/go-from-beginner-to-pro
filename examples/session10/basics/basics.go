// Session 10 — interfaces: behavior, satisfied implicitly.
// Run:  go run examples/session10/basics/basics.go
package main

import (
	"fmt"
	"math"
)

// An interface is a set of method signatures.
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func main() {
	// Both Circle and Rectangle satisfy Shape automatically —
	// we never wrote "implements". The methods match, so it works.
	var s Shape

	s = Circle{Radius: 5}
	fmt.Printf("circle:    area=%.2f perimeter=%.2f\n", s.Area(), s.Perimeter())

	s = Rectangle{Width: 3, Height: 4}
	fmt.Printf("rectangle: area=%.2f perimeter=%.2f\n", s.Area(), s.Perimeter())
}
