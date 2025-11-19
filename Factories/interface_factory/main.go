package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct {
	name     string
	age      int
	EyeCount int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Println("Sorry I am too tired")
}
func (p *person) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

func NewPerson(name string, age int) Person {
	if age >= 100 {
		return &tiredPerson{
			name: name,
			age:  age,
		}
	}
	return &person{
		name: name,
		age:  age,
	}
}

// This way you only expose the factory function and the interface, keeping the implementation details hidden.
func main() {
	p := NewPerson("Alice", 100)

	p.SayHello()
}
