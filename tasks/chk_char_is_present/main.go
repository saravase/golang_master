package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("I am fine", "am"))
	fmt.Println(strings.ContainsAny("How are you", "how"))
	fmt.Println(strings.ContainsAny("How are you", "lhj"))
	fmt.Println(strings.ContainsAny("How are you", "w | j"))
	fmt.Println(strings.ContainsAny("How are you", "w & j"))
	fmt.Println(strings.ContainsAny("How are you", "w & u"))
}
