package main

import "fmt"

func main() {
	var slice []int

	slice = make([]int, 4, 4)
	fmt.Println(len(slice), cap(slice))
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 1, 3, 4, 5)
	fmt.Println(len(slice), cap(slice))
	arr := []int{11, 13, 14, 15}
	slice = append(slice, arr...)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice[:], slice[:3]...)
	fmt.Println(len(slice), cap(slice))

	for num := range slice {
		fmt.Printf("The index %d is %d\n", num, slice[num])
	}

	for index, num := range slice {
		fmt.Printf("The index %d is %d\n", index, num)
	}

	for _, num := range slice {
		fmt.Printf("The number %d\n", num)
	}

	map1 := make(map[string]string)
	map1["A"] = "America"
	map1["I"] = "India"

	for key, val := range map1 {
		fmt.Printf("The key : %s value: %s\n", key, val)
	}

}
