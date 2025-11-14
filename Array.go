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
func Arr1(){
	var arr1 [5] int

	for i:=0; i < len(arr1); i++{
		arr1[i] = i*2
		fmt.Printf("the value of array at index %d is %d\n", i, arr1[i])
	}

	// for i:=0; i < len(arr1); i++{
	// 	fmt.Printf("the value of array at index %d is %d\n", i, arr1[i])
	// }
fmt.Println("------------with range------------")
	for i := range arr1{
		arr1[i] = i *2

		fmt.Printf("the value of array at index %d is %d\n", i, arr1[i])
	}

}









