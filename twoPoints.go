package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p *point) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p *point) Scale(s float64){
	p.x = p.x *s
	p.y = p.y *s
	return
}
func twoPoints() {
	p1 := new(point)
	p1.x =3
	p1.y =4
	fmt.Printf("The lenght of the vector is: %f\n", p1.Abs())

	p2 := &point{4,5}
	fmt.Printf("The lenght of vector p2 is : %f\n", p2.Abs())

		p1.Scale(5)
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs() )
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f\n", p1.x, p1.y)
}