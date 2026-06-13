// Session 08 — methods: value vs pointer receivers, and methods on any type.
// Run:  go run examples/session08/methods/methods.go
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

// Value receiver: reads the struct (gets a copy). Good for read-only methods.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver modifying a copy — the change is LOST. (Anti-example.)
func (r Rectangle) ScaleCopy(factor float64) {
	r.Width *= factor // only changes the local copy
}

// Pointer receiver: modifies the REAL struct. Use this when mutating.
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

// Methods work on any named type, not just structs.
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("area:", rect.Area()) // 12

	rect.ScaleCopy(2)
	fmt.Printf("after ScaleCopy (unchanged): %+v\n", rect) // still 3x4

	rect.Scale(2) // Go auto-takes the address for pointer-receiver calls
	fmt.Printf("after Scale (changed):       %+v\n", rect) // 6x8
	fmt.Println("new area:", rect.Area())                  // 48

	temp := Celsius(100)
	fmt.Printf("%.1f°C = %.1f°F\n", float64(temp), temp.ToFahrenheit())
}
