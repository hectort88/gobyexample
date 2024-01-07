package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func methods() {
	separator("Methods")
	rect := rectangle{width: 10, height: 5}
	fmt.Println(rect)
	fmt.Println("area:", rect.area())
	fmt.Println("perimeter:", rect.perimeter())
}
