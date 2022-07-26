package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	fmt.Println(dilbert)

	dilbert.Salary -= 5000
	fmt.Println(dilbert)

	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"

	// above statement is equivalent to this
	// (*employeeOfTheMonth).Position += " ..."

	var e1 Employee
	var e2 Employee
	fmt.Printf("%p\t%p\n", &e1, &e2)
	for i := 0; i < 20; i++ {
		var emp Employee
		fmt.Printf("%p\n", &emp)
	}

	var e3 *Employee = new(Employee)
	fmt.Printf("e3: %p\t%p\n", e3, &e3)
}
