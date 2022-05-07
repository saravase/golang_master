package main

import (
	"fmt"
	"strconv"
)

func main() {
	t := true

	// strconv.FormatBool
	s := strconv.FormatBool(t)
	fmt.Println(s)

	fmt.Println(fmt.Sprintf("%v", t))
}
