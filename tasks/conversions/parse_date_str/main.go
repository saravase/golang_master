package main

import (
	"fmt"
	"time"
)

func main() {
	input := "2026-02-28"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)
	fmt.Println(t.Format("02-Jan-2006"))

	t, _ = time.Parse(time.RFC1123, "Tue, 04 Feb 2008 15:04:05 MST")
	fmt.Println(t)

}
