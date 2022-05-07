package main

import (
	"fmt"
)

func main() {
	a := []int{1, 3, 4, 5, 6, 7, 8, 3, 5, 4, 6, 4, 6}
	m := map[int]bool{}

	for _, v := range a {
		if !m[v] {
			m[v] = true
		}
	}

	var u []int

	for k, _ := range m {
		u = append(u, k)
	}

	fmt.Println(u)

}
