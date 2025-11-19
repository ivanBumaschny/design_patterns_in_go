package main

import "fmt"

// a Prototype Factory, related to the Prototype design patterns, lets use create "templates" of objects that can be cloned to create new objects. This effectively initializes objects based on a prototype instance, allowing for efficient object creation without the need to know the details of their construction.
type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
	Boss
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{
			Position:     "Developer",
			AnnualIncome: 60000,
		}
	case Manager:
		return &Employee{
			Position:     "Manager",
			AnnualIncome: 80000,
		}
	case Boss:
		return &Employee{
			Position:     "Boss",
			AnnualIncome: 100000,
		}
	default:
		panic("Unsupported Role")
	}
}
func main() {
	m := NewEmployee(Manager)
	m.Name = "Alice"
	fmt.Println(m)
}
