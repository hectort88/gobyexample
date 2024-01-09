// Starting with version 1.18,
// Go has added support for generics,
// also known as type parameters.
package main

import "fmt"

// takes a map of any type and returns a slice of its keys.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(value T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: value}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: value}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elements []T
	for e := lst.head; e != nil; e = e.next {
		elements = append(elements, e.val)
	}
	return elements
}

func myGenerics() {
	separator("Generics")
	m := map[int]string{1: "1", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	fmt.Println("list:", lst.GetAll())
}
