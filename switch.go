package main

import (
	"fmt"
)

func switchTest (){
	switch num := 100; {
	case num > 200:
		fmt.Println("not grater than 200")

	case num < 30:
		fmt.Println("Number is not less than 30")

	case num <=90:
		fmt.Println("number is not grater than 90")

	default:
		fmt.Printf("The number is %d\n", num)
	}
}


func testCase(){
	name := "Something"

	switch {
	case name == "bukola" :
		fmt.Println("nmae is not correct")

	case name == "shade" :
		 fmt.Println("somehow close")

	default:
	fmt.Println("Name is %s", name)
	}
}