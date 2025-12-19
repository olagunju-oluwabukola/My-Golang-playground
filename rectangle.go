package main

import "fmt"

type rectangle struct {
	lenght, width int
}

func (r *rectangle) Area() int {
	return r.lenght * r.width
}

func (r *rectangle) perimeter() int {
	return 2 * (r.lenght + r.width)
}

func mainRectangle() {
	rect := rectangle{4,5}
	// rect.lenght = 10
	// rect.width = 12
	fmt.Println("the value of the struct are : ", rect)
	fmt.Println("the are of the rectangle is :", rect.Area())
	fmt.Println("The perimeter of the rectangle is :", rect.perimeter())

}
