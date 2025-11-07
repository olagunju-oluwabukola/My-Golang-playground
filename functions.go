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