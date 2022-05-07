package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "134"

	// strconv.Atoi
	n, _ := strconv.Atoi(s)
	fmt.Println(n)

	// strconv.ParseInt
	m, _ := strconv.ParseInt(s, 10, 64)
	fmt.Println(m)

	// fmt.Sscanf
	c := "id:34"
	var o int
	fmt.Sscanf(c, "id:%d", &o)
	fmt.Println(o)
}
