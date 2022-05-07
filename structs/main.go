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

	var p1 person
	p1.name = "primz"
	p1.fathername = "optimus"
	p1.gender = "female"
	p1.age = 31
	p1.height = 180
	p1.weight = 50
}

func PrintPerson(p person){
	fmt.Println(p.name)
	fmt.Println(p.fathername)
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.height)
	fmt.Println(p.weight)
}