package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println(append(a1, a3...))
}
