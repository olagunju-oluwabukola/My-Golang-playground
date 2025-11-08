package main

import "fmt"

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