package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int32
}

// Composition vs inheritance
type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age of %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 24
	ftEmployee.id = 5
	fmt.Printf("%+v\n", ftEmployee)
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee.Person) // Full time employee has the person in his composition

}
