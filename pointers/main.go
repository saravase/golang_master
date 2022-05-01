package main

import "fmt"

func main() {

	var var1 = 100
	var p1 = &var1
	var p2 = new(int)
	*p2 = 200

	fmt.Println(var1)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(&var1)

	*p1 = 300

	fmt.Println(var1)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(&var1)

	var p3 = &p2
	**p3 = 600

	fmt.Println(*p2)
	fmt.Println(**p3)
	fmt.Println(&p2)
	fmt.Println(&p3)

	fmt.Println(*p2 + 1)
	fmt.Println(**p3 + 5)
}
