package main

import "fmt"

func ageCheck() {
	var age int
	fmt.Print("How old are you")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You're still a minor")
	} else if age == 18 {
		fmt.Println("you are not eligible")
	}else {
		fmt.Println("You're almost there")
	}
}