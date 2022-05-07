package main

import (
	"fmt"
	"unicode"
)

func main() {
	chars := []rune{'a', '7', '$', 'd', '!', 'b', '4', '`'}

	for _, c := range chars {
		if unicode.IsLetter(c) {
			fmt.Printf("%c is a letter\n", c)
		} else {
			fmt.Printf("%c is not a letter\n", c)
		}
	}

}
