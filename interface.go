package main

import "fmt"
type shaper interface{
	Area() float32
}

type square struct{
	side float32
}

func  (sq *square) Area() float32{
	return sq.side * sq.side
}

func interfaceMain () {
sq1:= new(square)
sq1.side = 5
var areaIntF shaper
areaIntF = sq1
// shorter form
// areaIntF := shaper(sq1)
  fmt.Printf("The square has area: %f\n", areaIntF.Area())
}




