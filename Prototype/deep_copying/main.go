package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	name    string
	Address *Address
}

func main() {
	alice := Person{
		name: "Alice",
		Address: &Address{
			StreetAddress: "123 London Road",
			City:          "London",
			Country:       "UK",
		},
	}

	bob := alice                                   // This is a shallow copy, address pointer is shared
	bob.name = "Bob"                               // ok
	bob.Address.StreetAddress = "321 Baker Street" // oops, Alice's address is changed too

	fmt.Println(alice, alice.Address)
	fmt.Println(bob, bob.Address)

	alice.Address.StreetAddress = "123 London Road" // reset

	// deep copying
	// this is the correct solution, albeit skipping the obvious part that building a new Address each time does not make any sence, and should be made more abstract (this is implemented like this here only to show the inner workings of deep copying)
	bob = alice
	bob.Address = &Address{
		StreetAddress: alice.Address.StreetAddress,
		City:          alice.Address.City,
		Country:       alice.Address.Country,
	}
	bob.name = "Bob"
	bob.Address.StreetAddress = "321 Baker Street" // ok, Alice's address is not changed

	println(alice.name, alice.Address.StreetAddress)
	println(bob.name, bob.Address.StreetAddress)
}
