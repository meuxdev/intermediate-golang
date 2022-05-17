package main

import "fmt"

type Employee struct {
	id   int
	name string
}

// RECEIVER FUNCTIONS
// e *Employee -> reference to the class with the pointer
// this way you can create the getters and setters
// not also getters and setters, you can create any function assigned to the class.
// Use pointer receivers when you want to update the properties of the struct.
// Use value receivers when you just want to use the values of the instance but not mod.
func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) setName(name string) {
	e.name = name
}

func main() {
	e := Employee{}
	fmt.Printf("%+v\n", e)
	e.id = 1
	e.name = "Alejandro"
	fmt.Printf("%+v\n", e)
	e.setId(10) // own instance of the obj is sending as reference to the function
	fmt.Printf("%+v\n", e)
	fmt.Printf("Name: %s\n", e.getName())
	fmt.Printf("Name: %d\n", e.getId())
}
