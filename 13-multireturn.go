package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func multireturns() {
	separator("Miltiple Returns")
	a, b := vals()
	fmt.Println(a, b)

	// discard value using _
	_, c := vals()
	fmt.Println(c)
}
