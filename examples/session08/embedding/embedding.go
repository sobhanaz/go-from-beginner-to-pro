// Session 08 — struct embedding: composition over inheritance.
// Run:  go run examples/session08/embedding/embedding.go
package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return a.Name + " makes a sound"
}

// Dog EMBEDS Animal (no field name) — its fields and methods are promoted.
type Dog struct {
	Animal
	Breed string
}

// A constructor (NewX convention) that wires up the embedded part.
func NewDog(name, breed string) Dog {
	return Dog{
		Animal: Animal{Name: name},
		Breed:  breed,
	}
}

func main() {
	d := NewDog("Rex", "Labrador")

	// Promoted field and method — accessed as if they were Dog's own.
	fmt.Println("Name:", d.Name)     // promoted from Animal
	fmt.Println("Speak:", d.Speak()) // promoted method
	fmt.Println("Breed:", d.Breed)   // Dog's own field

	// You can still reach the embedded struct explicitly if needed.
	fmt.Println("via embed:", d.Animal.Name)
}
