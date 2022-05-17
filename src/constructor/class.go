package main

import "fmt"

type Employee struct {
	id        int    // zero values -> NOO null values :( int = 0
	name      string // ""
	vacations bool   // false
}

func NewEmployee(id int, name string, vacations bool) *Employee {
	return &Employee{
		id:        id,
		name:      name,
		vacations: vacations,
	}
}

func main() {
	// Way 1
	e := Employee{}
	// Way 2

	e2 := Employee{
		id:        1,
		name:      "Name",
		vacations: true,
	}

	// Way 3
	e3 := new(Employee) // this returns a reference &e3 pointer

	fmt.Printf("%+v\n", e)
	fmt.Printf("%+v\n", e2)
	fmt.Printf("%+v\n", *e3)
	e3.id = 21
	e3.name = "Aljenadro"
	e3.vacations = false
	fmt.Printf("%+v\n", *e3)

	// Way 4 "The best" More flexible, allows you to make more things before creating the
	// 																																						object.
	e4 := NewEmployee(12345, "ProPlayer", true)
	fmt.Printf("%+v\n", *e4)

}
