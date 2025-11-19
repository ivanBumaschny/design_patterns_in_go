package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

/// ------ We want to implement different factories for each role in the company

// Functional approach (return a function not an object)
// This is more idiomatic in Go, and is good when you dont need to hold a state. Does not implement a Factory interface.
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:         name,
			Position:     position,
			AnnualIncome: annualIncome,
		}
	}
}

// Structural approach (return an object that implements a Factory interface)
// This is more appropriate when you need to hold state or have more complex logic in the factory itself. Usually implements a Factory interface.
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (f *EmployeeFactory) CreateEmployee(name string) *Employee {
	return &Employee{
		Name:         name,
		Position:     f.Position,
		AnnualIncome: f.AnnualIncome,
	}
}

func NewEmployeeFactoryStruct(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{
		Position:     position,
		AnnualIncome: annualIncome,
	}
}

func main() {
	// These are functions
	// We only provide the Dinamic parts (name) when creating an employee, since the rest of the data is "fixed" (at least not that much variation) and can be decided on the fly like this. In this cases, functional implementation is more concise.
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Alice")
	manager := managerFactory("Bob")

	println(developer.Name, "is a", developer.Position, "earning", developer.AnnualIncome)
	println(manager.Name, "is a", manager.Position, "earning", manager.AnnualIncome)

	// This is a Factory object
	// In this case, we might want to have more complex logic in the factory itself, so having an object that can hold state and methods is more appropriate. For example, you can access the Annual income via the factory itself (bossFactory.AnnualIncome) which you cant with the functional approach.
	// For a Specialized object like this, the reciever must already know the inner workings, in this case that there is a CreateEmployee method thats necesary to create the object... A good practice would be to implement a Factory interface that exposes only the CreateEmployee method.
	bossFactory := NewEmployeeFactoryStruct("CEO", 100000)
	boss := bossFactory.CreateEmployee("Charlie")
	fmt.Println(boss)
}
