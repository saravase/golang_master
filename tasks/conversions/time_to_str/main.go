package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().String())

	t := time.Date(2022, 03, 28, 03, 50, 16, 0, time.UTC)
	fmt.Println(t.String())
}
