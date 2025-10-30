package main

import "fmt"

func myArray  (){
	var users [50] string
	users [10] = "Bola"
	users [30] = "Alex"
	fmt.Println("Lenth of the Array \n", len(users))
	var list = [2] string{"Ola", "lagos"}
	fmt.Println(list[1])

}



func mySlice (){
	var models [] string
	models = append(models, "Camry", "GLE")
	fmt.Println(models)
	// fmt.Print(models[3])
	fmt.Println(models[0])
	fmt.Println(len(models))

}