package main

import (
	"fmt"
	"math"
)

func main() {
	var1 := math.Ceil(2.80)
	var2 := math.Ceil(0.93)
	var3 := math.Ceil(-7.80)
	var4 := math.Max(5, 9)
	var5 := math.Max(5, -9)
	var6 := math.Max(-5, -9)
	var7 := math.Min(5, 9)
	var8 := math.Min(5, -9)
	var9 := math.Min(-5, -9)
	var10 := math.Remainder(6, 9)
	var11 := math.Sqrt(4)
	var12 := math.Sqrt(16)
	var13 := math.Mod(9, 4)

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
}
