package main

import "fmt"

func main() {
	quotes := [3]string{"All is well", "Health is wealth", "Food is everything"}
	numbers := [4]int{1, 2, 3, 4}
	digits := [5]int{1, 2, 3, 4, 5}
	count := [3]int{}
	count[0] = 1
	count[1] = 2
	count[2] = 3
	values := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println(quotes)
	fmt.Println(numbers)
	fmt.Println(digits)
	fmt.Println(count)
	fmt.Println(values)

	fmt.Println(numbers[1] + numbers[2] + digits[3])

	fmt.Println(len(values))

	// Multi dimensional arrays
	multiArray := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	multiArray1 := [...][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}, {1, 2, 3}}

	fmt.Println(multiArray)
	fmt.Println(multiArray1)
	multiArray1[2][1] = 198
	fmt.Println(multiArray1)
}
