package main

import "fmt"

// To extend an already functional builder, we can use a functional programming approach

type Person struct {
	name, position string
}

type personMod func(*Person)
type PersonBuilder struct {
	// List of modifications that are gonna be applied to the person
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}

// Extending the functionality just by adding new "actions", without adding new builders or modifying previous ones
func (b *PersonBuilder) WorkAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

// With this setup, its very easy to add new functionalities just by extending the actions that a person can take, without the need of modifying the builder itself or adding new builders
func main() {
	pb := PersonBuilder{}
	p := pb.Called("Dmitri").WorkAsA("Developer").Build()
	fmt.Println(*p)
}
