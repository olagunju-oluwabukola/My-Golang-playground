package main

import (
	"fmt"
	"log"
	"strconv"
)

func intCon() {
	var big int64 = 129
	var little int8
	little = int8(big)
	fmt.Println(little)
}

func fltCon(){
	var x int = 100
	var y float32 = float32(x)
	fmt.Printf("%.9f\n", y)
}
func int2str (){
	b := 2
	name:= "Sammy"
	lines := 34
	fmt.Printf("the typpe of %d is %T\n", b,b)
	a:= strconv.Itoa(b)
	fmt.Printf("The tyoe of %v is a %T\n", a ,a)

	c := strconv.Itoa(12)
	fmt.Printf("The type of %v is %T\n", c,c)
	fmt.Printf("%q\n", c)
	fmt.Println("congratulations "  + name  + " you just wrote "  + strconv.Itoa(lines) +  " lines of code! ")
	points:= 34.67
	fl:= fmt.Sprint(32.6)
	fmt.Printf("the type of %v is %T\n", fl, fl)
	fmt.Println("Sammy has " + fmt.Sprint(points) + " points" )
}

func strCon (){
	lines_yesterday := "100"
	lines_today := "167"

	yesterday, err:= strconv.Atoi(lines_yesterday)
	if err != nil{
		log.Fatal(err)

	}

	today, err := strconv.Atoi(lines_today)
	if err != nil{
		log.Fatal(err)
	}
	lines_more := today - yesterday

	fmt.Println(lines_more)
}


func str() {
	a := "not a number"
	b, err := strconv.Atoi(a)
	fmt.Println(b)
	fmt.Println(err)
	fmt.Println(a)

	x:= "my string"
	y := []byte(x)
	z := string(y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
