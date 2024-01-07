package main

import "fmt"

type person struct {
	name string
	age  uint
}

func createPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func mystructs() {
	separator("Structs")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 24})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(createPerson("Created Person"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println("Copied element")
	fmt.Println(sp)
	sp.age = 44
	fmt.Println(s)
	fmt.Println(sp)
}
