package main

import (
	"fmt"
	"strings"
)

func main() {
	var1 := strings.ToUpper("primz")
	var2 := strings.ToLower("PRIMZ")
	var3 := strings.ToTitle("I am optimus primz")
	var4 := strings.TrimSpace(" space")
	var5 := strings.Trim("space space", "p")
	var6 := strings.TrimLeft("primz", "pr")
	var7 := strings.TrimRight("primz", "z")
	var8 := strings.Count("optimus primz", "p")
	var9 := strings.Contains("optimus primz", "z")
	var10 := strings.ContainsAny("optimus primz", "zmp")
	var11 := strings.Index("optimus", "m")
	var12 := strings.IndexAny("optimus primz", "p")
	var13 := strings.LastIndex("optimus primz", "p")
	var14 := strings.LastIndexAny("optimus primz", "p")
	var15 := strings.Replace("optimus primz", "p", ".", 3)
	var16 := strings.Split("optimus pimz", " ")
	var17 := strings.SplitAfter("optimus primz primz", " ")
	var18 := strings.Join([]string{"optimus", "primz", "saro"}, "-")
	var19 := strings.Repeat("optimus", 5)

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	fmt.Println(var7)
	fmt.Println(var8)
	fmt.Println(var9)
	fmt.Println(var10)
	fmt.Println(var11)
	fmt.Println(var12)
	fmt.Println(var13)
	fmt.Println(var14)
	fmt.Println(var15)
	fmt.Println(var16)
	fmt.Println(var17)
	fmt.Println(var18)
	fmt.Println(var19)
}
