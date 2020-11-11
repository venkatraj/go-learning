package main

import "fmt"

// It seems to make it easy to handle error
// Seem we don't need <Either<Failure, Success>> type of code
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
