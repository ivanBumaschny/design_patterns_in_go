// LSP
// Liskov Substitution Principle (LSP)
// An API that works with a base class, should also work with the derived classes
// Go does not have POO so it is not directly applicable
package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

// ------- This inheritance implementation breaks the principle
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.width = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = s.width
}
func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

// The principle, in this case, would mean that we should be able to use Sized interfaces without breaking the inner workings of the objects
//--------------------

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	if actualArea != expectedArea {
		fmt.Print("Expected an Area of ", expectedArea, ", but got ", actualArea, "\n")
	} else {
		fmt.Print("Area of ", actualArea, "\n")
	}
}
func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
