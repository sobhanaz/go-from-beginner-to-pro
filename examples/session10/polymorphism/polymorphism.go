// Session 10 — polymorphism: one function, many types; and Stringer.
// Run:  go run examples/session10/polymorphism/polymorphism.go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct{ Radius float64 }
type Square struct{ Side float64 }

func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }
func (s Square) Area() float64 { return s.Side * s.Side }

// Works on ANY Shape — the essence of polymorphism.
func totalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

// A type satisfying fmt.Stringer prints using its String() method.
type Color struct{ R, G, B int }

func (c Color) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

func main() {
	// A slice holding different concrete types, all Shapes.
	shapes := []Shape{
		Circle{Radius: 2},
		Square{Side: 3},
		Circle{Radius: 1},
	}
	fmt.Printf("total area = %.2f\n", totalArea(shapes))

	// fmt automatically calls String().
	fmt.Println("favorite color:", Color{255, 0, 0})
}
