package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// bytes, encoding/gob
type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}

func main() {
	alice := Person{
		name: "Alice",
		Address: &Address{
			StreetAddress: "123 London Road",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Bob", "Charlie"},
	}

	// This is a prime example of Prototype pattern via serialization.
	// We use a predefined object (alice) and create a deep copy of it via serialization. We do not care about internat pointer managing because of the serialization process.
	bob := alice.DeepCopy()
	bob.name = "Bob"
	bob.Address.StreetAddress = "321 Baker Street"
	bob.Friends = append(bob.Friends, "David")

	fmt.Println(alice, alice.Address)
	fmt.Println(bob, bob.Address)
}
