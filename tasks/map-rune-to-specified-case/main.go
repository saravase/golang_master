package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%c \n", unicode.To(unicode.UpperCase, 'c'))
	fmt.Printf("%c \n", unicode.To(unicode.LowerCase, 'X'))
	fmt.Printf("%c \n", unicode.To(unicode.TitleCase, 'c'))
}
