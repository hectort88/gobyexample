package main

import (
	"fmt"
	"unicode/utf8"
)

func stringrunes() {
	separator("String and runes")
	const s = "สวัสดี"
	fmt.Println(s, "len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// to get the len of the runes
	fmt.Println(s, "rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
