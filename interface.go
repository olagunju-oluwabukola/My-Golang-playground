package main

import (
	"fmt"
)
type shaper interface{
	Area() float32
}

type square struct{
	side float32
}

func  (sq *square) Area() float32{
	return sq.side * sq.side
}

type Rectangle struct{
	 lenght, width float32
}

func (r Rectangle) Area() float32{
	return r.lenght * r.width
}
func interfaceMain () {
	r := Rectangle{3,5}
	q := &square{8}

	shapes := []shaper{r,q}
	fmt.Println("looping shapes for area...")

	for n, _ := range shapes{
		fmt.Println("shape details:", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}
sq1:= new(square)
sq1.side = 5.

var areaIntF shaper
areaIntF = sq1

switch t := areaIntF.(type){
case *square:
	fmt.Printf("The type square %T with value %v\n", t,t)
case *Rectangle :
	fmt.Println("The type square %T with value %v", t,t)
	  default:
      fmt.Printf("Unexpected type %T", t)
    }
}

//  shorter form
//  areaIntF := shaper(sq1)
//   fmt.Printf("The square has area: %f\n", areaIntF.Area())





