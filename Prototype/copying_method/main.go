package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

// An easy implementation would be to associate a DeepCopy method to each struct that needs deep copying
func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

type Person struct {
	name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p // Defines a copy of everything that _can_ be copied shallowly (by value)
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
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

	// You can organize your objects to have a DeepCopy() method. However, it is not ideal if you have to implement it for every struct that needs it. You need to check for each struct the behaviour of each member (pointer, slice, map, value, interface, etc) and implement the deep copy accordingly.
	bob := alice.DeepCopy() // This is a deep copy, address pointer is NOT shared
	bob.name = "Bob"
	bob.Address.StreetAddress = "321 Baker Street"
	bob.Friends = append(bob.Friends, "David")

	fmt.Println(alice, alice.Address)
	fmt.Println(bob, bob.Address)
}
