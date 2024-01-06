package main

import "fmt"

func variables() {
	separator("Variables")
	// declaring with var
	var a = "Initial"
	fmt.Println(a)

	// declare multiple
	var b, c int = 3, 4
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// default value with type int 0
	var e int
	fmt.Println(e)

	f := "Linux"
	fmt.Println(f)
}
