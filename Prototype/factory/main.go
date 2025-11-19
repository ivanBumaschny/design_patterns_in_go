package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Employee{}
	_ = d.Decode(&result)
	return &result
}

var mainOffice = Employee{
	"", Address{0, "123 East Dr", "London"},
}

var auxOffice = Employee{
	"", Address{0, "66 West Dr", "London"},
}

// Utility function, thats why new is in lower case
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func main() {
	alice := NewMainOfficeEmployee("Alice", 100)
	bob := NewAuxOfficeEmployee("Bob", 200)

	fmt.Println(*alice)
	fmt.Println(*bob)
}
