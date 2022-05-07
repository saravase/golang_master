package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Gender  string
	Address Address
}

type Address struct {
	City    string
	Pincode int
}

func main() {

	a := Address{
		City:    "Chennai",
		Pincode: 654763,
	}
	p := &Person{
		Name:    "optimus",
		Age:     34,
		Gender:  "Male",
		Address: a,
	}

	d, _ := json.Marshal(p)
	fmt.Println(string(d))

	d, _ = json.MarshalIndent(p, "", " ")
	fmt.Println(string(d))

}
