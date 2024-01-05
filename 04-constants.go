package main

import (
	"fmt"
	"math"
)

const s string = "Constant"

func constants() {
	fmt.Println("Constants--------------------")
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
