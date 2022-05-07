package main

import (
	"fmt"
	"unicode"
)

func main() {
	chars := []rune{'a', '7', '$', 'd', '!', 'b', '4', '`'}

	for _, c := range chars {
		if unicode.IsPunct(c) {
			fmt.Printf("%c is a punctuation\n", c)
		} else {
			fmt.Printf("%c is not a punctuation\n", c)
		}
	}

}
