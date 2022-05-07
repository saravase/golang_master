package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Index of are : %d\n", strings.Index("How are you", "are"))
	fmt.Printf("Index of am : %d\n", strings.Index("I am fine", "am"))

	fmt.Printf("Index of are : %d\n", strings.IndexAny("How are you", "yas"))
	fmt.Printf("Index of am : %d\n", strings.IndexAny("I am fine", "nll"))

	fmt.Printf("Index of are : %d\n", strings.IndexByte("How are you", 'r'))
	fmt.Printf("Index of am : %d\n", strings.IndexByte("I am fine", 'H'))
}
