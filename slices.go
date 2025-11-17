package main

import "fmt"

func sliceCheck() {
	var slice [5]int
	 s := slice[1:3]
	 fmt.Println(s)

	for i := 0; i < len(slice); i++ {
		slice[i] = i
		fmt.Println(i)
	}
	 fmt.Println(s)
	 fmt.Printf("the lenght of the original array is %d\n", len(slice))
	 fmt.Printf("the lenght of the slice is %d\n", len(s))
	 fmt.Printf("the capacity of the original array is %d\n", cap(slice))
	 fmt.Printf("the capacity of the slice is %d\n", cap(s))

}

func makeSlice(){
	slice1 := [10] int {1,2,3,}
}
