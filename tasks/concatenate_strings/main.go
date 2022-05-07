package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	s1 := "how"
	s2 := "are"
	s3 := "you"

	s4 := s1 + s2 + s3

	fmt.Println(s4)

	var s5 bytes.Buffer

	s5.WriteString(s1)
	s5.WriteString(s2)
	s5.WriteString(s3)

	fmt.Println(s5.String())

	fmt.Println(fmt.Sprintf("%s%s%s", s1, s2, s3))

	fmt.Println(strings.Join([]string{s1, s2, s3}, " "))

	var s6 strings.Builder
	s6.WriteString(s1)
	s6.WriteString(s2)
	s6.WriteString(s3)
	fmt.Println(s6.String())

}
