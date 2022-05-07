package main

import "fmt"

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float32
}

func main() {
	e1 := Employee{
		ID:       1,
		Name:     "optimus",
		Position: "Developer",
		Salary:   100000,
	}
	e2 := Employee{
		ID:       2,
		Name:     "primz",
		Position: "Tester",
		Salary:   200000,
	}
	e3 := Employee{
		ID:       3,
		Name:     "saravana",
		Position: "Manager",
		Salary:   300000,
	}

	emap := make(map[Employee]int)

	emap[e1] = 1
	emap[e2] = 2
	emap[e3] = 3

	fmt.Println(emap)

	for e, v := range emap {
		fmt.Printf("Employee ID : %d\n", e.ID)
		fmt.Printf("Employee Name: %s\n", e.Name)
		fmt.Printf("Employee Position: %s\n", e.Position)
		fmt.Printf("Employee Salary: %f\n", e.Salary)
		fmt.Printf("Employee Value : %d\n", v)
	}

	delete(emap, e3)

	fmt.Println(emap)

}
