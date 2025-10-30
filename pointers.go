package main

import "fmt"

func memoryCheck() {
	var num = 10
	fmt.Printf("This variable %d is located at %p in memory", num, &num)
}