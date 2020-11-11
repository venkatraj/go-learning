package main

import "fmt"

// Yah! we don't need rest operator (js)
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sum1(a int, nums ...int) {
	fmt.Print(nums, " ")
	total := a
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// and no need for spread operator (js)
	sum(nums...)

	sum1(2, 1, 2)
	sum1(2, 1, 2, 3)

	nums = []int{1, 2, 3, 4}
	// and no need for spread operator (js)
	sum1(2, nums...)
}
