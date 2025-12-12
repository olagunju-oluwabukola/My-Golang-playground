package main

import (
	"fmt"
	"math"
)

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

func mainArray(){
	list := [4]int{10, 76, 90,75,}
	x:= sum(&list)
	fmt.Printf("the sum of the array is: %d", x)
}
	func sum(a *[4]int) (sum int){
		// for i:= 0; i< len(*a); i++ {
		// 	sum += (*a)[i]
		// }
		// return

		for _, v :=range a {
			sum += v
		}
		return
	}

func mainLoop(){
	var numbers [14] int

 loopArray(&numbers)

	fmt.Printf("lsist of numbers : %d\n", numbers)
}

func loopArray (b *[14] int){
	for i:=0; i < len(b); i++{
		(*b)[i] = i
	}

}


func arrayLoop(){
	list:= []int{8,10,3,5,6}
	sum := 0
	max := math.MinInt
	min := math.MaxInt

	for _, v := range list{
		sum += v
		if v < min {
			min = v
	}
	if v > max {
		max = v
	}
	}
	average := float64(sum) / float64(len(list))
	fmt.Println("Array:", list)
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", average)
	fmt.Println("Minimum:", min)
	fmt.Println("Maximum:", max)
}