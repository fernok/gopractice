package main

import "fmt"

// Struct takes over the role of classes in Go
type Rectangle struct {
	Name          string
	Width, Height float64
}

// Methods do not take part in the struct.
// Instead, methods are made independently outside the struct
// and receivers are used to indicate what struct the method belongs to.
func (r Rectangle) Area() {
	fmt.Println("The area of", r.Name, "is", r.Width*r.Height)
}

// Two types of receivers exist ---
// Value receivers and Pointer receivers.
// Value receivers cannot change the field value of the struct
func (r Rectangle) StayTheSame(w, h float64) {
	r.Width = w
	r.Height = h
}

// Pointer receivers can change the field value of the struct
func (r *Rectangle) Mutate(w, h float64) {
	r.Width = w
	r.Height = h
}

func main() {
	fmt.Println("Starting object oriented examples. ")

	fmt.Println("1. Initializing a struct")
	fmt.Println("... r := &Rectangle{\"rect1\", 10.0, 12.5}")
	r := &Rectangle{"rect1", 10.0, 12.5}

	fmt.Println("2. Trying the Area() method")
	fmt.Println("...r.Area()")
	fmt.Print("...")
	r.Area()

	fmt.Println("3. Trying the two types of receivers")
	fmt.Println("...r.StayTheSame(15.0, 20.0)")
	r.StayTheSame(15.0, 20.0)
	fmt.Println("...r.Area()")
	fmt.Print("...")
	r.Area()
	fmt.Println("...r.Mutate(15.0, 20.0)")
	r.Mutate(15.0, 20.0)
	fmt.Println("...r.Area()")
	fmt.Print("...")
	r.Area()
}
