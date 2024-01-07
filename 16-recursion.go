package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func recursion() {
	separator("Recursion")
	fmt.Println("fib 7:", fib(7))
	fmt.Println("factorial 7:", factorial(7))
}
