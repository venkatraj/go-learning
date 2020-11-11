package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial Value:", i)

	zeroVal(i)
	fmt.Println("After calling zeroVal()", i)

	zeroPtr(&i)
	fmt.Println("After calling zeroPtr()", i)

	fmt.Println("Address of i", &i)

}
