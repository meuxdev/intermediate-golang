package main

import "fmt"

type Person struct {
	name string
	age  int32
}

type Employee struct {
	id int32
}

// Composition vs inheritance
type FullTimeEmployee struct {
	Person
	Employee
	enDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int32
}

func NewTemporaryEmployee(id, age, taxRate int32, name string) *TemporaryEmployee {

	return &TemporaryEmployee{
		Person: Person{
			name: name,
			age:  age,
		},
		Employee: Employee{
			id: id,
		},
		taxRate: taxRate,
	}
}

func (tempEmployee TemporaryEmployee) getMessage() string {
	return fmt.Sprintf("***************************************\nEmployee Type: Temporary Employee\n Name: %s\n Age: %d Years Old\n Employee ID: %d\n Tax Rate: %d\n****************************************", tempEmployee.name, tempEmployee.age, tempEmployee.id, tempEmployee.taxRate)
}

type PrintInfo interface {
	getMessage() string
}

// Implementing the interface
func EmployeePrintMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
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
	tEmployee := NewTemporaryEmployee(23, 32, 12, "Richard Thomson")
	EmployeePrintMessage(*tEmployee)
	EmployeePrintMessage(ftEmployee)

}
