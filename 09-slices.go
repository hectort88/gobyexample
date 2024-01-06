package main

import (
	"fmt"
	"slices"
)

func myslices() {
	separator("Slices")

	// zero length slice
	var s []string
	fmt.Println("uninit:", s, ", isNil:", s == nil, ", len0:", len(s) == 0)

	// non-zero length slice
	s = make([]string, 3)
	fmt.Println("empty:", s, ", isNil:", s == nil, ", len:", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "cd"
	fmt.Println(s)
	fmt.Println(s[2])

	s = append(s, "e")
	s = append(s, "f", "g")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := c[2:5]
	fmt.Println("sl1:", l)
	l = c[:5]
	fmt.Println("sl2:", l)
	l = c[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
