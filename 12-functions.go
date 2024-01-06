package main

import "fmt"

func factorial(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func plus(x int, y int) int {
	return x + y
}

func plusPlus(x, y, z int) int {
	return plus(x, y) + z
}

func functions() {
	separator("functions")
	fmt.Println("factorial 5:", factorial(5))
	fmt.Println("factorial 7:", factorial(7))
	fmt.Println("factorial 4:", factorial(4))
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
