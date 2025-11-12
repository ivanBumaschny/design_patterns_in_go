// Dependency Inversion Principle (DIP)
// HLM should not depend on LLM
// --> High level Modules / Low Level Modules
// Both should depend on abstractions

package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
	// ..
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Low-Level Module (interface)
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// Low-Level Module
type Relationships struct {
	relations []Info
}

// Since in this case we are still operating inside the low level module, we can access the inner parts of the module without breaking the principle
func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}
func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// High-Level Module
type Research_DIP_break struct {
	// break DIP -> This should not depend on a low level module
	relationships Relationships
}

type Research struct {
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func (r *Research_DIP_break) Investigate() {
	// This breaks the principle because, if Research desides to change the way it stores the information from a Slice to anything else, the code would break. The high level structure depends on the low level structure
	relations := r.relationships.relations
	for _, rel := range relations {
		if rel.from.name == "John" && rel.relationship == Parent {
			fmt.Println("John has a child called ", rel.to.name)
		}
	}
}
func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	// r := Research_DIP_break{relationships}
	// r.Investigate()
	r := Research{&relationships}
	r.Investigate()

	// In the correct implementation, the Investigate() is a method of the low level module. If it were a high level implementatino, it would break the principle because it depends on the low level.
	// In Go, the necesary abstractions for this to "work" are interfaces
}
