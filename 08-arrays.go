package main

import "fmt"

func arrays() {
	separator("Arrays")
	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// initialize with a fixed number
	b := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("dcl:", b)
	fmt.Printf("type: %T\n", b)

	// let the compiler infer the size
	c := [...]int32{3, 5, 7, 8, 9}
	fmt.Print("c: ", c)
	fmt.Printf(" has len: %d and is of type: %T\n", len(c), c)

	// Matrix
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = j + i
		}
	}
	fmt.Println(twoD)

}
