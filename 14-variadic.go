package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func variadic() {
	separator("Variadic Functions")
	fmt.Println("sum 1, 2:", sumAll(1, 2))
	fmt.Println("sum 1, 2, 3:", sumAll(1, 2, 3))

	// from a slice
	nums := []int{2, 5, 7, 8}
	fmt.Println("sum 2, 5, 7, 8:", sumAll(nums...))
}
