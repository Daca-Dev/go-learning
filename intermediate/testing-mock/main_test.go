package main

import "testing"

func TestGetFulltimeEmployeeById(t *testing.T) {
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
				GetPersonByDNI = func(DNI string) (Person, error) {
					return Person{
						name: "David",
						age:  26,
						DNI:  "1",
					}, nil
				}

				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:      1,
						postion: "Senior Technician developer",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					name: "David",
					age:  26,
					DNI:  "1",
				},
				Employee: Employee{
					id:      1,
					postion: "Senior Technician developer",
				},
			},
		},
	}
	// rewrite your functions (mocking)
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	// made the test
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeId(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}
		if ft != test.expectedEmployee {
			t.Errorf("The FullTimeEmployee isn't the same as espected")
		}
	}

	// return the original functions
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
