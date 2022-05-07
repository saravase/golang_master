package main

import "fmt"

func main() {

	fmt.Println("Printing")
	fmt.Println("multiline Strings ")
	fmt.Println("in Go!!")

	fmt.Println("Hi primz...")

	s := `
	Printing
	multiline Strings 
	in Go!!
	`
	fmt.Printf("%s", s)

	s1 := "Printing\n" +
		"multiline Strings\n" +
		"in Go!!"
	fmt.Printf("%s\n", s1)

}
