package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	// sub slice from slice
	fmt.Println(a[3:])

	// delete
	fmt.Println(Delete(a, 4))
	fmt.Println(Delete(a, 9))

}

func Delete(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}
