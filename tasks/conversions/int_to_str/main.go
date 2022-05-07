package main

import (
	"fmt"
	"strconv"
)

func main() {

	n := 34

	// strconv.Itoa
	s := strconv.Itoa(n)
	fmt.Println(s)

	// strconv.FormatInt
	s = strconv.FormatInt(int64(n), 10)
	fmt.Println(s)

	// fmt.Sprint
	s = fmt.Sprint(n)
	fmt.Println(s)
}
