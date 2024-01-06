package main

import "fmt"

func ranges() {
	separator("Ranges")
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Printf("Sum of nums is %d\n", sum)

	for i, num := range nums {
		fmt.Println("index:", i, "value:", num)
	}

	kvs := map[string]string{"a": "One", "b": "Two", "c": "Three"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	gostr := "Go ⚡"
	fmt.Println("len for", gostr, ":", len(gostr))
	for i, c := range gostr {
		fmt.Println(i, c)
	}
}
