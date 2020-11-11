package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	test, prs := m["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("test:", test)

	// _ blank identifier, will not be used
	// the second identifier, `hasKey` is optional
	// it is used to find whether key is present
	// or not. Because non existing keys return zero'd values
	fmt.Println("no key:", m["k3"])
	_, hasKey := m["k2"]
	fmt.Println("hasKey k2?:", hasKey)

	// single line declration and initialization
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
