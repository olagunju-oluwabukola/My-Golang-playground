package main

import (
	"fmt"
	"math/rand"
)

func funTest() {
	lastName := "johnny"

	fmt.Println("Hey I'm the first to execute\n")
	greeting(lastName)
	fmt.Println("I should run last\n")


}

func greeting(name string){
	fmt.Printf("Hiiii, my mame is %v\n", name)
	name = "onegirllikethat"
	fmt.Printf("my name is %v for the 3rd time\n", name)
}

func f1(a,b,c  int) int{
	return  a+b+c
}

func f2(a,b int) (int,int,int){
	n1 := a + b
	n2 := a -b
	n3 := a * b
	return n1,n2,n3
}

func f3 (){
	fmt.Println(f1(f2(20,10)))
}

func product (x,y,z int) int{
	sum, _  :=  fmt.Println(x*y*z)
	return sum

}

// fplus := func(x, y int) int { return x + y }

func f(){
	for i:=0; i<4; i++{
		g:= func(i int){
			fmt.Printf("%d", i)
		}

		g(i)
		fmt.Printf("-is a type of %T and has the value of %v\n", g,g)
	}
}

func adder(a,b int) int{
	result := a+b
	return result
}

func randomNum (){
	var num = rand.Intn(10) +1
	fmt.Printf("The random num is %v\n", num)

	 num = rand.Intn(10) +1
	fmt.Printf("The random num is %v\n", num)
}


func mul(g1,g2 int) (int, int, int) {
	// return  g1* g2, g1 + g2, g1- g2
	a1 := g1* g2
	a2 :=  g1 + g2
	a3 :=g1- g2

	return a1,a2,a3
}

// closures
func text2(myfunc func(int)int){
	fmt.Println(myfunc(4))

	t:= func (i int) int {
	return i*i
}

fmt.Println(t(5))

}



func closures(){
	test := func (x int)int {
		return x * -1
	}
	// fmt.Println(test)
	text2(test)
	fmt.Println(mul(4,7))
}
