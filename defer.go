package main

import "fmt"

func func1() {
	fmt.Println("Hello")
	defer func2()
	fmt.Println("Hello world")

}

func func2(){
	fmt.Println("Hey there, I'm deferred")
}

func forTest(){
	for i :=0; i<= 5; i++{
		defer fmt.Printf("with defer : %d\n", i)
		fmt.Printf("!defer: %d\n", i)
	}
}

func trace ( s string,){
	fmt.Println("Entering:",s)
}

func untrace (s string){
	fmt.Println("leaving :",s)
}

func funcA(){
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func funcB(){
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	funcA()
}