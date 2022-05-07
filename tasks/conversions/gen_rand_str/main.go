package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("15 letters : %s\n", genRandStr(15))
	fmt.Printf("35 letters : %s\n", genRandStr(35))
	fmt.Printf("5 letters : %s\n", genRandStr(5))

}

func genRandStr(n int) string {
	ls := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)

	for i := range s {
		s[i] = ls[rand.Intn(len(ls))]
	}

	return string(s)
}
