package main

import "fmt"

type AreaInterface interface {
	Area() float32
}

type perimeterInterface interface {
	perimeter() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	lenght float32
	width  float32
}

func (s *Square) Area() float32 {
		return s.side * s.side
	}

	func (t *Triangle) Area() float32{
		return 0.5 * t.lenght * t.width
	}

	func (s *Square) perimeter() float32{
		return 4 *s.side
	}

func shahpesMain() {
	var areaIntf AreaInterface
	var periIntf perimeterInterface
	sq := new(Square)
	sq.side = 10
	tr := new(Triangle)
	tr.lenght = 12
	tr.width = 8

	areaIntf = sq
	areaIntf = tr
	periIntf = sq
	fmt.Printf("The area of the square is %f\n", areaIntf.Area())
	fmt.Printf("The area of the triangle is  %f\n", areaIntf.Area())
	fmt.Printf("The perimeter of the square is %f\n", periIntf.perimeter())

}