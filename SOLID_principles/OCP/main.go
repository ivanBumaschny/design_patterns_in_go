// Open-Closed Principle (OCP)
// open for extension, closed for modification

package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small  Size = iota
	medium Size = iota
	large  Size = iota
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	// ..
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// If we want to add another filter, we need to add a new method that does it and extend the set up, instead of adding code to the already tested and implemented methods
// You create a specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (c SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == c.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

// Suppose you want to add a combinatory spefication, you do it by generating a new type, never modifying the ones you already have
type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) &&
		a.second.IsSatisfied(p)
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Printf("GreenProducts (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}

	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	// combinatory implementation
	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large green products:\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}

// The Types are open for extension, meaning that the Specification interface is implemented in such a way that, if you want to change/add a filter you can just add a new specification that implements the interface, whithout the need of chaning the inner workings of the interface itself
// You are unlikely, also, to chang the BetterFilter type, which continues to implement this principle
// There should be no need to change the already implemented types
