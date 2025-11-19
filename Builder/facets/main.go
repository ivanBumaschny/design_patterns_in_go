package main

import "fmt"

// Imagine we want a different builder for each segmented part of the structure
type Person struct {
	// Address
	StreetAdress, PostCode, City string

	// Job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{
		// &Person{}, -> We aggregate a person builder with different builders
		&Person{},
	}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

// We essentially built a tiny DSL (Domain Specific Language) for this structures
// Person Address Builder section
type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAdress string) *PersonAddressBuilder {
	pab.person.StreetAdress = streetAdress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}
func (pab *PersonAddressBuilder) WithPostCode(postcode string) *PersonAddressBuilder {
	pab.person.PostCode = postcode
	return pab
}

// Person Job Builder section
type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyname string) *PersonJobBuilder {
	pjb.person.CompanyName = companyname
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}
func (pjb *PersonJobBuilder) Earning(annualincome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualincome
	return pjb
}

func main() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("123 London Road").
		In("London").
		WithPostCode("LN1424").
		Works().
		At("Databricks").
		AsA("Programmer").
		Earning(120000)
	person := pb.Build()
	fmt.Println(person)
}
