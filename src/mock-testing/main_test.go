package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						Name: "Alejandro Andrade",
						DNI:  "12asdf1asd2213",
						Age:  35,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "Alejandro Andrade",
					DNI:  "12asdf1asd2213",
					Age:  35,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeById
	originalGetGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetGetPersonByDni
}
