package main

import "fmt"

type person struct {
	name       string
	fathername string
	gender     string
	age        int
	height     int
	weight     int
}

func main() {
	p := person{
		name:       "optimus",
		fathername: "primz",
		gender:     "male",
		age:        24,
		height:     178,
		weight:     54,
	}

	fmt.Println(p)
}
